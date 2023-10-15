package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type IDynamoDBInterface interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
