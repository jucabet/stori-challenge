package entities

import (
	"jucabet/stori-challenge/send-reports/internal/domain/enums"
	"time"
)

type Transaction struct {
	Type             enums.DBTypeRegisters
	ID               string
	Date             time.Time
	TransactionValue float64
	FileChargeID     string
}
