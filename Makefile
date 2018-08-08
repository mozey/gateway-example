# Copied from
# https://gist.github.com/lantins/e83477d8bccab83f078d

# binary name to kill/restart
PROG = dev.out

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

# run formatting tool and build
build: dependencies clean
	/usr/bin/env bash -c scripts/build.dev.sh

# run unit tests with code coverage
test: dependencies
	go test -cover

# generate code coverage report
coverage: test
	go test -coverprofile=.coverage.out
	go tool cover -html=.coverage.out

# attempt to kill running server
kill:
	@echo kill
	-@killall -9 $(PROG) 2>/dev/null || true

# attempt to build and start server
restart:
	@echo restart
	@make kill
	@make build; (if [ "$$?" -eq 0 ]; then (./${PROG} &); fi)

# watch .go files for changes then recompile & try to start server
# will also kill server after ctrl+c
dev: dependencies
	@make restart
	@fswatch -or ./ -e dev.out | xargs -n1 -I{} make restart || make kill
