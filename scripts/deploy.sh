#!/usr/bin/env bash

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_LAMBDA_NAME=${APP_LAMBDA_NAME}
AWS_PROFILE=${AWS_PROFILE}

# Confirm profile
read -p "AWS_PROFILE = ${AWS_PROFILE} continue (y)? " -n 1 -r
echo ""
if [[ ${REPLY} =~ ^[Yy]$ ]]
then
    :
else
    echo "Abort"
    exit 1
fi

echo "Deploying lambda fn ${APP_LAMBDA_NAME}..."
aws lambda update-function-code --function-name ${APP_LAMBDA_NAME} \
--zip-file fileb://${APP_DIR}/build/main.zip

