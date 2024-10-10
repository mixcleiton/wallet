package response

// @BasePath /v1
// @securityDefinitions.basic BasicAuth
// @tag(name="ExtractResponse", description="Operations about extract")
type ExtractResponse struct {
	IdUUID    string  `json:"idUuid"`
	WalletId  int     `json:"walletId"`
	Type      int     `json:"type"`
	Value     float64 `json:"value"`
	Status    int     `json:"status"`
	CreatedAt string  `json:"createdAt"`
}
