#!/bin/bash

# Define variables
DOCKER_IMAGE_NAME="secret-sidecar"
DOCKER_IMAGE_TAG="latest"
ACR_NAME="acrasgardeomainrnd001.azurecr.io"

# Build Docker image
docker build -t $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG .

# Tag the Docker image for ACR
docker tag $DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG $ACR_NAME/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG

# Push the Docker image to ACR
docker push $ACR_NAME/$DOCKER_IMAGE_NAME:$DOCKER_IMAGE_TAG
