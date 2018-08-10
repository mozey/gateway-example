#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

GOPATH=${GOPATH}

# Compile config util
export APP_DEBUG=true
APP_DIR=${GOPATH}/src/github.com/mozey/gateway
cd ${APP_DIR}
go build -ldflags "-X main.AppDir=${APP_DIR}" -o ./config ./cmd/config

# Create dev config
cp ${APP_DIR}/config.dev.sample.json ${APP_DIR}/config.dev.json
${APP_DIR}/config \
-key APP_DIR -value ${APP_DIR} \
-update

# Create prod config
cp ${APP_DIR}/config.prod.sample.json ${APP_DIR}/config.prod.json
${APP_DIR}/config -env prod \
-key APP_DIR -value ${APP_DIR} \
-update


