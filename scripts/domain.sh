#!/usr/bin/env bash

# Set (e) exit on error
# Set (u) no-unset to exit on undefined variable
set -eu
# If any command in a pipeline fails,
# that return code will be used as the
# return code of the whole pipeline.
bash -c 'set -o pipefail'

E_BADARGS=100

# Must fail with "unbound variable" if these are not set
APP_DIR=${APP_DIR}
APP_FN_NAME=${APP_FN_NAME}
APP_API_PATH=${APP_API_PATH}
APP_API_CUSTOM=${APP_API_CUSTOM}
APP_API_DOMAIN=${APP_API_DOMAIN}
APP_REGION=${APP_REGION}
APP_API=${APP_API}
APP_API_STAGE_NAME=${APP_API_STAGE_NAME}

# ..............................................................................
# WARNING 2018-08-05 `aws apigateway create-domain-name` will fail with
# "Certificate must be in 'us-east-1'" for certs requested in a different region
CERT_REGION=us-east-1

APP_CERT_ARN=$(aws acm list-certificates --region ${CERT_REGION} | \
jq -r ".CertificateSummaryList[] | select(.DomainName == \"${APP_API_CUSTOM}\") | .CertificateArn")
if [ "${APP_CERT_ARN}" = "" ]
then
    echo "Requesting cert for ${APP_API_CUSTOM}"
    echo ""
    aws acm request-certificate \
    --region ${CERT_REGION} \
    --domain-name ${APP_API_CUSTOM} --validation-method DNS
    APP_CERT_ARN=$(aws acm list-certificates --region ${CERT_REGION} | \
    jq -r ".CertificateSummaryList[] | select(.DomainName == \"${APP_API_CUSTOM}\") | .CertificateArn")
fi

# ..............................................................................

DNS_VALIDATION=$(aws acm describe-certificate --region ${CERT_REGION} \
--certificate-arn ${APP_CERT_ARN} | \
jq -r .Certificate.DomainValidationOptions[0].ResourceRecord)
DNS_VALIDATION_CNAME=$(echo ${DNS_VALIDATION} | jq -r .Name)
DNS_VALIDATION_VALUE=$(echo ${DNS_VALIDATION} | jq -r .Value)
APP_DNS_HOSTED_ZONE=$(aws route53 list-hosted-zones | \
jq -r ".HostedZones[] | select(.Name == \"${APP_API_DOMAIN}.\") | .Id")
if [ "${APP_DNS_HOSTED_ZONE}" = "" ]
then
    echo "Invalid APP_API_DOMAIN: no matching hosted zones"
    exit ${E_BADARGS}
fi

CREATE_CNAME=$(aws route53 list-resource-record-sets --hosted-zone-id ${APP_DNS_HOSTED_ZONE} | \
jq -r ".ResourceRecordSets[] | select(.Name == \"${DNS_VALIDATION_CNAME}\") | .Name")
if [ "${CREATE_CNAME}" = "" ]
then
    echo "Creating CNAME record for DSN validation"
    echo ""
    echo "
    {
        \"Comment\": \"DNS validation for custom domain\",
        \"Changes\": [
            {
                \"Action\": \"CREATE\",
                \"ResourceRecordSet\": {
                    \"Name\": \"${DNS_VALIDATION_CNAME}\",
                    \"Type\": \"CNAME\",
                    \"TTL\": 300,
                    \"ResourceRecords\": [
                        {
                            \"Value\": \"${DNS_VALIDATION_VALUE}\"
                        }
                    ]
                }
            }
        ]
    }
    " > ${APP_DIR}/change-resource-record-sets.json
    aws route53 change-resource-record-sets --hosted-zone-id ${APP_DNS_HOSTED_ZONE} \
    --change-batch file://${APP_DIR}/change-resource-record-sets.json
fi

# ..............................................................................
echo "Check cert status"
echo ""
APP_CERT_STATUS=$(aws acm describe-certificate --region ${CERT_REGION} \
--certificate-arn ${APP_CERT_ARN} | \
jq -r .Certificate.Status)
if [ "${APP_CERT_STATUS}" != "ISSUED" ]
then
    echo "Invalid APP_CERT_STATUS ${APP_CERT_STATUS}"
    echo "It might take a while for validation to complete"
    echo ""
    exit ${E_BADARGS}
fi
echo "Certificate is verified"

# ..............................................................................
CREATE_API_DOMAIN=0
aws apigateway get-domain-name --domain-name ${APP_API_CUSTOM} > /dev/null \
|| CREATE_API_DOMAIN=1
if [ ${CREATE_API_DOMAIN} -eq 1 ]
then
    echo "Create API domain"
    echo ""
    aws apigateway create-domain-name \
    --domain-name ${APP_API_CUSTOM} \
    --certificate-name ${APP_API_CUSTOM} \
    --region ${APP_REGION} \
    --certificate-arn ${APP_CERT_ARN}
    APP_API_TARGET=$(aws apigateway get-domain-name \
    --domain-name ${APP_API_CUSTOM} | \
    jq -r .distributionDomainName)
fi

## ..............................................................................
BASE_PATH=${APP_API_PATH}
if [ "${APP_API_PATH}" = "" ]
then
    # Trying to delete an empty base path will error
    BASE_PATH="(none)"
fi
# TODO Filter on base path when checking if mapping exists
PATH_MAPPINGS=$(aws apigateway get-base-path-mappings \
--domain-name ${APP_API_CUSTOM} | \
jq ".items | length")
if [ "${PATH_MAPPINGS}" = "0" ]
then
    echo "Create API path mapping"
    echo ""
    aws apigateway create-base-path-mapping \
    --base-path ${BASE_PATH} \
    --domain-name ${APP_API_CUSTOM} \
    --rest-api-id ${APP_API} \
    --stage ${APP_API_STAGE_NAME} \
    --region ${APP_REGION}
fi

APP_API_CUSTOM_ENDPOINT="https://${APP_API_CUSTOM}${APP_API_PATH}"

# ..............................................................................
CREATE_CNAME=$(aws route53 list-resource-record-sets --hosted-zone-id ${APP_DNS_HOSTED_ZONE} | \
jq -r ".ResourceRecordSets[] | select(.Name == \"${APP_API_CUSTOM}.\") | .Name")
if [ "${CREATE_CNAME}" = "" ]
then
    echo "Creating CNAME record for ${APP_API_CUSTOM}"
    echo ""
    echo "
    {
        \"Comment\": \"Custom domain for lambda fn ${APP_FN_NAME}\",
        \"Changes\": [
            {
                \"Action\": \"CREATE\",
                \"ResourceRecordSet\": {
                    \"Name\": \"${APP_API_CUSTOM}\",
                    \"Type\": \"CNAME\",
                    \"TTL\": 300,
                    \"ResourceRecords\": [
                        {
                            \"Value\": \"${APP_API_TARGET}\"
                        }
                    ]
                }
            }
        ]
    }
    " > ${APP_DIR}/change-resource-record-sets.json
    aws route53 change-resource-record-sets --hosted-zone-id ${APP_DNS_HOSTED_ZONE} \
    --change-batch file://${APP_DIR}/change-resource-record-sets.json
fi

# ..............................................................................

${APP_DIR}/config -env prod \
-key "APP_CERT_ARN" -value "${APP_CERT_ARN}" \
-key "APP_API_CUSTOM_ENDPOINT" -value "${APP_API_CUSTOM_ENDPOINT}" \
-key "APP_DNS_HOSTED_ZONE" -value "${APP_DNS_HOSTED_ZONE}" \
-update

