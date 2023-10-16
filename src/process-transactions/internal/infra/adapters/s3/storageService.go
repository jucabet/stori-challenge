package s3

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"
	"jucabet/stori-challenge/process-transactions/internal/infra/adapters/s3/utils"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (adapter *S3Adapter) GetFilesByFolder(folder enums.StorageFolders) ([]string, error) {
	itemOutput, err := adapter.client.ListObjects(context.Background(), &s3.ListObjectsInput{
		Bucket: aws.String(adapter.bucket),
		Prefix: aws.String(string(folder)),
	})
	if err != nil {
		return nil, err
	}

	files := []string{}
	for _, key := range itemOutput.Contents {
		path := string(*key.Key)
		filename := filepath.Base(path)
		if filename != string(folder) {
			files = append(files, filename)
		}
	}

	return files, nil
}

func (adapter *S3Adapter) GetFileContentByName(folder enums.StorageFolders, name string) ([]*dtos.TransactionsFileInfoDto, error) {
	itemOutput, err := adapter.client.GetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(adapter.bucket),
			Key:    aws.String(string(folder) + name),
		},
	)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(itemOutput.Body)

	var index = 0
	var response []*dtos.TransactionsFileInfoDto

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if len(record) != 3 {
			return nil, errors.New("InvalidFileFormat")
		}

		if index > 0 {
			tx, err := utils.MapRecordtoTXFileDto(record)
			if err != nil {
				return nil, err
			}

			response = append(response, tx)
		}

		index++
	}

	return response, nil
}
