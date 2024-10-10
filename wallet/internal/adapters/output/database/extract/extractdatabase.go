package database

import (
	"database/sql"

	"br.com.cleiton/wallet/internal/domain/entities"
)

type extractDatabase struct {
	db *sql.DB
}

func NewExtract(db *sql.DB) IExtractDatabase {
	return &extractDatabase{db: db}
}

func (e *extractDatabase) GetExtract(walletId string, documentNumber string, page int, size int) ([]entities.Extract, error) {
	offset := (page - 1) * size

	rows, err := e.db.Query("SELECT evt.id, evt.type_id, evt.value, evt.created_at "+
		"FROM extract e INNER JOIN wallet w ON w.id = e.wallet_id "+
		"INNER JOIN event evt ON evt.wallet_id = w.id "+
		"WHERE w.id = $1 AND w.document_number = $2 "+
		"ORDER BY e.created_at desc LIMIT $3 OFFSET $4", walletId, documentNumber, size, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var extracts []entities.Extract
	for rows.Next() {
		var extract entities.Extract
		err := rows.Scan(&extract.Id, &extract.Type, &extract.Value, &extract.CreateAt)
		if err != nil {
			return nil, err
		}

		extracts = append(extracts, extract)
	}

	return extracts, nil
}
