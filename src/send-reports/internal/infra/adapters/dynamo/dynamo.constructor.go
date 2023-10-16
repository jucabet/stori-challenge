package dynamo

import (
	"jucabet/stori-challenge/send-reports/internal/infra/interfaces"
)

type DynamoDBAdapter struct {
	client    interfaces.IDynamoDBInterface
	tableName string
}

func NewDynamoDBAdapter(
	client interfaces.IDynamoDBInterface,
	tableName string,
) *DynamoDBAdapter {
	return &DynamoDBAdapter{
		client:    client,
		tableName: tableName,
	}
}
