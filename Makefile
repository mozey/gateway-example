# Copied from
# https://gist.github.com/lantins/e83477d8bccab83f078d

# binary name to kill/restart
PROG_API = gateway-api.out
PROG_CONSOLE = gateway-console.out

dependencies:
	@command -v fswatch --version >/dev/null 2>&1 || \
	{ printf >&2 "Install fswatch, run: brew install fswatch\n"; exit 1; }

# default targets to run when only running `make`
default: dependencies test

# ..............................................................................
clean:
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
	gotest ./cmd...
	gotest ./pkg...
	gotest ./internal...


# api...........................................................................
# Local server with live reload
build.api.dev: dependencies clean
	/usr/bin/env bash -c scripts/api/build.dev.sh

# attempt to kill running server
kill.api.dev:
	@echo kill.api.dev
	-@killall -9 $(PROG_API) 2>/dev/null || true

# attempt to build and start server
restart.api.dev:
	@echo restart.api.dev
	@make kill.api.dev
	@make build.api.dev; (if [ "$$?" -eq 0 ]; then (./${PROG_API} &); fi)

# watch .go files for changes then recompile & try to start server
# will also kill server after ctrl+c
# fswatch includes everything unless an exclusion filter says otherwise
# https://stackoverflow.com/a/37237681/639133
api: dependencies
	@make restart.api.dev
	@fswatch -or --exclude ".*" --include "\\.go$$" ./ | \
	xargs -n1 -I{} make restart.api.dev || make kill.api.dev

# console.......................................................................
# Local server with live reload
build.console.dev: dependencies clean
	/usr/bin/env bash -c scripts/console/build.dev.sh

# attempt to kill running server
kill.console.dev:
	@echo kill.console.dev
	-@killall -9 $(PROG_CONSOLE) 2>/dev/null || true

# attempt to build and start server
restart.console.dev:
	@echo restart.console.dev
	@make kill.console.dev
	@make build.console.dev; (if [ "$$?" -eq 0 ]; then (./${PROG_CONSOLE} &); fi)

# watch .go files for changes then recompile & try to start server
# will also kill server after ctrl+c
# fswatch includes everything unless an exclusion filter says otherwise
# https://stackoverflow.com/a/37237681/639133
console: dependencies
	@make restart.console.dev
	@fswatch -or --exclude ".*" --include "\\.go$$" ./ | \
	xargs -n1 -I{} make restart.console.dev || make kill.console.dev

# lambda........................................................................

build.api: clean
ifeq ($(AWS_PROFILE),aws-local)
	echo "Build must use prod env"
	exit 1
endif
	/usr/bin/env bash -c scripts/api/build.sh

build.console: clean
ifeq ($(AWS_PROFILE),aws-local)
	echo "Build must use prod env"
	exit 1
endif
	/usr/bin/env bash -c scripts/console/build.sh






