package utils

import (
	"errors"
	"jucabet/stori-challenge/process-transactions/internal/domain/consts"
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"strconv"
	"time"
)

func MapRecordtoTXFileDto(record []string) (*dtos.TransactionsFileInfoDto, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, errors.New("InvalidIDFromFile: " + record[0])
	}

	date, err := time.Parse(consts.FileDateFormat, record[1])
	if err != nil {
		return nil, errors.New("InvalidDateFromFile: " + record[1])
	}

	tx, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return nil, errors.New("InvalidTXValueFromFile: " + record[1])
	}

	return &dtos.TransactionsFileInfoDto{
		ID:          id,
		Date:        date,
		Transaction: tx,
	}, nil
}
