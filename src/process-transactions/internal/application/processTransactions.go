package application

import (
	"errors"
	"jucabet/stori-challenge/process-transactions/internal/domain/dtos"
	"jucabet/stori-challenge/process-transactions/internal/domain/entities"
	"jucabet/stori-challenge/process-transactions/internal/domain/enums"
	"jucabet/stori-challenge/process-transactions/internal/domain/utils"
	"strconv"
	"sync"
	"time"
)

func (pt *ProcessTransaction) ProcessTransactions() error {
	files, err := pt.storageService.GetFilesByFolder(enums.IncomingFiles)
	if err != nil {
		utils.ErrorLog("ProcessTransactions", err)
		return err
	}

	if len(files) == 0 {
		return errors.New("FilesNotFound")
	}

	err = processFiles(pt, files)
	if err != nil {
		utils.ErrorLog("ProcessTransactions", err)
		return err
	}

	return nil
}

var processFiles = func(pt *ProcessTransaction, files []string) error {
	wg := sync.WaitGroup{}
	for _, filename := range files {
		txFileInfo, err := pt.storageService.GetFileContentByName(enums.IncomingFiles, filename)
		if err != nil {
			utils.ErrorLog("processFiles", err)
			return err
		}

		wg.Add(1)
		go func(
			filename string,
			txFileInfo []*dtos.TransactionsFileInfoDto,
		) {
			defer wg.Done()

			err := processFileTransactions(pt, filename, txFileInfo)
			if err != nil {
				utils.ErrorLog("processFiles", err)
			}

			err = moveFile(pt, filename)
			if err != nil {
				utils.ErrorLog("processFiles", err)
			}
		}(filename, txFileInfo)
	}

	wg.Wait()
	return nil
}

var processFileTransactions = func(
	pt *ProcessTransaction,
	filename string,
	txFileInfo []*dtos.TransactionsFileInfoDto,
) error {
	fileChargeID, err := registerFileChargeInfo(pt, filename)
	if err != nil {
		utils.ErrorLog("processTransactions", err)
		return err
	}

	wg := sync.WaitGroup{}
	for _, transaction := range txFileInfo {
		wg.Add(1)

		go func(tx *dtos.TransactionsFileInfoDto) {
			defer wg.Done()

			err := registerTransaction(pt, fileChargeID, tx)
			if err != nil {
				utils.ErrorLog("processTransactions", err)
			}
		}(transaction)
	}

	wg.Wait()

	err = pt.topicRepo.SendMessageToReport(fileChargeID)
	if err != nil {
		utils.ErrorLog("processTransactions", err)
		return err
	}

	return nil
}

var registerFileChargeInfo = func(
	pt *ProcessTransaction,
	filename string,
) (string, error) {
	fileChargeID := utils.GenerateRandomCode(15)
	chargeInfor := &entities.FileCharge{
		Type:     enums.FileCharge,
		ID:       fileChargeID,
		Date:     time.Now(),
		FileName: filename,
	}

	err := pt.databaseRepo.SaveFileInfo(chargeInfor)
	return fileChargeID, err
}

var registerTransaction = func(
	pt *ProcessTransaction,
	fileChargeID string,
	transaction *dtos.TransactionsFileInfoDto,
) error {
	id := strconv.Itoa(transaction.ID)
	tx := &entities.Transaction{
		Type:             enums.Transaction,
		ID:               id,
		Date:             transaction.Date,
		TransactionValue: transaction.Transaction,
		FileChargeID:     fileChargeID,
	}

	return pt.databaseRepo.SaveTransaction(tx)
}

var moveFile = func(
	pt *ProcessTransaction,
	filename string,
) error {
	src := &dtos.MoveFileDto{
		Folder:   enums.IncomingFiles,
		Filename: filename,
	}

	dest := &dtos.MoveFileDto{
		Folder:   enums.ProcessedFiles,
		Filename: filename,
	}

	return pt.storageRepo.MoveFile(src, dest)
}
