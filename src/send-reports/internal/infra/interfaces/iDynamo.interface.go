package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type IDynamoDBInterface interface {
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}
