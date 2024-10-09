package ports

import "br.com.cleiton/wallet/internal/domain/entities"

type ListWallet interface {
	GetWalletInfo(walletId string, documentNumber string) (*entities.Wallet, error)
}
