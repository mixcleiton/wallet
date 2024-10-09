package database

import (
	"database/sql"

	"br.com.cleiton/wallet/internal/domain/entities"
)

type WalletDatabase struct {
	db *sql.DB
}

func NewWallet() IWalletDatabase {
	return &WalletDatabase{}
}

func (wd *WalletDatabase) Save(wallet entities.Wallet) error {
	_, err := wd.db.Exec("INSERT INTO wallet(id, saldo, document_number, ) VALUES ()")
	if err != nil {
		return err
	}

	return nil
}

func (wd *WalletDatabase) GetWallet(walletId string, documentNumber string) (*entities.Wallet, error) {
	rows, err := wd.db.Query("SELECT w.id_uuid, w.document_number, w.saldo, w.created_at "+
		"FROM wallets w WHERE w.id = $1 AND w.document_number = $2 ", walletId, documentNumber)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var wallet entities.Wallet
	for rows.Next() {
		err := rows.Scan(&wallet.IdUUID, &wallet.DocumentNumber, &wallet.Saldo, &wallet.CreateAt)
		if err != nil {
			return nil, err
		}
	}

	return &wallet, nil
}
