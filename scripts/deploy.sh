#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

# ..............................................................................
# REVERT to a previous version
# ssh to the deployment server and
# rename build/main-${APP_VERSION}.zip to build/main.zip
# then run the deploy script manually
# ..............................................................................

SERVERS=(
#   SERVER_IP       APP_DIR
    1.2.3.4         /path/to/go/src/github.com/mozey/gateway
)

EXPECTED_ARGS=1
OPTIONAL_ARGS=2

if [ $# -lt ${EXPECTED_ARGS} ]
then
  echo "Usage:"
  echo "  `basename $0` AWS_PROFILE KEYFILE [TARGET]"
  echo ""
  echo "This script deploys the lambda function using an intermediate server"
  exit 1
fi

AWS_PROFILE="$1"
KEYFILE="$2"

if [ $# -eq ${OPTIONAL_ARGS} ]
then
    TARGET="$3"
else
    TARGET="origin/master"
    echo "Using default TARGET: ${TARGET}"
fi

read -p "Have you committed and pushed (y/n)? " -n 1 -r
echo    # move to a new line
if [[ ${REPLY} =~ ^[Yy]$ ]]
then
    for (( i = 0 ; i < ${#SERVERS[@]} ; i = i + 2 )) do
        SERVER_IP=${SERVERS[$i]}
        APP_DIR=${SERVERS[$i+1]}
        ssh -i ${KEYFILE} ubuntu@${SERVER_IP} \
            "cd ${APP_DIR} && git fetch --all && git checkout --force ${TARGET}"
        ssh -i ${KEYFILE} ubuntu@${SERVER_IP} \
            'source /home/ubuntu/.profile && cd '${APP_DIR}' && eval "$(./config -env prod)" && make build && ./scripts/deploy.local.sh '${AWS_PROFILE}' PROMPT_DISABLED'
    done
fi
