package entities

type Status int

const (
	WAITING_PROCESS Status = 1
	PROCESSED       Status = 2
	COMPLETED       Status = 3
	CANCELED        Status = 4
)
