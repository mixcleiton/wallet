package ports

import "br.com.cleiton/wallet/internal/domain/entities"

type CreateWallet interface {
	Create(wallet entities.Wallet) error
}
