package main

import (
	"context"
	"jucabet/stori-challenge/process-transactions/internal/domain/utils"
	"jucabet/stori-challenge/process-transactions/internal/infra/deps/factory"
	lambdahandler "jucabet/stori-challenge/process-transactions/internal/infra/entrypoints/lambdaHandler"
	"jucabet/stori-challenge/process-transactions/internal/infra/enums"
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
	if os.Getenv("ENV") == string(enums.LOCAL) {
		result, err := lambdahandler.HandleRequest(context.Background())
		if err != nil {
			panic(err)
		}

		utils.Info("main", result)
		os.Exit(1)
	}

	lambda.Start(lambdahandler.HandleRequest)
}
