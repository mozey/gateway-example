#!/usr/bin/env bash

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_LAMBDA_NAME=${APP_LAMBDA_NAME}

echo "Deploying lambda fn ${APP_LAMBDA_NAME}..."
aws lambda update-function-code --function-name ${APP_LAMBDA_NAME} \
--zip-file fileb://${APP_DIR}/build/main.zip

