#!/bin/bash
docker build -t munsy/art .
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push munsy/art