package ports

import "br.com.cleiton/wallet/internal/domain/entities"

type ExtractWallet interface {
	GetExtract(walletId string, documentNumber string, page int, size int) ([]entities.Extract, error)
}
