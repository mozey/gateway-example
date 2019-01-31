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
APP_LAMBDA_NAME_API=${APP_LAMBDA_NAME_API}
APP_LAMBDA_HANDLER_API=${APP_LAMBDA_HANDLER_API}
APP_REGION=${APP_REGION}
AWS_PROFILE=${AWS_PROFILE}
APP_LAMBDA_POLICY_NAME=${APP_LAMBDA_POLICY_NAME}

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

prompt_continue "Create lambda function ${APP_LAMBDA_NAME_API}"

APP_ACCOUNT=$(aws sts get-caller-identity | jq -r .Account)

${APP_DIR}/config -env prod \
-key "APP_ACCOUNT" -value "${APP_ACCOUNT}" \
-key "APP_LAMBDA_POLICY_NAME" -value "${APP_LAMBDA_POLICY_NAME}" \

aws iam create-role --role-name ${APP_LAMBDA_NAME_API} \
--assume-role-policy-document file://${APP_DIR}/deployments/aws/trust-policy.json
APP_LAMBDA_ROLE_ARN_API="arn:aws:iam::${APP_ACCOUNT}:role/${APP_LAMBDA_NAME_API}"
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_ROLE_ARN_API" -value "${APP_LAMBDA_ROLE_ARN_API}"

APP_LAMBDA_POLICY_ARN=arn:aws:iam::${APP_ACCOUNT}:policy/${APP_LAMBDA_POLICY_NAME}
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_POLICY_ARN" -value "${APP_LAMBDA_POLICY_ARN}"

aws iam attach-role-policy --role-name ${APP_LAMBDA_NAME_API} \
--policy-arn ${APP_LAMBDA_POLICY_ARN}

aws lambda create-function --function-name ${APP_LAMBDA_NAME_API} --runtime go1.x \
--role ${APP_LAMBDA_ROLE_ARN_API} \
--handler ${APP_LAMBDA_HANDLER_API} --zip-file fileb://${APP_DIR}/build/main.zip \
--timeout 300 \
--memory-size 128

# NOTE Use lambda default service key for KMS encryption of env vars
LAMBDA_ENV_CSV=$(${APP_DIR}/config -env prod -csv)
aws lambda update-function-configuration \
--function-name ${APP_LAMBDA_NAME_API} \
--kms-key-arn "" \
--environment Variables={${LAMBDA_ENV_CSV}}


# API...........................................................................

echo "Creating API deployment..."
echo ""

APP_GW_ID_API_NAME=${APP_LAMBDA_NAME_API}

APP_GW_ID_API=$(aws apigateway create-rest-api --name ${APP_LAMBDA_NAME_API} | \
jq -r .id)
${APP_DIR}/config -env prod \
-key "APP_GW_ID_API" -value "${APP_GW_ID_API}"

# Get APP_GW_ID_API by name
APP_GW_ID_API=$(aws apigateway get-rest-apis | \
jq -r ".items[]  | select(.name == \"${APP_LAMBDA_NAME_API}\") | .id")

APP_GW_ROOT_API=$(aws apigateway get-resources --rest-api-id ${APP_GW_ID_API} | \
jq -r .items[0].id)
${APP_DIR}/config -env prod \
-key "APP_GW_ROOT_API" -value "${APP_GW_ROOT_API}"

APP_GW_PROXY_API=$(aws apigateway create-resource \
--rest-api-id ${APP_GW_ID_API} \
--parent-id ${APP_GW_ROOT_API} \
--path-part "{proxy+}" | jq -r .id)
${APP_DIR}/config -env prod \
-key "APP_GW_PROXY_API" -value "${APP_GW_PROXY_API}"

aws apigateway put-method --rest-api-id ${APP_GW_ID_API} \
--resource-id ${APP_GW_ROOT_API} --http-method ANY \
--authorization-type NONE

aws apigateway put-method --rest-api-id ${APP_GW_ID_API} \
--resource-id ${APP_GW_PROXY_API} --http-method ANY \
--authorization-type NONE

APP_LAMBDA_ARN_API=$(aws lambda get-function \
--function-name ${APP_LAMBDA_NAME_API} | jq -r .Configuration.FunctionArn)
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_ARN_API" -value "${APP_LAMBDA_ARN_API}"

aws apigateway put-integration --rest-api-id ${APP_GW_ID_API} \
--resource-id ${APP_GW_ROOT_API} --http-method ANY --type AWS \
--integration-http-method POST \
--uri arn:aws:apigateway:${APP_REGION}:lambda:path/2015-03-31/functions/${APP_LAMBDA_ARN_API}/invocations
# aws apigateway get-integration --rest-api-id ${APP_GW_ID_API} --resource-id ${APP_GW_ROOT_API} --http-method ANY

aws apigateway put-integration --rest-api-id ${APP_GW_ID_API} \
--resource-id ${APP_GW_PROXY_API} --http-method ANY --type AWS_PROXY \
--integration-http-method POST \
--uri arn:aws:apigateway:${APP_REGION}:lambda:path/2015-03-31/functions/${APP_LAMBDA_ARN_API}/invocations
# aws apigateway get-integration --rest-api-id ${APP_GW_ID_API} --resource-id ${APP_GW_PROXY_API} --http-method ANY

APP_LAMBDA_PERM_API=$(uuidgen)
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_PERM_API" -value "${APP_LAMBDA_PERM_API}"

aws lambda add-permission --function-name ${APP_LAMBDA_NAME_API} \
--statement-id ${APP_LAMBDA_PERM_API} \
--action lambda:InvokeFunction --principal apigateway.amazonaws.com \
--source-arn arn:aws:execute-api:${APP_REGION}:${APP_ACCOUNT}:${APP_GW_ID_API}/*/*/*

# Add multi stage logic here...
APP_GW_STAGE_NAME_API=main
${APP_DIR}/config -env prod \
-key "APP_GW_STAGE_NAME_API" -value "${APP_GW_STAGE_NAME_API}"

aws apigateway create-deployment --rest-api-id ${APP_GW_ID_API} \
--stage-name ${APP_GW_STAGE_NAME_API}

APP_LAMBDA_BASE_API="https://${APP_GW_ID_API}.execute-api.${APP_REGION}.amazonaws.com/${APP_GW_STAGE_NAME_API}"
${APP_DIR}/config -env prod \
-key "APP_LAMBDA_BASE_API" -value "${APP_LAMBDA_BASE_API}"

# ..............................................................................
echo "Done"

