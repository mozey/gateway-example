# Copied from
# https://gist.github.com/lantins/e83477d8bccab83f078d

# binary name to kill/restart
PROG_DEV = dev.out

# targets not associated with files
.PHONY: dependencies default build test coverage clean kill restart serve

dependencies:
	@command -v fswatch --version >/dev/null 2>&1 || \
	{ printf >&2 "Install fswatch, run: brew install fswatch\n"; exit 1; }

# default targets to run when only running `make`
default: dependencies test

# dev...........................................................................
# Local server with live reload
clean.dev:
	go clean

# Tests are cached,
# to re-run all tests clear it first
test:
	go clean -testcache
	@make test.cache

test.cache: dependencies
	./config -env dev -compare prod || \
    { printf >&2 "dev and prod keys don't match\n"; exit 1; }
ifneq ($(AWS_PROFILE),aws-local)
	echo "Tests must use dev env"
	exit 1
endif
	gotest ./pkg...
	gotest ./internal...

build.dev: dependencies clean.dev
	/usr/bin/env bash -c scripts/build.dev.sh

# attempt to kill running server
kill.dev:
	@echo kill.dev
	-@killall -9 $(PROG_DEV) 2>/dev/null || true

# attempt to build and start server
restart.dev:
	@echo restart.dev
	@make kill.dev
	@make build.dev; (if [ "$$?" -eq 0 ]; then (./${PROG_DEV} &); fi)

# watch .go files for changes then recompile & try to start server
# will also kill server after ctrl+c
# fswatch includes everything unless an exclusion filter says otherwise
# https://stackoverflow.com/a/37237681/639133
dev: dependencies
	@make restart.dev
	@fswatch -or --exclude ".*" --include "\\.go$$" ./ | \
	xargs -n1 -I{} make restart.dev || make kill.dev

# lambda........................................................................

clean:
	go clean

build: clean
ifeq ($(AWS_PROFILE),aws-local)
	echo "Build must use prod env"
	exit 1
endif
	/usr/bin/env bash -c scripts/build.sh

# Deploy does not use the makefile






