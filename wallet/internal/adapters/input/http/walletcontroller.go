package http

import (
	"encoding/json"
	"io"
	"net/http"

	"br.com.cleiton/wallet/internal/adapters/input/http/request"
	"br.com.cleiton/wallet/internal/domain/entities"
	"br.com.cleiton/wallet/internal/domain/ports"
	"github.com/labstack/echo/v4"
)

type walletController struct {
	createWalletUC ports.CreateWallet
}

func New(createWalletUC ports.CreateWallet) walletController {
	return walletController{createWalletUC: createWalletUC}
}

func (w *walletController) CreateWallet(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	var walletRequest request.WalletRequest
	err = json.Unmarshal(body, &walletRequest)
	if err != nil {
		return err
	}

	walletEntity := entities.Wallet{
		DocumentNumber: walletRequest.DocumentNumber,
		IdUUID:         walletRequest.IdUUID,
	}

	w.createWalletUC.Create(walletEntity)

	return c.JSON(http.StatusOK, nil)
}
