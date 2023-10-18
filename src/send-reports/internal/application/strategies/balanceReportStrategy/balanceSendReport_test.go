package balancereportstrategy

import (
	"errors"
	"fmt"
	"jucabet/stori-challenge/send-reports/internal/application/strategies/balanceReportStrategy/consts"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/entities"
	"jucabet/stori-challenge/send-reports/internal/domain/enums"
	"jucabet/stori-challenge/send-reports/mocks"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendReport(t *testing.T) {
	databaseService := new(mocks.IDatabaseService)
	mailerService := new(mocks.IMailerService)

	strategy := NewBalanceReportStrategy(databaseService, mailerService)

	contacts := []*entities.Contact{
		{
			Type:        enums.Contact,
			ID:          "1",
			ContactName: "Juanca",
			Email:       "mock@mock.com",
		},
	}

	response := map[string]interface{}{
		"totalTXJuly":         2,
		"totalTXAugust":       2,
		"totalBalance":        "39.74",
		"averageDebitAmount":  "-15.38",
		"averageCreditAmount": "35.25",
	}

	mailData := &dtos.SendEmailDto{
		Email:   contacts[0].Email,
		Name:    contacts[0].ContactName,
		Subject: consts.BalanceReportSubjet,
		Content: "HTML",
	}

	mockReplaceTemplateValues := replaceTemplateValues
	replaceTemplateValues = func(contactName string, reportData map[string]interface{}) string {
		return "HTML"
	}

	t.Run("error getting contacts", func(t *testing.T) {
		databaseService.On("GetContacts").
			Return(contacts, errors.New("mockErrorGetContacts")).
			Once()

		err := strategy.SendReport(response)
		assert.Error(t, err, "mockErrorGetContacts")
	})

	t.Run("error sending mail", func(t *testing.T) {
		databaseService.On("GetContacts").
			Return(contacts, nil).
			Once()

		mailerService.On("SendEmail", mailData).
			Return(errors.New("mockErrorSendEmail")).
			Once()

		err := strategy.SendReport(response)
		assert.NoError(t, err)
	})

	t.Run("sucess", func(t *testing.T) {
		databaseService.On("GetContacts").
			Return(contacts, nil).
			Once()

		mailerService.On("SendEmail", mailData).
			Return(nil).
			Once()

		err := strategy.SendReport(response)
		assert.NoError(t, err)
	})

	replaceTemplateValues = mockReplaceTemplateValues
}

func TestReplaceTemplateValues(t *testing.T) {
	response := map[string]interface{}{
		"totalTXJuly":         2,
		"totalTXAugust":       2,
		"totalBalance":        "39.74",
		"averageDebitAmount":  "-15.38",
		"averageCreditAmount": "35.25",
	}

	contactName := "Juanca"

	template := consts.BalanceReportContent
	template = strings.ReplaceAll(template, "{{totalTXJuly}}", fmt.Sprint(response["totalTXJuly"]))
	template = strings.ReplaceAll(template, "{{totalTXAugust}}", fmt.Sprint(response["totalTXAugust"]))
	template = strings.ReplaceAll(template, "{{totalBalance}}", fmt.Sprint(response["totalBalance"]))
	template = strings.ReplaceAll(template, "{{averageDebitAmount}}", fmt.Sprint(response["averageDebitAmount"]))
	template = strings.ReplaceAll(template, "{{averageCreditAmount}}", fmt.Sprint(response["averageCreditAmount"]))
	template = strings.ReplaceAll(template, "{{userName}}", contactName)

	newTemplate := replaceTemplateValues(contactName, response)
	assert.Equal(t, newTemplate, template)
}
