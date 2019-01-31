#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

EXPECTED_ARGS=1
OPTIONAL_ARGS=2

if [ $# -lt ${EXPECTED_ARGS} ]
then
  echo "Usage:"
  echo "  `basename $0` AWS_PROFILE [PROMPT]"
  echo ""
  echo "This script deploys the lambda function."
  echo "It should not be executing directly, use the Makefile instead."
  echo "Set PROMPT to PROMPT_DISABLED to disable prompting"
  exit 1
fi

AWS_PROFILE="$1"

if [ $# -eq ${OPTIONAL_ARGS} ]
then
    PROMPT="$2"
else
    PROMPT="PROMPT_ENABLED"
fi

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_LAMBDA_NAME_API=${APP_LAMBDA_NAME_API}

#if [ ${AWS_PROFILE} != "xxx" ]
#then
#    echo "This script requires the admin AWS_PROFILE"
#    exit 1
#fi

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
    prompt_continue "Deploy lambda fn ${APP_LAMBDA_NAME_API} to prod?"
fi

echo
echo "Updating function code..................................................."
echo
aws lambda update-function-code --function-name ${APP_LAMBDA_NAME_API} \
--zip-file fileb://${APP_DIR}/build/main.zip

echo
echo "Updating function config................................................."
echo
# TODO Combine with fn above?
# NOTE Use lambda default service key for KMS encryption of env vars
LAMBDA_ENV_CSV=$(${APP_DIR}/config -env prod -csv)
aws lambda update-function-configuration \
--function-name ${APP_LAMBDA_NAME_API} \
--kms-key-arn "" \
--environment Variables={${LAMBDA_ENV_CSV}}

