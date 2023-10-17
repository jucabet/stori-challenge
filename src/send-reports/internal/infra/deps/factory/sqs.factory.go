package factory

import (
	sqsAdapter "jucabet/stori-challenge/send-reports/internal/infra/adapters/sqs"
	"jucabet/stori-challenge/send-reports/internal/infra/deps/utils"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func NewSQSAdapter() (*sqsAdapter.SQSAdapter, error) {
	cfg, err := utils.NewAWSConfig(os.Getenv("AWS_REGION_PROJECT"))
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(cfg)

	return sqsAdapter.NewSQSAdapter(client, os.Getenv("AWS_SQS_NAME")), nil
}
