package utils

import (
	"context"
	"jucabet/stori-challenge/process-transactions/internal/domain/utils"
	"jucabet/stori-challenge/process-transactions/internal/infra/enums"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/ratelimit"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
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
			config.WithRetryer(func() aws.Retryer {
				return retry.NewStandard(func(so *retry.StandardOptions) {
					so.RateLimiter = ratelimit.NewTokenRateLimit(1000000)
				})
			}),
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
