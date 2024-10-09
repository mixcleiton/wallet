package cmd

import (
	"br.com.cleiton/wallet/internal/adapters/input/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartWallet() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	wallet := http.New()
	e.POST("/wallet", wallet.CreateWallet)

	e.Logger.Fatal(e.Start(":8080"))
}
