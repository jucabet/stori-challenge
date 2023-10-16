package balancereportstrategy

import (
	"fmt"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/entities"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"
	"time"
)

func (br *BalanceReportStrategy) BuildReport(data *dtos.SendReportDto) (map[string]interface{}, error) {
	transactions, err := br.databaseService.GetTransactionsByFileChargeID(data.FileChargeID)
	if err != nil {
		utils.ErrorLog("BuildReport", err)
		return nil, err
	}

	repotData := getTransactionsData(transactions)

	return repotData, nil
}

var getTransactionsData = func(transactions []*entities.Transaction) map[string]interface{} {
	var totalTXJuly, totalTXAugust int
	var DebitAmount, CreditAmount, totalBalance float64
	// variables for count number of credict and debit transactions for average
	var countDebitAmunt, countCreditAmount int

	for _, transaction := range transactions {
		identifyMonthTransaction(transaction, &totalTXJuly, &totalTXAugust)
		balanceWithTransaction(
			transaction,
			&DebitAmount,
			&CreditAmount,
			&totalBalance,
			&countDebitAmunt,
			&countCreditAmount,
		)
	}

	return map[string]interface{}{
		"totalTXJuly":         totalTXJuly,
		"totalTXAugust":       totalTXAugust,
		"totalBalance":        fmt.Sprintf("%.2f", totalBalance),
		"averageDebitAmount":  fmt.Sprintf("%.2f", (DebitAmount / float64(countDebitAmunt))),
		"averageCreditAmount": fmt.Sprintf("%.2f", (CreditAmount / float64(countCreditAmount))),
	}
}

var identifyMonthTransaction = func(
	transaction *entities.Transaction,
	totalTXJuly,
	totalTXAugust *int,
) {
	if transaction.Date.Month() == time.July {
		(*totalTXJuly)++
	}

	if transaction.Date.Month() == time.August {
		(*totalTXAugust)++
	}
}

var balanceWithTransaction = func(
	transaction *entities.Transaction,
	DebitAmount,
	CreditAmount,
	totalBalance *float64,
	countDebitAmunt,
	countCreditAmount *int,
) {
	if transaction.TransactionValue < 0 {
		(*DebitAmount) = (*DebitAmount) + transaction.TransactionValue
		(*countDebitAmunt)++
	}

	if transaction.TransactionValue >= 0 {
		(*CreditAmount) = (*CreditAmount) + transaction.TransactionValue
		(*countCreditAmount)++
	}

	// Algorithmic sum
	(*totalBalance) = (*totalBalance) + transaction.TransactionValue
}
