package s3

import (
	"jucabet/stori-challenge/process-transactions/internal/infra/interfaces"
)

type S3Adapter struct {
	client interfaces.IS3Interface
	bucket string
}

func NewS3Adapter(
	client interfaces.IS3Interface,
	bucket string,
) *S3Adapter {
	return &S3Adapter{
		client: client,
		bucket: bucket,
	}
}
