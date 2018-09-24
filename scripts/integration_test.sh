#!/bin/sh
set -e

export ETCDCTL_API=3

snap_server="etcd_snap"
no_snap_server="etcd_no_snap"

RETRIES=20
until nc -z $snap_server 2379 >/dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
	echo "Waiting for ${snap_server} server, $((RETRIES--)) remaining attempts..."
	sleep 1
done

value=$(etcdctl --endpoints ${snap_server}:2379 get --print-value-only hello)
[ $value ]

RETRIES=20
until nc -z $no_snap_server 2379 >/dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
	echo "Waiting for ${no_snap_server} server, $((RETRIES--)) remaining attempts..."
	sleep 1
done

value=$(etcdctl --endpoints ${no_snap_server}:2379 get --print-value-only hello)
[ ! $value ]
