package usecases

import (
	database "br.com.cleiton/wallet/internal/adapters/output/database/extract"
	"br.com.cleiton/wallet/internal/domain/entities"
)

type extractWalletUC struct {
	extractDatabase database.IExtractDatabase
}

func NewExtractUC(extractDatabase database.IExtractDatabase) extractWalletUC {
	return extractWalletUC{extractDatabase: extractDatabase}
}

func (e *extractWalletUC) GetExtract(walletId string, documentNumber string, page int64, size int64) ([]entities.Extract, error) {
	return e.extractDatabase.GetExtract(walletId, documentNumber, page, size)
}
