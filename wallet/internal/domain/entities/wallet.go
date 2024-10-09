package entities

import "time"

type Wallet struct {
	Id             int64
	Saldo          float64
	CreateAt       time.Time
	UpdatedAt      time.Time
	User           string
	DocumentNumber string
	IdUUID         string
}
