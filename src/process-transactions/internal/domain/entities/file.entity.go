package entities

import (
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"
	"time"
)

type FileCharge struct {
	Type     enums.DBTypeRegisters
	ID       string
	Date     time.Time
	FileName string
}
