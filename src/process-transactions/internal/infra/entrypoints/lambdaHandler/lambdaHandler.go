package lambdahandler

import (
	"context"
	"jucabet/stori-challenge/process-transactions/internal/application"
	"jucabet/stori-challenge/process-transactions/internal/infra/deps/factory"
)

func HandleRequest(ctx context.Context) (string, error) {
	usecase, err := buildUsecase()
	if err != nil {
		return "", err
	}

	err = usecase.ProcessTransactions()
	if err != nil {
		return "", err
	}

	return "sucess", err
}

var buildUsecase = func() (*application.ProcessTransaction, error) {
	s3Adapter, err := factory.NewS3Adapter()
	if err != nil {
		return nil, err
	}

	dynamoAdapter, err := factory.NewDynamoDBAdapter()
	if err != nil {
		return nil, err
	}

	sqsAdapter, err := factory.NewSQSAdapter()
	if err != nil {
		return nil, err
	}

	usecase := application.NewProcessTransactions(
		s3Adapter,
		s3Adapter,
		dynamoAdapter,
		sqsAdapter,
	)

	return usecase, nil
}
