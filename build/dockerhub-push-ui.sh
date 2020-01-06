#!/bin/bash
cp ./ui/Dockerfile ./Dockerfile
docker build -t munsy/art:ui .
rm ./Dockerfile
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push munsy/art:ui