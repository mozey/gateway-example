#!/usr/bin/env bash

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_FN_NAME=${APP_FN_NAME}

echo "Deploying lambda fn ${APP_FN_NAME}..."
aws lambda update-function-code --function-name ${APP_FN_NAME} \
--zip-file fileb://${APP_DIR}/build/main.zip

