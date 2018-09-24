#!/bin/sh

docker container run -v $PWD/scripts:/scripts \
	--network etcd-minio-snapshot_minio \
	gcr.io/etcd-development/etcd:v3.3.9 /scripts/integration_test.sh
