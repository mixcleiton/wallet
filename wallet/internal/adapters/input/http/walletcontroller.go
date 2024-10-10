package http

import (
	"encoding/json"
	"io"
	"net/http"

	"br.com.cleiton/wallet/internal/adapters/input/http/request"
	"br.com.cleiton/wallet/internal/adapters/input/http/response"
	"br.com.cleiton/wallet/internal/domain/entities"
	"br.com.cleiton/wallet/internal/domain/ports"
	"github.com/labstack/echo/v4"
)

type walletController struct {
	createWalletUC ports.CreateWallet
	listWalletUC   ports.ListWallet
}

func NewWalletController(createWalletUC ports.CreateWallet, listWalletUC ports.ListWallet) walletController {
	return walletController{createWalletUC: createWalletUC, listWalletUC: listWalletUC}
}

// @Summary      Post wallet info
// @Description  Post wallet information
// @Tags        wallet
// @Accept      json
// @Produce     json
// @Param wallet body request.WalletRequest true "Wallet Identification"
// @Success     201
// @Failure     400 Bad Request
// @Router      /api/v1/wallet [post]
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

	err = w.createWalletUC.Create(walletEntity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

// @Summary      Get wallet info
// @Description  Get wallet information
// @Tags        wallet
// @Accept      json
// @Produce     json
// @Param walletId path string true "Wallet ID"
// @Param documentNumber path string true "Document Number"
// @Success     200 {object} response.WalletResponse
// @Failure     400 Bad Request
// @Router      /api/v1/wallet/{walletId}/{documentNumber} [get]
func (w *walletController) GetWalletInfo(c echo.Context) error {
	walletId := c.Param("walletId")
	documentNumber := c.Param("documentNumber")

	wallet, err := w.listWalletUC.GetWalletInfo(walletId, documentNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if wallet.Id == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}

	walletResponse := response.WalletResponse{
		Saldo:          wallet.Saldo,
		CreateAt:       wallet.CreateAt.Format("2006-01-02"),
		DocumentNumber: wallet.DocumentNumber,
		IdUUID:         wallet.IdUUID,
	}

	return c.JSON(http.StatusOK, walletResponse)
}
