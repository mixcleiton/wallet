package database

import (
	"database/sql"

	"br.com.cleiton/wallet/internal/domain/entities"
)

type WalletDatabase struct {
	db *sql.DB
}

func NewWallet(db *sql.DB) IWalletDatabase {
	return &WalletDatabase{db: db}
}

func (wd *WalletDatabase) Save(wallet entities.Wallet) error {
	_, err := wd.db.Exec("INSERT INTO wallet(saldo, document_number, created_at, \"user\", id_uuid) "+
		"VALUES ($1, $2, $3, $4, $5)", wallet.Saldo, wallet.DocumentNumber, wallet.CreateAt, wallet.User, wallet.IdUUID)
	if err != nil {
		return err
	}

	return nil
}

func (wd *WalletDatabase) GetWallet(walletId string, documentNumber string) (*entities.Wallet, error) {
	rows, err := wd.db.Query("SELECT w.id, w.id_uuid, w.document_number, w.saldo, w.created_at "+
		"FROM wallet w WHERE w.id_uuid = $1 AND w.document_number = $2 ", walletId, documentNumber)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var wallet entities.Wallet
	for rows.Next() {
		err := rows.Scan(&wallet.Id, &wallet.IdUUID, &wallet.DocumentNumber, &wallet.Saldo, &wallet.CreateAt)
		if err != nil {
			return nil, err
		}
	}

	return &wallet, nil
}
