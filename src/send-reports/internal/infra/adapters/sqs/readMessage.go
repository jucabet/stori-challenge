package sqs

import (
	"context"
	"encoding/json"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (adapter *SQSAdapter) ReadMessage() (map[string]string, error) {
	result, err := adapter.getQueueUrl()
	if err != nil {
		utils.ErrorLog("ReadMessage", err)
		return nil, err
	}

	msgResult, err := adapter.client.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            result.QueueUrl,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   int32(60),
	})
	if err != nil {
		utils.ErrorLog("ReadMessage", err)
		return nil, err
	}

	if len(msgResult.Messages) == 0 {
		return map[string]string{}, nil
	}

	data := map[string]string{}
	json.Unmarshal([]byte(*msgResult.Messages[0].Body), &data)
	data["messageId"] = *msgResult.Messages[0].ReceiptHandle

	return data, nil
}
