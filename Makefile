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

# clean up
clean:
	go clean

# dev...........................................................................
# Local server with live reload

build.dev: dependencies clean
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
dev: dependencies
	@make restart.dev
	@fswatch -or ./ -e dev.out | \
	xargs -n1 -I{} make restart.dev || make kill.dev

# container.....................................................................
# TODO Docker container with live reload

# lambda........................................................................
test: dependencies
	gotest ./internal...

build: clean test
	/usr/bin/env bash -c scripts/build.sh

deploy: build
	/usr/bin/env bash -c scripts/deploy.sh







