package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func TestParserUnitTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}

func (s *ParserTestSuite) TearDownTest() {
	os.Clearenv()
}

func (s *ParserTestSuite) Test_Wow() {

	accessKey := "minio_access_key"
	secretKey := "secret_key"
	minioEndpoint := "minio:9090"
	etcdSnapshotBucket := "a_bucket"
	etcdObjectName := "object_name"

	os.Setenv("EMS_MINIO_ACCESS_KEY", accessKey)
	os.Setenv("EMS_MINIO_SECRET_KEY", secretKey)
	os.Setenv("EMS_MINIO_ENDPOINT", minioEndpoint)
	os.Setenv("EMS_ETCD_SNAPSHOT_BUCKET", etcdSnapshotBucket)
	os.Setenv("EMS_ETCD_OBJECT_NAME", etcdObjectName)

	spec, err := ParseENV()
	s.Require().NoError(err)

	s.Equal(accessKey, spec.MinioAccessKey)
	s.Equal(secretKey, spec.MinioSecretKey)
	s.Equal(minioEndpoint, spec.MinioEndpoint)
	s.Equal(etcdSnapshotBucket, spec.EtcdSnapshotBucket)
	s.Equal(etcdObjectName, spec.EtcdObjectName)
}
