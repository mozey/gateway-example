#!/usr/bin/env bash

aws lambda get-function --function-name mozey-gateway

# Create VPC and Subnets

# Add function to VPC
aws lambda update-function-configuration --function-name mozey-gateway --vpc-config SubnetIds=subnet-xxx,subnet-xxx,subnet-xxx,SecurityGroupIds=sg-xxx

# Enable Outgoing Internet Access within VPC
# https://medium.com/@philippholly/aws-lambda-enable-outgoing-internet-access-within-vpc-8dd250e11e12
# https://medium.com/financial-engines-techblog/aws-lambdas-with-a-static-outgoing-ip-5174a1e70245

# Create a NAT Gateway in a subnet


# Give the lambda function internet access

# TODO
