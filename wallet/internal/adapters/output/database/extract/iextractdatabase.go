package database

import "br.com.cleiton/wallet/internal/domain/entities"

type IExtractDatabase interface {
	GetExtract(walletId string, documentNumber string, page int64, size int64) ([]entities.Extract, error)
}
