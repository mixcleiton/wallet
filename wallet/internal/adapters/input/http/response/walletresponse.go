package response

// @BasePath /v1
// @securityDefinitions.basic BasicAuth
// @tag(name="WalletResponse", description="Information about wallet")
type WalletResponse struct {
	Saldo          float64
	CreateAt       string
	DocumentNumber string
	IdUUID         string
}
