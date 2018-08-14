#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

# Env
GOPATH=${GOPATH}
APP_DIR=${APP_DIR}
export APP_DEBUG=true

# Build config util
go get github.com/mozey/config
cd ${GOPATH}/src/github.com/mozey/config
go build \
-ldflags "-X main.AppDir=${APP_DIR}" \
-o ${APP_DIR}/config ./cmd/config

# Git hooks
chmod u+x ${APP_DIR}/githooks/*.sh
ln -sf ${APP_DIR}/githooks/pre-commit.sh ${APP_DIR}/.git/hooks/pre-commit

read -p "Reset config files (y)? " -n 1 -r
echo ""
echo ""
if [[ ${REPLY} =~ ^[Yy]$ ]]
then
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
fi


