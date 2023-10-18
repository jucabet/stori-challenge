package balancereportstrategy

import (
	"errors"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/entities"
	"jucabet/stori-challenge/send-reports/internal/domain/enums"
	"jucabet/stori-challenge/send-reports/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildReport(t *testing.T) {
	databaseService := new(mocks.IDatabaseService)
	mailerService := new(mocks.IMailerService)

	strategy := NewBalanceReportStrategy(databaseService, mailerService)
	data := &dtos.SendReportDto{
		ReportType:   enums.BalanceReport,
		FileChargeID: "1234abc",
	}
	sendData := map[string]interface{}{
		"total": 10,
	}
	transactions := []*entities.Transaction{}

	mockGetTransactionsData := getTransactionsData
	getTransactionsData = func(transactions []*entities.Transaction) map[string]interface{} {
		return sendData
	}

	t.Run("error in Get Transactions By File Charge ID", func(t *testing.T) {
		databaseService.On("GetTransactionsByFileChargeID", data.FileChargeID).
			Return(nil, errors.New("mockErrorGetTransactionsByFileChargeID")).
			Once()

		_, err := strategy.BuildReport(data)
		assert.Error(t, err, "mockErrorGetTransactionsByFileChargeID")
	})

	t.Run("sucess", func(t *testing.T) {
		databaseService.On("GetTransactionsByFileChargeID", data.FileChargeID).
			Return(transactions, nil).
			Once()

		received, err := strategy.BuildReport(data)
		assert.NoError(t, err)
		assert.Equal(t, received, sendData)
	})

	getTransactionsData = mockGetTransactionsData
}

func TestGetTransactionsData(t *testing.T) {
	july, _ := time.Parse("1/2", "7/12")
	august, _ := time.Parse("1/2", "8/12")
	fileChargeID := "1234abc"

	transactions := []*entities.Transaction{
		{
			Type:             enums.Transaction,
			ID:               "1",
			Date:             july,
			TransactionValue: 60.5,
			FileChargeID:     fileChargeID,
		},
		{
			Type:             enums.Transaction,
			ID:               "2",
			Date:             july,
			TransactionValue: -10.3,
			FileChargeID:     fileChargeID,
		},
		{
			Type:             enums.Transaction,
			ID:               "3",
			Date:             august,
			TransactionValue: -20.46,
			FileChargeID:     fileChargeID,
		},
		{
			Type:             enums.Transaction,
			ID:               "4",
			Date:             august,
			TransactionValue: 10,
			FileChargeID:     fileChargeID,
		},
	}

	responseExpected := map[string]interface{}{
		"totalTXJuly":         2,
		"totalTXAugust":       2,
		"totalBalance":        "39.74",
		"averageDebitAmount":  "-15.38",
		"averageCreditAmount": "35.25",
	}

	responseReceived := getTransactionsData(transactions)

	assert.Equal(t, responseExpected, responseReceived)
}
