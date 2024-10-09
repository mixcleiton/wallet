package usecases

import (
	"errors"
	"time"

	database "br.com.cleiton/wallet/internal/adapters/output/database/wallet"
	"br.com.cleiton/wallet/internal/domain/entities"
	"br.com.cleiton/wallet/internal/domain/ports"
)

type createWalletUC struct {
	db database.IWalletDatabase
}

func NewCreateUC(db database.IWalletDatabase) ports.CreateWallet {
	return &createWalletUC{db: db}
}

func (c *createWalletUC) Create(wallet entities.Wallet) error {
	if wallet.IdUUID == "" {
		return errors.New("identificação da carteira é obrigatória")
	}

	if wallet.DocumentNumber == "" {
		return errors.New("número do documento é obrigatório")
	}

	wallet.CreateAt = time.Now()
	wallet.UpdatedAt = time.Now()
	wallet.Saldo = 0
	wallet.User = "sistema"

	return c.db.Save(wallet)
}
