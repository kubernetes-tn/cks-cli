
export $(cat .env | xargs);
aws cloudformation describe-stacks \
  --stack-name $CFN_STACK_NAME \
  --region $AWS_REGION \
  --query "Stacks[0].Outputs[?OutputKey=='Bucket'].OutputValue" --output text
