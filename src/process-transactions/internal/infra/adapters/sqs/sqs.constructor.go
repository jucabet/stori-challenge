package sqs

import (
	"jucabet/stori-challenge/process-transactions/internal/infra/interfaces"
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
