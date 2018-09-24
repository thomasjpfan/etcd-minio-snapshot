TAG?=master
TRAVIS_COMMIT?=master

build:
	go build -mod vendor -o etcd-minio-snapshot

unittest:
	go test ./... --run UnitTest

image:
	docker image build -t thomasjpfan/etcd-minio-snapshot:$(TAG) \
	--label "org.opencontainers.image.revision=$(TRAVIS_COMMIT)" .
