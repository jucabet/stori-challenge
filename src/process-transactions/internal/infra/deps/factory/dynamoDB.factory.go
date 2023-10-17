package factory

import (
	"jucabet/stori-challenge/process-transactions/internal/infra/adapters/dynamo"
	"jucabet/stori-challenge/process-transactions/internal/infra/deps/utils"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBAdapter() (*dynamo.DynamoDBAdapter, error) {
	cfg, err := utils.NewAWSConfig(os.Getenv("AWS_REGION_PROJECT"))
	if err != nil {
		return nil, err
	}

	client := dynamodb.NewFromConfig(cfg)

	return dynamo.NewDynamoDBAdapter(client, os.Getenv("AWS_DYNAMO_TABLE_NAME")), nil
}
