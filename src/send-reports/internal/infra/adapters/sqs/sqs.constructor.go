package sqs

import (
	"context"
	"jucabet/stori-challenge/send-reports/internal/infra/interfaces"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSAdapter struct {
	client    interfaces.ISQSInterface
	queueName string
}

func NewSQSAdapter(client interfaces.ISQSInterface, queueName string) *SQSAdapter {
	return &SQSAdapter{
		client:    client,
		queueName: queueName,
	}
}

func (adapter *SQSAdapter) getQueueUrl() (*sqs.GetQueueUrlOutput, error) {
	result, err := adapter.client.GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
		QueueName: &adapter.queueName,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
