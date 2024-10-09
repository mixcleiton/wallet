package usecases

import (
	"errors"

	database "br.com.cleiton/wallet/internal/adapters/output/database/wallet"
	"br.com.cleiton/wallet/internal/domain/entities"
)

type listWalletUC struct {
	walletDatabase database.IWalletDatabase
}

func NewListWalletInfo(walletDatabase database.IWalletDatabase) listWalletUC {
	return listWalletUC{walletDatabase: walletDatabase}
}

func (l *listWalletUC) GetWalletInfo(walletId string, documentNumber string) (*entities.Wallet, error) {

	if walletId == "" {
		return nil, errors.New("identificação da carteira é obrigatória")
	}

	if documentNumber == "" {
		return nil, errors.New("número do documento do dono da carteira é obrigatória")
	}

	return l.walletDatabase.GetWallet(walletId, documentNumber)
}
