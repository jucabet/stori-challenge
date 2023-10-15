package prodsecrets

type ProdSecretAdapter struct{}

func NewProdSecretAdapter() *ProdSecretAdapter {
	return &ProdSecretAdapter{}
}

func (adapter *ProdSecretAdapter) LoadSecrets() error {
	return nil
}
