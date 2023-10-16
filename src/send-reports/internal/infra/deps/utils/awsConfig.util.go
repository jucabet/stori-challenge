package utils

import (
	"context"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"
	"jucabet/stori-challenge/send-reports/internal/infra/enums"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewAWSConfig(region string) (aws.Config, error) {
	var cfg aws.Config
	var err error

	if os.Getenv("ENV") == string(enums.LOCAL) || os.Getenv("ENV") == string(enums.DOCKER) {
		customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           os.Getenv("URL_ENDPOINT_LOCALSTACK"),
				SigningRegion: region,
			}, nil
		})

		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
			config.WithEndpointResolverWithOptions(customResolver),
		)
		if err != nil {
			utils.ErrorLog("NewAWSConfig", err)
			return cfg, err
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			utils.ErrorLog("NewAWSConfig", err)
			return cfg, err
		}
	}

	return cfg, nil
}
