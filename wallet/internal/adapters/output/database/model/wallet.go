package model

import "time"

type Wallet struct {
	Id        string    `json:"id"`
	Saldo     float64   `json:"saldo"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_ad"`
	User      string    `json:"user"`
}
