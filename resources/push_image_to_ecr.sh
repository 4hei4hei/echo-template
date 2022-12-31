#!bin/bash
set -euo pipefail

# AWS profile config
AWS_PROFILE=hogehoge
REGION=ap-northeast-1

# Path to a target Dockerfile
IMAGE_PATH=$1
# Name you would like to give to an image
IMAGE_NAME=$2
# Tag you want to put on am image
IMAGE_TAG=$3

# Get AWS Account ID
AWS_ACCOUNT_ID=$(aws sts-get-caller-identity --profile ${AWS_PROFILE} | jq 'Account' | sed 's/"//g')

# Build image
docker images
docker build -t "${IMAGE_NAME}:${IMAGE_TAG}"

ECR_ENDPOINT="${AWS_ACCOUNT_ID}.dkr.ecr.${REGION}.amazonaws.com"
aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR_ENDPOINT}

# Get image ID
IAMGE_ID=$(docker images ${IMAGE_NAME} --format {{.ID}})

# Push image
docker tag "${IAMGE_ID} ${ECR_ENDPOINT}/${IMAGE_NAME}:${IMAGE_TAG}"
docker push ${ECR_ENDPOINT}/${IMAGE_NAME}:${IMAGE_TAG}"
