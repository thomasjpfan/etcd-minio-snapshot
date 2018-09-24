package main

import (
	"fmt"
	"log"
	"time"

	minio "github.com/minio/minio-go"
)

// DownloadClient wraps a minoClient
type DownloadClient struct {
	*minio.Client
}

// NewDownloadClientFromSpec creates a DownloadClient from spec
func NewDownloadClientFromSpec(spec Specification) (*DownloadClient, error) {
	cl, err := minio.New(
		spec.MinioEndpoint, spec.MinioAccessKey, spec.MinioSecretKey, false,
	)
	if err != nil {
		return nil, err
	}
	return &DownloadClient{cl}, nil
}

// Get downlaiods file from minio
func (d DownloadClient) Get(bucket, object, path string) error {

	bucketExist, err := d.BucketExists(bucket)
	if err != nil {
		return err
	}
	if !bucketExist {
		return fmt.Errorf("Bucket %s did not exist", bucket)
	}

	return d.FGetObject(bucket, object, path, minio.GetObjectOptions{})
}

// WaitForClient waits for minio instance be callable
func (d DownloadClient) WaitForClient(retries, interval int) error {
	_, err := d.ListBuckets()
	if err == nil {
		return nil
	}
	for i := 1; i < retries; i++ {
		log.Printf("Waiting for minio, %d remaining attempts...", i)
		time.Sleep(time.Millisecond * time.Duration(interval))
		_, err = d.ListBuckets()
		if err != nil {
			continue
		}
		return nil
	}
	return fmt.Errorf("Unable to connect to minio")
}
