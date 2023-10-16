package s3

import (
	"context"
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (adapter *S3Adapter) MoveFile(src, dest *dtos.MoveFileDto) error {
	err := adapter.copyFile(src, dest)
	if err != nil {
		return err
	}

	err = adapter.deleteFile(src.Folder, src.Filename)
	if err != nil {
		return err
	}

	return nil
}

func (adapter *S3Adapter) copyFile(src, dest *dtos.MoveFileDto) error {
	_, err := adapter.client.CopyObject(context.Background(), &s3.CopyObjectInput{
		Bucket:     aws.String(adapter.bucket),
		CopySource: aws.String(adapter.bucket + "/" + string(src.Folder) + src.Filename),
		Key:        aws.String(string(dest.Folder) + dest.Filename),
		ACL:        types.ObjectCannedACLBucketOwnerFullControl,
	})
	if err != nil {
		return err
	}

	return nil
}

func (adapter *S3Adapter) deleteFile(folder enums.StorageFolders, filename string) error {
	_, err := adapter.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(adapter.bucket),
		Key:    aws.String(string(folder) + filename),
	})
	if err != nil {
		return err
	}

	return nil
}
