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

EXPECTED_ARGS=3
OPTIONAL_ARGS=4

if [ $# -lt ${EXPECTED_ARGS} ]
then
  echo "Usage:"
  echo "  `basename $0` AWS_PROFILE FN STAGE [PROMPT]"
  echo ""
  echo "This script deploys the lambda function."
  echo "It should not be executing directly, use the Makefile instead."
  echo "Set PROMPT to PROMPT_DISABLED to disable prompting"
  exit 1
fi

AWS_PROFILE="$1"

# Function and stage must be lowercase
FN=$(echo "$2" | tr '[:upper:]' '[:lower:]')
STAGE=$(echo "$3" | tr '[:upper:]' '[:lower:]')
# Namespaces must be uppercase
NS_FN=$(echo "_$2" | tr '[:lower:]' '[:upper:]')
if [ ${STAGE} != "main" ]
then
    NS_STAGE=$(echo "_$3" | tr '[:lower:]' '[:upper:]')
else
    # Main has empty stage namespace
    NS_STAGE=""
fi

if [ $# -eq ${OPTIONAL_ARGS} ]
then
    PROMPT="$4"
else
    PROMPT="PROMPT_ENABLED"
fi

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}

# Must fail if these are not set in config
APP_VAR="APP_LAMBDA_NAME${NS_FN}${NS_STAGE}"
if ! APP_LAMBDA_NAME=`printenv ${APP_VAR}`;
then
    echo "undefined config ${APP_VAR}"
    exit 1
fi

function prompt_continue() {
    read -p "${1} continue (y)? " -n 1 -r
    echo ""
    if [[ ${REPLY} =~ ^[Yy]$ ]]
    then
        :
    else
        echo "Abort"
        exit 1
    fi
}

if [ ${PROMPT} != "PROMPT_DISABLED" ]
then
    prompt_continue "Deploy lambda fn ${APP_LAMBDA_NAME} to prod?"
fi

echo
echo "Updating function code..................................................."
echo
${DEBUG} aws lambda update-function-code --function-name ${APP_LAMBDA_NAME} \
--zip-file fileb://${APP_DIR}/build/${FN}/main.zip

echo
echo "Updating function config................................................."
echo
# TODO Combine with fn above?
# NOTE Use lambda default service key for KMS encryption of env vars
LAMBDA_ENV_CSV=$(${APP_DIR}/config -env prod -csv)
${DEBUG} aws lambda update-function-configuration \
--function-name ${APP_LAMBDA_NAME} \
--kms-key-arn "" \
--environment Variables={${LAMBDA_ENV_CSV}}

