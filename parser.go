package main

import (
	"github.com/kelseyhightower/envconfig"
)

// Specification are env variables used by ems
// split_words is used by envconfig
type Specification struct {
	MinioAccessKey     string `split_words:"true"`
	MinioSecretKey     string `split_words:"true"`
	MinioEndpoint      string `split_words:"true"`
	EtcdSnapshotBucket string `split_words:"true"`
	EtcdObjectName     string `split_words:"true"`
}

// ParseENV pareses env for variables
func ParseENV() (*Specification, error) {
	s := Specification{}

	err := envconfig.Process("EMS", &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
