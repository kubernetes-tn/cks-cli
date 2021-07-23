#!/bin/bash -xe

# This is a standard Bash trick to get the directory the script is running in.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )";
export $(cat .env | xargs);
# Deploy
aws cloudformation deploy --stack-name $CFN_STACK_NAME \
   --template-file deploy/cloudformation/template.yaml \
   --parameter-overrides \
      Environment=${CFN_ENVIRONMENT} \
      SSLCertificateARN=${CFN_SSL_CERT_ARN} \
      Hosts=${CFN_HOSTS} \
   --region $AWS_REGION \
   --capabilities CAPABILITY_IAM \
   --tags Env=$CFN_ENVIRONMENT ServiceName=cks-cli;

