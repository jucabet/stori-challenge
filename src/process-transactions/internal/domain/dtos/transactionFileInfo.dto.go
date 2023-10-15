package dtos

import "time"

type TransactionsFileInfoDto struct {
	ID          int
	Date        time.Time
	Transaction float64
}
