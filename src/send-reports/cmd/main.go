package main

import (
	"context"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"
	"jucabet/stori-challenge/send-reports/internal/infra/deps/factory"
	lambdahandler "jucabet/stori-challenge/send-reports/internal/infra/entrypoints/lambdaHandler"
	"jucabet/stori-challenge/send-reports/internal/infra/enums"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	secretAdapter := factory.NewSecretAdapter()
	err := secretAdapter.LoadSecrets()
	if err != nil {
		panic(err)
	}

	runHandler()
}

func runHandler() {
	sqsAdapter, err := factory.NewSQSAdapter()
	if err != nil {
		panic(err)
	}

	if os.Getenv("ENV") == string(enums.LOCAL) || os.Getenv("ENV") == string(enums.DOCKER) {
		handlerFunc := lambdahandler.Handler(sqsAdapter)

		result, err := handlerFunc(context.Background())
		if err != nil {
			panic(err)
		}

		utils.Info("main", result)
		os.Exit(1)
	}

	lambda.Start(lambdahandler.Handler(sqsAdapter))
}
