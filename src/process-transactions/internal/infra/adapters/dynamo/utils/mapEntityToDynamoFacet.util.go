package utils

import (
	"fmt"
	"jucabet/stori-challenge/process-transactions/internal/domain/consts"
	"jucabet/stori-challenge/process-transactions/internal/domain/entities"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func MapFileChargeInfoEntityToDynamoDto(fileChargeInfo *entities.FileCharge) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"type":     &types.AttributeValueMemberS{Value: string(fileChargeInfo.Type)},
		"id":       &types.AttributeValueMemberS{Value: fileChargeInfo.ID},
		"date":     &types.AttributeValueMemberS{Value: fileChargeInfo.Date.Format(consts.DatabaseDateFormat)},
		"fileName": &types.AttributeValueMemberS{Value: fileChargeInfo.FileName},
	}
}

func MapTransationEntityToDynamoDto(transaction *entities.Transaction) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"type":             &types.AttributeValueMemberS{Value: string(transaction.Type)},
		"id":               &types.AttributeValueMemberS{Value: transaction.ID},
		"date":             &types.AttributeValueMemberS{Value: transaction.Date.Format(consts.DatabaseDateFormat)},
		"transactionValue": &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", transaction.TransactionValue)},
	}
}
