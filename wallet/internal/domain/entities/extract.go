package entities

import "time"

type Extract struct {
	Id        int64
	IdUUID    string
	WalletId  int64
	Type      int64
	Value     float64
	Status    int64
	CreateAt  time.Time
	UpdatedAt time.Time
}
