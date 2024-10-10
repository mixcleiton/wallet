package requestdto

type EventRequest struct {
	UUID        string  `json:"IdUUID"`
	WalletUUID  string  `json:"wallet_uuid"`
	Type        int     `json:"type_id"`
	Value       float64 `json:"value"`
	EventId     string  `json:"event_uuid"`
	Description string  `json:"description"`
	Id          int     `json:"id"`
}
