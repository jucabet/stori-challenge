package application

import (
	"errors"
	"jucabet/stori-challenge/send-reports/internal/application/interfaces"
	balancereportstrategy "jucabet/stori-challenge/send-reports/internal/application/strategies/balanceReportStrategy"
	"jucabet/stori-challenge/send-reports/internal/domain/dtos"
	"jucabet/stori-challenge/send-reports/internal/domain/enums"
	"jucabet/stori-challenge/send-reports/internal/domain/utils"
)

func (sr *SendReport) SendReportHandler(data *dtos.SendReportDto) error {
	strategy := getReportStrategy(sr, data.ReportType)
	if strategy == nil {
		return errors.New("ReportTypeNotFound: " + string(data.ReportType))
	}

	reportData, err := strategy.BuildReport(data)
	if err != nil {
		return err
	}

	err = strategy.SendReport(reportData)
	if err != nil {
		return err
	}

	return nil
}

var getReportStrategy = func(sr *SendReport, reportType enums.ReportType) interfaces.IStategiesInterface {
	utils.Info("getReportStrategy", reportType)
	switch reportType {
	case enums.BalanceReport:
		return balancereportstrategy.NewBalanceReportStrategy(sr.databaseService, sr.mailerService)
	default:
		return nil
	}
}
