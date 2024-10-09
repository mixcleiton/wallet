package database

import "br.com.cleiton/wallet/internal/domain/entities"

type IExtractDatabase interface {
	GetExtract(walletId string, documentNumber string, page int, size int) ([]entities.Extract, error)
}
