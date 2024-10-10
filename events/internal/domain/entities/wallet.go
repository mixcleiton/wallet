package entities

import "time"

type Wallet struct {
	Id             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           string
	DocumentNumber string
	UUID           string
	Saldo          float64
}
