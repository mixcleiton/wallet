package database

import "br.com.cleiton/wallet/internal/domain/entities"

type IWalletDatabase interface {
	Save(wallet entities.Wallet) error
	GetWallet(walletId string, documentNumber string) (*entities.Wallet, error)
}
