#!/usr/bin/env bash

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_LAMBDA_NAME=${APP_LAMBDA_NAME}
AWS_PROFILE=${AWS_PROFILE}

function prompt_continue() {
    read -p "${1} AWS_PROFILE=${AWS_PROFILE} continue (y)? " -n 1 -r
    echo ""
    if [[ ${REPLY} =~ ^[Yy]$ ]]
    then
        :
    else
        echo "Abort"
        exit 1
    fi
}

prompt_continue "Deploying lambda fn ${APP_LAMBDA_NAME}..."

aws lambda update-function-code --function-name ${APP_LAMBDA_NAME} \
--zip-file fileb://${APP_DIR}/build/main.zip

# NOTE Use lambda default service key for KMS encryption of env vars
LAMBDA_ENV_CSV=$(${APP_DIR}/config -env prod -csv)
aws lambda update-function-configuration \
--function-name ${APP_LAMBDA_NAME} \
--kms-key-arn "" \
--environment Variables={${LAMBDA_ENV_CSV}}
