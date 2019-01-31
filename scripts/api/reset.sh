#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

APP_DIR=${APP_DIR}
APP_LAMBDA_NAME_API=${APP_LAMBDA_NAME_API}
APP_GW_SUBDOMAIN=${APP_GW_SUBDOMAIN}
APP_GW_PATH_API=${APP_GW_PATH_API}
APP_LAMBDA_POLICY_NAME=${APP_LAMBDA_POLICY_NAME}
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

prompt_continue "Delete lambda fn and API ${APP_LAMBDA_NAME_API}"


# Other managed/inline policies to detach/delete?
APP_LAMBDA_POLICY_ARN=$(aws iam list-attached-role-policies \
--role-name ${APP_LAMBDA_NAME_API} | \
jq -r ".AttachedPolicies[] | select(.PolicyName == \"${APP_LAMBDA_POLICY_NAME}\") | .PolicyArn") \
|| APP_LAMBDA_POLICY_ARN=""
if [ "${APP_LAMBDA_POLICY_ARN}" != "" ]
then
    echo "Detaching policy ${APP_LAMBDA_POLICY_ARN}"
    aws iam detach-role-policy --role-name ${APP_LAMBDA_NAME_API} \
    --policy-arn ${APP_LAMBDA_POLICY_ARN}
fi

DELETE_ROLE=1
aws iam get-role --role-name ${APP_LAMBDA_NAME_API} > /dev/null || DELETE_ROLE=0
if [ ${DELETE_ROLE} -eq 1 ]
then
    echo "Deleting IAM role"
    aws iam delete-role --role-name ${APP_LAMBDA_NAME_API}
fi

DELETE_FN=1
aws lambda get-function --function-name ${APP_LAMBDA_NAME_API} > /dev/null \
|| DELETE_FN=0
if [ ${DELETE_FN} -eq 1 ]
then
    echo "Deleting lambda fn"
    aws lambda delete-function --function-name ${APP_LAMBDA_NAME_API}
fi

# TODO Delete all path mappings?
BASE_PATH=${APP_GW_PATH_API}
if [ "${APP_GW_PATH_API}" = "" ]
then
    # Trying to delete an empty base path will error
    BASE_PATH="(none)"
fi
DELETE_BASE_PATH=1
aws apigateway get-base-path-mapping --domain-name ${APP_GW_SUBDOMAIN} \
--base-path ${BASE_PATH} > /dev/null \
|| DELETE_BASE_PATH=0
if [ ${DELETE_FN} -eq 1 ]
then
    aws apigateway delete-base-path-mapping \
    --domain-name ${APP_GW_SUBDOMAIN} --base-path ${BASE_PATH}
fi

APP_GW_ID_API=$(aws apigateway get-rest-apis | \
jq -r ".items[]  | select(.name == \"${APP_LAMBDA_NAME_API}\") | .id") \
|| APP_GW_ID_API=""
if [ "${APP_GW_ID_API}" != "" ]
then
    echo "Deleting API ${APP_GW_ID_API}"
    aws apigateway delete-rest-api --rest-api-id ${APP_GW_ID_API}
fi

echo ""
echo "Done"

