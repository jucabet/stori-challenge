package factory

import (
	s3Adapter "jucabet/stori-challenge/process-transactions/internal/infra/adapters/s3"
	"jucabet/stori-challenge/process-transactions/internal/infra/deps/utils"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Adapter() (*s3Adapter.S3Adapter, error) {
	cfg, err := utils.NewAWSConfig(os.Getenv("AWS_REGION_PROJECT"))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	return s3Adapter.NewS3Adapter(client, os.Getenv("AWS_BUCKET_NAME")), nil
}
