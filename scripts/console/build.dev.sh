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
AWS_PROFILE=${AWS_PROFILE}

if [ "${AWS_PROFILE}" != "aws-local" ]
then
    echo "dev build must use local services"
    exit 1
fi

echo "Building exe"
cd ${APP_DIR}
env go build \
-o ${APP_DIR}/gateway-console.out \
./cmd/dev/console

chmod u+x ${APP_DIR}/gateway-console.out
