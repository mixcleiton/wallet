package http

import (
	"log"
	"net/http"
	"strconv"

	"br.com.cleiton/wallet/internal/adapters/input/http/response"
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
// @Param walletId query string true "Wallet ID"
// @Param documentNumber query string true "Document Number"
// @Param page query int true "Page"
// @Param size query int true "Size"
// @Success     200 {object} []response.ExtractResponse
// @Failure     400 Bad Request
// @Router      /api/v1/extract [get]
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
		log.Printf("erro ao buscar extrato, erro %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	extractsResponse := make([]*response.ExtractResponse, 0)
	for _, extract := range extracts {
		log.Println(extract.IdUUID)
		extractResponse := &response.ExtractResponse{
			IdUUID:    extract.IdUUID,
			Status:    extract.Status,
			Value:     extract.Value,
			Type:      extract.Type,
			CreatedAt: extract.CreateAt.Format("2006-01-02"),
		}

		extractsResponse = append(extractsResponse, extractResponse)
	}

	return c.JSON(http.StatusOK, extractsResponse)
}
