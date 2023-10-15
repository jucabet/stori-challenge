package dynamo

import (
	"context"
	"jucabet/stori-challenge/process-transactions/internal/domain/entities"
	"jucabet/stori-challenge/process-transactions/internal/infra/adapters/dynamo/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (adapter *DynamoDBAdapter) SaveFileInfo(fileChargeInfo *entities.FileCharge) error {
	_, err := adapter.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(string(adapter.tableName)),
		Item:      utils.MapFileChargeInfoEntityToDynamoDto(fileChargeInfo),
	})
	if err != nil {
		return err
	}

	return nil
}

func (adapter *DynamoDBAdapter) SaveTransaction(transaction *entities.Transaction) error {
	_, err := adapter.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(string(adapter.tableName)),
		Item:      utils.MapTransationEntityToDynamoDto(transaction),
	})
	if err != nil {
		return err
	}

	return nil
}
