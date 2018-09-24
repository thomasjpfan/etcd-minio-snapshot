#!/bin/bash

docker_hub_name="thomasjpfan/etcd-snapshot"

master_image="${docker_hub_name}:master"
latest_image="${docker_hub_name}:latest"
release_image="${docker_hub_name}:${DOCKER_RELEASE_TAG}"

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

docker tag "$master_image" "$latest_image"
docker tag "$master_image" "$release_image"

docker push "$latest_image"
docker push "$release_image"