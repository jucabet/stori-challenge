package balancereportstrategy

import (
	"fmt"
	"jucabet/stori-challenge/send-reports/internal/application/strategies/balanceReportStrategy/consts"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/entities"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"
	"strings"
	"sync"
)

func (br *BalanceReportStrategy) SendReport(reportData map[string]interface{}) error {
	contacts, err := br.databaseService.GetContacts()
	if err != nil {
		utils.ErrorLog("SendReport", err)
		return err
	}

	wg := sync.WaitGroup{}
	for _, contact := range contacts {
		wg.Add(1)

		go func(dest *entities.Contact) {
			defer wg.Done()
			err := br.mailerService.SendEmail(&dtos.SendEmailDto{
				Email:   dest.Email,
				Name:    dest.ContactName,
				Subject: consts.BalanceReportSubjet,
				Content: replaceTemplateValues(dest.ContactName, reportData),
			})
			if err != nil {
				utils.ErrorLog("SendReport", err)
			}
		}(contact)
	}

	wg.Wait()
	return nil
}

var replaceTemplateValues = func(contactName string, reportData map[string]interface{}) string {
	template := consts.BalanceReportContent
	template = strings.ReplaceAll(template, "{{totalTXJuly}}", fmt.Sprint(reportData["totalTXJuly"]))
	template = strings.ReplaceAll(template, "{{totalTXAugust}}", fmt.Sprint(reportData["totalTXAugust"]))
	template = strings.ReplaceAll(template, "{{totalBalance}}", fmt.Sprint(reportData["totalBalance"]))
	template = strings.ReplaceAll(template, "{{averageDebitAmount}}", fmt.Sprint(reportData["averageDebitAmount"]))
	template = strings.ReplaceAll(template, "{{averageCreditAmount}}", fmt.Sprint(reportData["averageCreditAmount"]))
	template = strings.ReplaceAll(template, "{{userName}}", contactName)

	return template
}
