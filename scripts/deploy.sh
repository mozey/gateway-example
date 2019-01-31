#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

#DEBUG=echo
DEBUG=

# ..............................................................................
# REVERT to a previous version
# ssh to the deployment server and
# rename build/api/main-${APP_VERSION}.zip to build/api/main.zip
# then run the deploy script manually
# ..............................................................................

SERVERS=(
#   SERVER_IP       APP_DIR
    1.2.3.4         /path/to/go/src/github.com/mozey/gateway
)

EXPECTED_ARGS=3

if [ $# -lt ${EXPECTED_ARGS} ]
then
  echo "Usage:"
  echo "  `basename $0` KEYFILE FN STAGE"
  echo ""
  echo "This script deploys the lambda function using an intermediate server"
  exit 1
fi

KEYFILE="$1"
AWS_PROFILE=${AWS_PROFILE}

# Function and stage must be lowercase
FN=$(echo "$2" | tr '[:upper:]' '[:lower:]')
STAGE=$(echo "$3" | tr '[:upper:]' '[:lower:]')

# Always deploy from master.
# To revert to an older version,
# logon to deployment server,
# copy archived build and run,
# deploy.local.sh with
TARGET="origin/master"

read -p "Have you committed and pushed (y/n)? " -n 1 -r
echo    # move to a new line
if [[ ${REPLY} =~ ^[Yy]$ ]]
then
    for (( i = 0 ; i < ${#SERVERS[@]} ; i = i + 2 )) do
        SERVER_IP=${SERVERS[$i]}
        APP_DIR=${SERVERS[$i+1]}
        ${DEBUG} ssh -i ${KEYFILE} ubuntu@${SERVER_IP} \
            "cd ${APP_DIR} && git fetch --all && git checkout --force ${TARGET}"
        ${DEBUG} ssh -i ${KEYFILE} ubuntu@${SERVER_IP} \
            'source /home/ubuntu/.profile && cd '${APP_DIR}' && eval "$(./config -env prod)" && make build.'${FN}' && ./scripts/deploy.local.sh '${AWS_PROFILE}' '${FN}' '${STAGE}' PROMPT_DISABLED'
    done
fi
