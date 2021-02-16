#!/usr/bin/env bash

# Inspired by [pre-commit](https://golang.org/misc/git/pre-commit)
#
# scripts/config.sh will create a symlink at .git/hooks/pre-commit
#
# This script does not handle file names that contain spaces.

# WARNING Do not `set -eu`, otherwise `GO_FILES=...` will always fail
# and exit the script if no *.go files were changed

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
SCRIPT_NAME=$(basename -- "$0")
echo "${SCRIPT_DIR}/${SCRIPT_NAME}"

# Note that `gif diff` will not list un-staged files,
# but this script run after staging with `git commit -a` flag?
GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
[ -z "${GO_FILES}" ] && exit 0

# Run tests before merging with master
#echo "go vet..."
#go vet -x ${GO_FILES}

# Original script does not exit with error when gofmt complains?
echo "gofmt..."
gofmt -l ${GO_FILES}

#echo "golint..."
#golint -set_exit_status ${GO_FILES}


