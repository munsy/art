#!/bin/bash
cp ./api/Dockerfile ./Dockerfile
docker build -t munsy/art:api .
rm ./Dockerfile
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push munsy/art:api