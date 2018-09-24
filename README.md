# etcd minio snapshot

Before etcd starts up, a minio instance is queried for a snapshot to initialize the etcd db. If the minio instance does not have a snapshot, a new etcd instance is started.

## Environment Variables

- `EMS_MINIO_ACCESS_KEY`: minio access key
- `EMS_MINIO_SECRET_KEY`: minio secret key
- `EMS_MINIO_ENDPOINT`: endpoint of minio
- `EMS_ETCD_SNAPSHOT_BUCKET`: bucket snapshot is on
- `EMS_ETCD_OBJECT_NAME`: name of object in bucket
