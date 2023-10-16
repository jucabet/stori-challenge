package sqs

import (
	"context"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func (adapter *SQSAdapter) DeleteMessage(mesageID string) error {
	result, err := adapter.getQueueUrl()
	if err != nil {
		utils.ErrorLog("DeleteMessage", err)
		return err
	}

	dMInput := &sqs.DeleteMessageInput{
		QueueUrl:      result.QueueUrl,
		ReceiptHandle: aws.String(mesageID),
	}

	_, err = adapter.client.DeleteMessage(context.TODO(), dMInput)
	if err != nil {
		utils.ErrorLog("DeleteMessage", err)
		return err
	}

	return nil
}
