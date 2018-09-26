#!/bin/sh
set -e

export ETCDCTL_API=3

download-etcd-snapshot snapshot.db

if [ -e "snapshot.db" ]; then
	etcdctl snapshot restore --data-dir /etcd-data snapshot.db
	/usr/local/bin/etcd \
		--name s1 \
		--data-dir /etcd-data \
		--listen-client-urls http://0.0.0.0:2379 \
		--advertise-client-urls http://0.0.0.0:2379
fi

/usr/local/bin/etcd \
	--name s1 \
	--data-dir /etcd-data \
	--listen-client-urls http://0.0.0.0:2379 \
	--advertise-client-urls http://0.0.0.0:2379 \
	--listen-peer-urls http://0.0.0.0:2380 \
	--initial-advertise-peer-urls http://0.0.0.0:2380 \
	--initial-cluster s1=http://0.0.0.0:2380 \
	--initial-cluster-token tkn \
	--initial-cluster-state new
