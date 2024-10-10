package database

import (
	"database/sql"
	"log"

	"br.com.cleiton/events/internal/domain/entities"
)

type WalletDatabaseInterface interface {
	CountWalletByUUID(walletUUID string) (int, error)
	GetWalletByUUID(walletUUID string) (*entities.Wallet, error)
	UpdateWalletValueById(tx *sql.Tx, value float64, id int) (float64, error)
}

type walletDatabase struct {
	db *sql.DB
}

func NewWalletDatabase(db *sql.DB) walletDatabase {
	return walletDatabase{db: db}
}

func (w *walletDatabase) CountWalletByUUID(walletUUID string) (int, error) {
	rows, err := w.db.Query("SELECT COUNT(*) FROM wallet WHERE id_uuid = $1", walletUUID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (w *walletDatabase) GetWalletByUUID(walletUUID string) (*entities.Wallet, error) {
	row := w.db.QueryRow("SELECT w.id, w.created_at, w.document_number, w.id_uuid, w.saldo FROM wallet w WHERE w.id_uuid = $1", walletUUID)

	var wallet entities.Wallet
	err := row.Scan(&wallet.Id, &wallet.CreatedAt, &wallet.DocumentNumber, &wallet.UUID, &wallet.Saldo)
	if err != nil {
		return nil, err
	}

	return &wallet, err
}

func (w *walletDatabase) UpdateWalletValueById(tx *sql.Tx, value float64, id int) (float64, error) {
	var newValue float64
	err := tx.QueryRow("UPDATE wallet SET saldo = saldo + $1 WHERE id = $2 RETURNING saldo ", value, id).Scan(&newValue)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return newValue, nil
}
