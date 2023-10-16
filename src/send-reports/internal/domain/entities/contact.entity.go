package entities

import "jucabet/stori-challenge/send-reports/internal/domain/enums"

type Contact struct {
	Type        enums.DBTypeRegisters
	ID          string
	ContactName string
	Email       string
}
