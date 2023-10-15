package sqs

import (
	"context"
	"encoding/json"
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"
	"jucabet/stori-challenge/process-transactions/internal/domain/utils"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (adapter *SQSAdapter) SendMessageToReport(fileChargeID string) error {
	result, err := adapter.client.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: &adapter.queueName,
	})
	if err != nil {
		return err
	}

	body, _ := json.Marshal(map[string]string{
		"reportType":   string(enums.BalanceReport),
		"fileChargeId": fileChargeID,
	})

	messageOutput, err := adapter.client.SendMessage(context.Background(), &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{
			"reportType": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(enums.BalanceReport)),
			},
			"fileChargeId": {
				DataType:    aws.String("String"),
				StringValue: aws.String(fileChargeID),
			},
		},
		MessageBody: aws.String(string(body)),
		QueueUrl:    result.QueueUrl,
	})
	if err != nil {
		return err
	}

	utils.Info("SendMessageToReport", messageOutput)

	return nil
}
