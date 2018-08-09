#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

APP_DIR=${APP_DIR}

docker build \
-f ${APP_DIR}/Dockerfile \
-t mozey-gateway --rm \
${APP_DIR}

