package utils

import (
	"errors"
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"strconv"
	"strings"
	"time"
)

func MapRecordtoTXFileDto(record []string) (*dtos.TransactionsFileInfoDto, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return nil, errors.New("InvalidIDFromFile: " + record[0])
	}

	dateSplit := strings.Split(record[1], "/")
	if len(dateSplit) != 2 {
		return nil, errors.New("InvalidDateFromFile: " + record[1])
	}

	month, err := strconv.Atoi(dateSplit[0])
	if err != nil {
		return nil, errors.New("InvalidDateFromFile: " + record[1])
	}

	day, err := strconv.Atoi(dateSplit[1])
	if err != nil {
		return nil, errors.New("InvalidDateFromFile: " + record[1])
	}

	date := time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.Local)

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
