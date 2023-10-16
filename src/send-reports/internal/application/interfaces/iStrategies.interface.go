package interfaces

import "jucabet/stori-challenge/send-reports/internal/domain/dtos"

type IStategiesInterface interface {
	BuildReport(data *dtos.SendReportDto) (map[string]interface{}, error)
	SendReport(reportData map[string]interface{}) error
}
