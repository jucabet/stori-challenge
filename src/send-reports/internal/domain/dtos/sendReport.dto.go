package dtos

import "jucabet/stori-challenge/send-reports/internal/domain/enums"

type SendReportDto struct {
	ReportType   enums.ReportType
	FileChargeID string
}
