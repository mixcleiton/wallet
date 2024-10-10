package database

import (
	"database/sql"
	"log"
)

type ExtractDatabaseInterface interface {
	CreateExtract(tx *sql.Tx, eventId int) error
}

type extractDatabase struct {
	db *sql.DB
}

func NewExtractDatabase(db *sql.DB) walletDatabase {
	return walletDatabase{db: db}
}

func (w *walletDatabase) CreateExtract(tx *sql.Tx, eventId int) error {
	_, err := tx.Exec("INSERT INTO extract(event_id, created_at) values ($1, NOW())", eventId)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
