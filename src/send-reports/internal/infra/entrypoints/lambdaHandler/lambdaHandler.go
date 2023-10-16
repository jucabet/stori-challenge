package lambdahandler

import (
	"context"
	"jucabet/stori-challenge/send-reports/internal/application"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/enums"
	"jucabet/stori-challenge/send-reports/internal/infra/adapters/sqs"
	"jucabet/stori-challenge/send-reports/internal/infra/deps/factory"
)

func Handler(sqsAdapter *sqs.SQSAdapter) func(ctx context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		message, err := sqsAdapter.ReadMessage()
		if err != nil {
			return "", err
		}

		usecase, err := buildUsecase()
		if err != nil {
			return "", err
		}

		err = usecase.SendReportHandler(&dtos.SendReportDto{
			ReportType:   enums.ReportType(message["reportType"]),
			FileChargeID: message["fileChargeId"],
		})
		if err != nil {
			return "", err
		}

		err = sqsAdapter.DeleteMessage(message["messageId"])
		if err != nil {
			return "", err
		}

		return "sucess", nil
	}
}

var buildUsecase = func() (*application.SendReport, error) {
	dynamoAdapter, err := factory.NewDynamoDBAdapter()
	if err != nil {
		return nil, err
	}

	mailer, err := factory.NewMailerAdapter()
	if err != nil {
		return nil, err
	}

	usecase := application.NewSendReport(
		dynamoAdapter,
		mailer,
	)

	return usecase, nil
}
