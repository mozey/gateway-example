#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}

echo "Building exe"
cd ${APP_DIR}
env CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
-o build/container.out \
./cmd/dev

