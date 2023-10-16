package dynamo

import (
	"context"
	"jucabet/stori-challenge/send-reports/internal/domain/entities"
	domainEnums "jucabet/stori-challenge/send-reports/internal/domain/enums"
	"jucabet/stori-challenge/send-reports/internal/infra/adapters/dynamo/enums"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (adapter *DynamoDBAdapter) GetContacts() ([]*entities.Contact, error) {
	output, err := adapter.client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(adapter.tableName),
		KeyConditionExpression: aws.String("#type = :type"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":type": &types.AttributeValueMemberS{Value: string(domainEnums.Contact)},
		},
		ExpressionAttributeNames: map[string]string{
			"#type": "type",
		},
	})
	if err != nil {
		return nil, err
	}

	var contacts = []*entities.Contact{}
	err = attributevalue.UnmarshalListOfMaps(output.Items, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (adapter *DynamoDBAdapter) GetTransactionsByFileChargeID(fileChargeID string) ([]*entities.Transaction, error) {
	output, err := adapter.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(adapter.tableName),
		IndexName:              aws.String(string(enums.GSIFileChargeID)),
		KeyConditionExpression: aws.String("fileChargeId = :fileChargeId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":fileChargeId": &types.AttributeValueMemberS{Value: fileChargeID},
		},
	})
	if err != nil {
		return nil, err
	}

	var transactions = []*entities.Transaction{}
	err = attributevalue.UnmarshalListOfMaps(output.Items, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
