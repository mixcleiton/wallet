package http

import (
	"net/http"
	"strconv"

	"br.com.cleiton/wallet/internal/domain/ports"
	"github.com/labstack/echo/v4"
)

type extractController struct {
	extractWalletUC ports.ExtractWallet
}

func NewExtractController(extractWalletUC ports.ExtractWallet) *extractController {
	return &extractController{extractWalletUC: extractWalletUC}
}

// @Summary      Get wallet extract
// @Description  Get wallet extract information
// @Tags        extract
// @Accept      json
// @Produce     json
// @Param       Pet body pet.Pet true "Pet object that needs to be added to the store"
// @Success     200 {object} []
// @Failure     400 Bad Request
// @Router       /wallet/:walletId/:documentNumber [get]
func (e *extractController) GenerateExtract(c echo.Context) error {

	walletId := c.QueryParam("walletId")
	documentNumber := c.QueryParam("documentNumber")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = 10
	}

	extracts, err := e.extractWalletUC.GetExtract(walletId, documentNumber, page, size)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, extracts)
}
