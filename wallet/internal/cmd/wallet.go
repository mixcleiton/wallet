package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"br.com.cleiton/wallet/internal/adapters/input/http"
	extractDatabase "br.com.cleiton/wallet/internal/adapters/output/database/extract"
	walletDatabase "br.com.cleiton/wallet/internal/adapters/output/database/wallet"
	"br.com.cleiton/wallet/internal/config"
	"br.com.cleiton/wallet/internal/domain/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type startWalletCmd struct {
	databaseConfig *config.DatabaseConfig
}

func NewStartWallet() startWalletCmd {
	return startWalletCmd{}
}

func (s *startWalletCmd) StartWallet() {
	e := echo.New()
	var err error
	s.databaseConfig, err = config.LoadConfig("database.yaml")
	if err != nil {
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	db, err := sql.Open("postgres", s.configDB())
	if err != nil {
		panic(err)
	}

	dbWallet := walletDatabase.NewWallet(db)
	dbExtract := extractDatabase.NewExtract(db)
	createWalletUC := usecases.NewCreateUC(dbWallet)
	listWalletUC := usecases.NewListWalletInfo(dbWallet)
	extractUC := usecases.NewExtractUC(dbExtract)

	wallet := http.NewWalletController(createWalletUC, &listWalletUC)
	extractController := http.NewExtractController(&extractUC)

	g := e.Group("/api/v1")
	g.POST("/wallet", wallet.CreateWallet)
	g.GET("/wallet/:walletId/:documentNumber", wallet.GetWalletInfo)
	g.GET("/extract", extractController.GenerateExtract)
	e.Logger.Fatal(e.Start(":8089"))
}

func (s *startWalletCmd) configDB() string {
	urlDb := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"host.docker.internal", 5432, "root", "wallettest", "postgres")

	log.Println(urlDb)
	return urlDb
}
