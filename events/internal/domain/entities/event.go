package entities

import "time"

type Event struct {
	Id          int
	IdUUID      string
	WalletId    int
	WalletUUID  string
	Type        int
	Value       float64
	Status      int
	CreateAt    time.Time
	UpdatedAt   time.Time
	Description string
	EventUUID   string
	EventID     int
}
