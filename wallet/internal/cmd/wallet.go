package cmd

import (
	"database/sql"
	"fmt"

	"br.com.cleiton/wallet/internal/adapters/input/http"
	database "br.com.cleiton/wallet/internal/adapters/output/database/wallet"
	"br.com.cleiton/wallet/internal/config"
	"br.com.cleiton/wallet/internal/domain/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type startWalletCmd struct {
	databaseConfig config.DatabaseConfig
}

func NewStartWallet() startWalletCmd {
	return startWalletCmd{}
}

func (s *startWalletCmd) StartWallet() {
	e := echo.New()
	var err error
	s.databaseConfig, err = config.LoadConfig("./config/database.yaml")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	db, err := sql.Open("postgres", s.configDB())
	if err != nil {
		panic(err)
	}

	dbWallet := database.NewWallet(db)
	createWalletUC := usecases.NewCreateUC(dbWallet)
	listWalletUC := usecases.NewListWalletInfo(dbWallet)

	wallet := http.NewWalletController(createWalletUC, &listWalletUC)
	e.POST("/wallet", wallet.CreateWallet)
	e.GET("/wallet/:walletId/:documentNumber", wallet.GetWalletInfo)

	e.Logger.Fatal(e.Start(":8080"))
}

func (s *startWalletCmd) configDB() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.databaseConfig.Host, s.databaseConfig.Port, s.databaseConfig.User, s.databaseConfig.Password, s.databaseConfig.DBName)
}
