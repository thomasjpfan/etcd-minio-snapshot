package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	minio "github.com/minio/minio-go"
	"github.com/ory/dockertest"
	dc "github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/suite"
)

type DownloadTestSuite struct {
	suite.Suite
	Pool      *dockertest.Pool
	Resource  *dockertest.Resource
	DClient   *DownloadClient
	AccessKey string
	SecretKey string
	Endpoint  string
	TempDir   string
}

func TestDownloadUnitTestSuite(t *testing.T) {
	suite.Run(t, new(DownloadTestSuite))
}

func (s *DownloadTestSuite) SetupSuite() {
	pool, err := dockertest.NewPool("")
	s.Require().NoError(err)

	s.Pool = pool
	s.AccessKey = "MYACCESSKEY"
	s.SecretKey = "MYSECRETKEY"

	wd, err := os.Getwd()
	s.Require().NoError(err)
	testdata := filepath.Join(wd, "testdata")

	options := &dockertest.RunOptions{
		Repository: "minio/minio",
		Tag:        "RELEASE.2018-09-12T18-49-56Z",
		Cmd:        []string{"server", "/data"},
		PortBindings: map[dc.Port][]dc.PortBinding{
			"9000": []dc.PortBinding{{HostPort: "9000"}},
		},
		Env: []string{
			fmt.Sprintf("MINIO_ACCESS_KEY=%s", s.AccessKey), fmt.Sprintf("MINIO_SECRET_KEY=%s", s.SecretKey)},
		Mounts: []string{
			fmt.Sprintf("%s:/data", testdata),
		},
	}

	resource, err := pool.RunWithOptions(options)
	s.Resource = resource
	s.Require().NoError(err)

	s.Endpoint = fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp"))

	err = pool.Retry(func() error {
		minioClient, err := minio.New(
			s.Endpoint, s.AccessKey, s.SecretKey, false)
		_, err = minioClient.ListBuckets()
		if err != nil {
			return err
		}
		s.DClient = &DownloadClient{minioClient}
		return nil
	})
}

func (s *DownloadTestSuite) TearDownSuite() {
	s.Pool.Purge(s.Resource)
}

func (s *DownloadTestSuite) SetupTest() {
	dir, err := ioutil.TempDir("", "download")
	s.Require().NoError(err)
	s.TempDir = dir
}

func (s *DownloadTestSuite) TearDownTest() {
	os.RemoveAll(s.TempDir)
}

func (s *DownloadTestSuite) Test_Download_FileExists() {
	tmpfn := filepath.Join(s.TempDir, "dl.txt")
	err := s.DClient.Get("bucketpower", "hello.txt", tmpfn)
	s.Require().NoError(err)

	_, err = os.Stat(tmpfn)
	s.True(!os.IsNotExist(err))
}

func (s *DownloadTestSuite) Test_Download_FileDoesNotExists() {
	tmpfn := filepath.Join(s.TempDir, "dl.txt")
	err := s.DClient.Get("bucketpower", "hell.txt", tmpfn)
	s.Require().Error(err)

	_, err = os.Stat(tmpfn)
	s.True(os.IsNotExist(err))
}
func (s *DownloadTestSuite) Test_Download_FileBucketDoesNotExists() {

	tmpfn := filepath.Join(s.TempDir, "dl.txt")
	err := s.DClient.Get("bucketower", "hello.txt", tmpfn)
	s.Require().Error(err)

	_, err = os.Stat(tmpfn)
	s.True(os.IsNotExist(err))
}

func (s *DownloadTestSuite) Test_WaitForClient_Succeed() {
	err := s.DClient.WaitForClient(1, 100)
	s.NoError(err)
}

func (s *DownloadTestSuite) Test_WaitForClient_Fails() {
	minioClient, err := minio.New(
		s.Endpoint, s.AccessKey, "BAD_KEY", false)
	s.Require().NoError(err)

	dc := DownloadClient{minioClient}
	err = dc.WaitForClient(1, 100)

	s.Error(err)
}
