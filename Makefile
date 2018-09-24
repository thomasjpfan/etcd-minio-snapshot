TAG?=master

build:
	go build -mod vendor -o etcd-minio-snapshot

unittest:
	go test ./... --run UnitTest

image:
	docker image build -t thomasjpfan/etcd-minio-snapshot:$(TAG) .
