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
APP_LAMBDA_HANDLER_CONSOLE=${APP_LAMBDA_HANDLER_CONSOLE}

# Set version string for current revision,
# versions for future shared packages can be appended to APP_VERSION.
# Update config with version so it get set on ENV by deploy script
APP_GIT_REV=$(git -C ${APP_DIR} rev-parse --short --verify HEAD)
UTC_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
APP_VERSION="${UTC_DATE}#console=${APP_GIT_REV}"
${APP_DIR}/config -env prod \
-key "APP_VERSION_CONSOLE" -value "${APP_VERSION}"

mkdir -p ${APP_DIR}/build/console

# TODO Make binary reproducible,
# the binary's hash must be the same for the same git revision
# https://blog.filippo.io/reproducing-go-binaries-byte-by-byte
echo "Building exe"
cd ${APP_DIR}
env GOOS=linux GOARCH=amd64 go build \
-o build/console/${APP_LAMBDA_HANDLER_CONSOLE} \
./cmd/gateway/console
# TODO Save hash of binary to prod env

echo "Delete old build"
NAME="main.zip"
rm -f ${APP_DIR}/build/console/${NAME}

echo "Zip new build"
zip -j ${APP_DIR}/build/console/${NAME} ${APP_DIR}/build/console/${APP_LAMBDA_HANDLER_CONSOLE}
# Add more build artifacts here...

echo "Backup build with version"
cp ${APP_DIR}/build/console/${NAME} ${APP_DIR}/build/console/main-${APP_VERSION}.zip
# TODO Only keep last X backups?

echo "List zip contents"
unzip -vl ${APP_DIR}/build/console/${NAME}
