package entities

import "time"

type Extract struct {
	Id        int
	IdUUID    string
	WalletId  int
	Type      int
	Value     float64
	Status    int
	CreateAt  time.Time
	UpdatedAt time.Time
}
