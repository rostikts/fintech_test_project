package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rostikts/fintech_test_project/internal/transaction"
	"net/http"
)

type transactionHandler struct {
	service transaction.Service
}

type parseDocumentsRequest struct {
	Url string `json:"url"`
}
type parseDocumentResponse struct {
	Success int64 `json:"success"`
	Failed  int64 `json:"failed"`
}

var allowedFilters = []string{"transaction_id", "terminal_id", "status", "payment_type", "from", "to", "payment_narrative"}

func NewTransactionHandler(service transaction.Service) transaction.Handler {
	return transactionHandler{service: service}
}

func (h transactionHandler) ParseDocuments(ctx echo.Context) error {
	var body parseDocumentsRequest

	if err := ctx.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "incorrect url is provided")
	}
	if len(body.Url) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "url is not provided")
	}

	success, failed, err := h.service.ParseDocument(body.Url)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, parseDocumentResponse{
		Success: success,
		Failed:  failed,
	})
}

func (h transactionHandler) GetTransactions(ctx echo.Context) error {
	filters := extractFilters(ctx)

	result, err := h.service.GetTransactions(filters)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}

func extractFilters(ctx echo.Context) map[string]string {
	result := make(map[string]string)
	for _, filter := range allowedFilters {
		value := ctx.QueryParams().Get(filter)
		if value != "" {
			result[filter] = value
		}
	}
	return result
}
