package application

import (
	"errors"
	"jucabet/stori-challenge/send-reports/internal/application/interfaces"
	balancereportstrategy "jucabet/stori-challenge/send-reports/internal/application/strategies/balanceReportStrategy"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/enums"
	"jucabet/stori-challenge/send-reports/internal/domain/ports"
	"jucabet/stori-challenge/send-reports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendReportHandler(t *testing.T) {
	databaseService := new(mocks.IDatabaseService)
	mailerService := new(mocks.IMailerService)
	strategy := new(mocks.IStategiesInterface)

	usecase := NewSendReport(databaseService, mailerService)

	mockGetReportStrategy := getReportStrategy
	data := &dtos.SendReportDto{
		ReportType:   enums.BalanceReport,
		FileChargeID: "1234abc",
	}

	t.Run("no strategy found", func(t *testing.T) {
		getReportStrategy = func(
			sr *SendReport,
			reportType enums.ReportType,
		) interfaces.IStategiesInterface {
			return nil
		}

		err := usecase.SendReportHandler(data)
		assert.Error(t, err, "ReportTypeNotFound: balance")
	})

	t.Run("error building report", func(t *testing.T) {
		getReportStrategy = func(
			sr *SendReport,
			reportType enums.ReportType,
		) interfaces.IStategiesInterface {
			return strategy
		}

		strategy.On("BuildReport", data).
			Return(nil, errors.New("mockErrorBuildReport")).
			Once()

		err := usecase.SendReportHandler(data)
		assert.Error(t, err, "mockErrorBuildReport")
	})

	t.Run("error sending report", func(t *testing.T) {
		sendData := map[string]interface{}{
			"total": 10,
		}

		getReportStrategy = func(
			sr *SendReport,
			reportType enums.ReportType,
		) interfaces.IStategiesInterface {
			return strategy
		}

		strategy.On("BuildReport", data).
			Return(sendData, nil).
			Once()

		strategy.On("SendReport", sendData).
			Return(errors.New("mockErrorSendReport")).
			Once()

		err := usecase.SendReportHandler(data)
		assert.Error(t, err, "mockErrorSendReport")
	})

	t.Run("sucess", func(t *testing.T) {
		sendData := map[string]interface{}{
			"total": 10,
		}

		getReportStrategy = func(
			sr *SendReport,
			reportType enums.ReportType,
		) interfaces.IStategiesInterface {
			return strategy
		}

		strategy.On("BuildReport", data).
			Return(sendData, nil).
			Once()

		strategy.On("SendReport", sendData).
			Return(nil).
			Once()

		err := usecase.SendReportHandler(data)
		assert.NoError(t, err)
	})

	getReportStrategy = mockGetReportStrategy
}

func TestGetReportStrategy(t *testing.T) {
	databaseService := new(mocks.IDatabaseService)
	mailerService := new(mocks.IMailerService)

	usecase := NewSendReport(databaseService, mailerService)

	mockNewBalanceReportStrategy := balancereportstrategy.NewBalanceReportStrategy
	balancereportstrategy.NewBalanceReportStrategy = func(
		databaseService ports.IDatabaseService,
		mailerService ports.IMailerService,
	) interfaces.IStategiesInterface {
		return &balancereportstrategy.BalanceReportStrategy{}
	}

	t.Run("balance report", func(t *testing.T) {
		strategy := getReportStrategy(usecase, enums.BalanceReport)
		assert.NotNil(t, strategy)
	})

	t.Run("no strategy", func(t *testing.T) {
		strategy := getReportStrategy(usecase, "")
		assert.Nil(t, strategy)
	})

	balancereportstrategy.NewBalanceReportStrategy = mockNewBalanceReportStrategy
}
