package localsecrets

import (
	"jucabet/stori-challenge/process-transactions/internal/domain/utils"

	"github.com/joho/godotenv"
)

type LocalSecretAdapter struct{}

func NewLocalSecretAdapter() *LocalSecretAdapter {
	return &LocalSecretAdapter{}
}

func (adapter *LocalSecretAdapter) LoadSecrets() error {
	err := godotenv.Load()
	if err != nil {
		utils.ErrorLog("LoadSecret", err)
		return err
	}

	return nil
}
