#!/bin/bash

docker_hub_name="${DOCKER_USERNAME}/etcd-snapshot"

master_image="${docker_hub_name}:master"
latest_image="${docker_hub_name}:latest"
release_image="${docker_hub_name}:${TRAVIS_COMMIT}"

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

docker tag "$master_image" "$latest_image"
docker tag "$master_image" "$release_image"

docker push "$latest_image"
docker push "$release_image"
