package model

import "time"

type Extract struct {
	Id        int64     `JSON:"id"`
	IdUUID    string    `JSON:"id_uuid"`
	WalletId  int64     `JSON:"wallet_id"`
	TypeId    int64     `JSON:"type_id"`
	Value     float64   `JSON:"value"`
	StatusId  int64     `JSON:"status_id"`
	CreatedAt time.Time `JSON:"created_at"`
	UpdatedAt time.Time `JSON:"updated_at"`
}
