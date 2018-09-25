#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_LAMBDA_NAME=${APP_LAMBDA_NAME}
APP_LAMBDA_HANDLER=${APP_LAMBDA_HANDLER}
APP_REGION=${APP_REGION}
AWS_PROFILE=${AWS_PROFILE}
DEPLOYMENTS_DIR=${APP_DIR}/deployments

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

# Lambda fn.....................................................................

prompt_continue "Create lambda function ${APP_LAMBDA_NAME}"

APP_ACCOUNT=$(aws sts get-caller-identity | jq -r .Account)
APP_LAMBDA_POLICY_NAME="S2S-Lambda"

${APP_DIR}/config -env prod \
-key "APP_ACCOUNT" -value "${APP_ACCOUNT}" \
-key "APP_LAMBDA_POLICY_NAME" -value "${APP_LAMBDA_POLICY_NAME}" \

aws iam create-role --role-name ${APP_LAMBDA_NAME} \
--assume-role-policy-document file://${DEPLOYMENTS_DIR}/aws/trust-policy.json
APP_LAMBDA_ROLE_ARN="arn:aws:iam::${APP_ACCOUNT}:role/${APP_LAMBDA_NAME}"
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_ROLE_ARN" -value "${APP_LAMBDA_ROLE_ARN}"

APP_LAMBDA_POLICY_ARN=arn:aws:iam::${APP_ACCOUNT}:policy/${APP_LAMBDA_POLICY_NAME}
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_POLICY_ARN" -value "${APP_LAMBDA_POLICY_ARN}"

aws iam attach-role-policy --role-name ${APP_LAMBDA_NAME} \
--policy-arn ${APP_LAMBDA_POLICY_ARN}

# NOTE the vpc-config requires EC2 permissions on the S2S-Lambda policy
aws lambda create-function --function-name ${APP_LAMBDA_NAME} --runtime go1.x \
--role ${APP_LAMBDA_ROLE_ARN} \
--handler ${APP_LAMBDA_HANDLER} --zip-file fileb://${APP_DIR}/build/main.zip \
--vpc-config file://${DEPLOYMENTS_DIR}/aws/vpc-config.json \
--timeout 300 \
--memory-size 128

# NOTE Use lambda default service key for KMS encryption of env vars
LAMBDA_ENV_CSV=$(${APP_DIR}/config -env prod -csv)
aws lambda update-function-configuration \
--function-name ${APP_LAMBDA_NAME} \
--kms-key-arn "" \
--environment Variables={${LAMBDA_ENV_CSV}}


# API...........................................................................

echo "Creating API deployment..."
echo ""

APP_API_NAME=${APP_LAMBDA_NAME}

APP_API=$(aws apigateway create-rest-api --name ${APP_LAMBDA_NAME} | \
jq -r .id)
${APP_DIR}/config -env prod \
-key "APP_API" -value "${APP_API}"

# Get APP_API by name
#APP_API=$(aws apigateway get-rest-apis | \
#jq -r ".items[]  | select(.name == \"${APP_LAMBDA_NAME}\") | .id")

APP_API_ROOT=$(aws apigateway get-resources --rest-api-id ${APP_API} | \
jq -r .items[0].id)
${APP_DIR}/config -env prod \
-key "APP_API_ROOT" -value "${APP_API_ROOT}"

APP_API_PROXY=$(aws apigateway create-resource \
--rest-api-id ${APP_API} \
--parent-id ${APP_API_ROOT} \
--path-part "{proxy+}" | jq -r .id)
${APP_DIR}/config -env prod \
-key "APP_API_PROXY" -value "${APP_API_PROXY}"

aws apigateway put-method --rest-api-id ${APP_API} \
--resource-id ${APP_API_ROOT} --http-method ANY \
--authorization-type NONE

aws apigateway put-method --rest-api-id ${APP_API} \
--resource-id ${APP_API_PROXY} --http-method ANY \
--authorization-type NONE

APP_LAMBDA_ARN=$(aws lambda get-function \
--function-name ${APP_LAMBDA_NAME} | jq -r .Configuration.FunctionArn)
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_ARN" -value "${APP_LAMBDA_ARN}"

aws apigateway put-integration --rest-api-id ${APP_API} \
--resource-id ${APP_API_ROOT} --http-method ANY --type AWS \
--integration-http-method POST \
--uri arn:aws:apigateway:${APP_REGION}:lambda:path/2015-03-31/functions/${APP_LAMBDA_ARN}/invocations
# aws apigateway get-integration --rest-api-id ${APP_API} --resource-id ${APP_API_ROOT} --http-method ANY

aws apigateway put-integration --rest-api-id ${APP_API} \
--resource-id ${APP_API_PROXY} --http-method ANY --type AWS_PROXY \
--integration-http-method POST \
--uri arn:aws:apigateway:${APP_REGION}:lambda:path/2015-03-31/functions/${APP_LAMBDA_ARN}/invocations
# aws apigateway get-integration --rest-api-id ${APP_API} --resource-id ${APP_API_PROXY} --http-method ANY

APP_LAMBDA_PERM=$(uuidgen)
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_PERM" -value "${APP_LAMBDA_PERM}"

aws lambda add-permission --function-name ${APP_LAMBDA_NAME} \
--statement-id ${APP_LAMBDA_PERM} \
--action lambda:InvokeFunction --principal apigateway.amazonaws.com \
--source-arn arn:aws:execute-api:${APP_REGION}:${APP_ACCOUNT}:${APP_API}/*/*/*

# Add multi stage logic here...
APP_API_STAGE_NAME=main
${APP_DIR}/config -env prod \
-key "APP_API_STAGE_NAME" -value "${APP_API_STAGE_NAME}"

aws apigateway create-deployment --rest-api-id ${APP_API} \
--stage-name ${APP_API_STAGE_NAME}

APP_LAMBDA_BASE="https://${APP_API}.execute-api.${APP_REGION}.amazonaws.com/${APP_API_STAGE_NAME}"
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_BASE" -value "${APP_LAMBDA_BASE}"

# ..............................................................................
echo "Done"

