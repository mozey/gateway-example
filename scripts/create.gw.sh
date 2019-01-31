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

EXPECTED_ARGS=1
OPTIONAL_ARGS=2

if [ $# -lt ${EXPECTED_ARGS} ]
then
  echo "Usage:"
  echo "  `basename $0` FN [STAGE]"
  echo ""
  echo "This script creates a lambda function"
  echo ""
  echo "Example:"
  echo "  `basename $0` api stage"
  exit 1
fi

# Function and stage must be lowercase
FN=$(echo "$1" | tr '[:upper:]' '[:lower:]')
# Namespaces must be uppercase
NS_FN=$(echo "_$1" | tr '[:lower:]' '[:upper:]')
if [ $# -eq ${OPTIONAL_ARGS} ]
then
    STAGE=$(echo "$2" | tr '[:upper:]' '[:lower:]')
    NS_STAGE=$(echo "_$2" | tr '[:lower:]' '[:upper:]')
else
    NS_STAGE=""
fi

# This script requires an admin profile
AWS_PROFILE=${AWS_PROFILE}

APP_DIR=${APP_DIR}
APP_REGION=${APP_REGION}

# Must fail if these are not set in config
APP_VAR="APP_LAMBDA_NAME${NS_FN}${NS_STAGE}"
if ! APP_LAMBDA_NAME=`printenv ${APP_VAR}`;
then
    echo "undefined config ${APP_VAR}"
    exit 1
fi
APP_VAR="APP_LAMBDA_HANDLER${NS_FN}${NS_STAGE}"
if ! APP_LAMBDA_HANDLER=`printenv ${APP_VAR}`;
then
    echo "undefined config ${APP_VAR}"
    exit 1
fi

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

# API...........................................................................

prompt_continue "Creating API deployment for ${APP_LAMBDA_NAME}"

APP_ACCOUNT=$(aws sts get-caller-identity | jq -r .Account)

APP_GW_ID=$(${DEBUG} aws apigateway create-rest-api --name ${APP_LAMBDA_NAME} | \
jq -r .id)
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_GW_ID${NS_FN}${NS_STAGE}" -value "${APP_GW_ID}"

# Get APP_GW_ID by name
APP_GW_ID=$(aws apigateway get-rest-apis | \
jq -r ".items[]  | select(.name == \"${APP_LAMBDA_NAME}\") | .id")

APP_GW_ROOT=$(aws apigateway get-resources --rest-api-id ${APP_GW_ID} | \
jq -r .items[0].id)
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_GW_ROOT${NS_FN}${NS_STAGE}" -value "${APP_GW_ROOT}"

APP_GW_PROXY=$(${DEBUG} aws apigateway create-resource \
--rest-api-id ${APP_GW_ID} \
--parent-id ${APP_GW_ROOT} \
--path-part "{proxy+}" | jq -r .id)
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_GW_PROXY${NS_FN}${NS_STAGE}" -value "${APP_GW_PROXY}"

${DEBUG} aws apigateway put-method --rest-api-id ${APP_GW_ID} \
--resource-id ${APP_GW_ROOT} --http-method ANY \
--authorization-type NONE

${DEBUG} aws apigateway put-method --rest-api-id ${APP_GW_ID} \
--resource-id ${APP_GW_PROXY} --http-method ANY \
--authorization-type NONE

APP_LAMBDA_ARN=$(aws lambda get-function \
--function-name ${APP_LAMBDA_NAME} | jq -r .Configuration.FunctionArn)
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_LAMBDA_ARN${NS_FN}${NS_STAGE}" -value "${APP_LAMBDA_ARN}"

${DEBUG} aws apigateway put-integration --rest-api-id ${APP_GW_ID} \
--resource-id ${APP_GW_ROOT} --http-method ANY --type AWS \
--integration-http-method POST \
--uri arn:aws:apigateway:${APP_REGION}:lambda:path/2015-03-31/functions/${APP_LAMBDA_ARN}/invocations
# aws apigateway get-integration --rest-api-id ${APP_GW_ID} --resource-id ${APP_GW_ROOT} --http-method ANY

${DEBUG} aws apigateway put-integration --rest-api-id ${APP_GW_ID} \
--resource-id ${APP_GW_PROXY} --http-method ANY --type AWS_PROXY \
--integration-http-method POST \
--uri arn:aws:apigateway:${APP_REGION}:lambda:path/2015-03-31/functions/${APP_LAMBDA_ARN}/invocations
# aws apigateway get-integration --rest-api-id ${APP_GW_ID} --resource-id ${APP_GW_PROXY} --http-method ANY

APP_LAMBDA_PERM=$(uuidgen)
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_LAMBDA_PERM${NS_FN}${NS_STAGE}" -value "${APP_LAMBDA_PERM}"

${DEBUG} aws lambda add-permission --function-name ${APP_LAMBDA_NAME} \
--statement-id ${APP_LAMBDA_PERM} \
--action lambda:InvokeFunction --principal apigateway.amazonaws.com \
--source-arn arn:aws:execute-api:${APP_REGION}:${APP_ACCOUNT}:${APP_GW_ID}/*/*/*

# Add multi stage logic here...
APP_GW_STAGE_NAME=main
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_GW_STAGE_NAME${NS_FN}${NS_STAGE}" -value "${APP_GW_STAGE_NAME}"

${DEBUG} aws apigateway create-deployment --rest-api-id ${APP_GW_ID} \
--stage-name ${APP_GW_STAGE_NAME}

APP_LAMBDA_BASE="https://${APP_GW_ID}.execute-api.${APP_REGION}.amazonaws.com/${APP_GW_STAGE_NAME}"
${DEBUG} ${APP_DIR}/config -env prod \
-key "APP_LAMBDA_BASE${NS_FN}${NS_STAGE}" -value "${APP_LAMBDA_BASE}"

# ..............................................................................
echo "GW Done"

