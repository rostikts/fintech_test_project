package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/internal/transaction"
	"net/http"
	"os"
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

// ParseDocuments godoc
// @Summary      parse document
// @Description  parse document with transactions
// @Tags         transactions
// @Accept       json
// @Produce      json
// @Param        payload          body     				 parseDocumentsRequest       false  "payload"
// @Success      201  {object}    parseDocumentResponse
// @Failure		 400  {object}    echo.HTTPError
// @Router       /transactions/parse [post]
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

// GetTransactions godoc
// @Summary      returns array of transactions
// @Description  returns array of filtered transactions
// @Tags         transactions
// @Produce      json
// @Param        terminal_id          query     int       false  "filter by terminal_id"
// @Param        transaction_id       query     int       false  "filter by transaction_id"
// @Param        status               query     string    false  "filter by status"
// @Param        payment_type         query     string    false  "filter by payment_type"
// @Param        from              	  query     string    false  "filter from start date"       Format(date)
// @Param        to                   query     string    false  "filter to ending date"        Format(date)
// @Param        payment_narrative    query     string    false  "partial match by narrative"
// @Success      200  {array}   models.Transaction
// @Failure		 400  {object}  echo.HTTPError
// @Router       /transactions [get]
func (h transactionHandler) GetTransactions(ctx echo.Context) error {
	filters := extractFilters(ctx)

	result, err := h.service.GetTransactions(filters)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, result)
}

// GetTransactionsCSV godoc
// @Summary      returns cvs with transactions
// @Description  returns csv file with filtered transactions
// @Tags         transactions
// @Produce      json
// @Param        terminal_id          query     int       false  "filter by terminal_id"
// @Param        transaction_id       query     int       false  "filter by transaction_id"
// @Param        status               query     string    false  "filter by status"
// @Param        payment_type         query     string    false  "filter by payment_type"
// @Param        from              	  query     string    false  "filter from start date"       Format(date)
// @Param        to                   query     string    false  "filter to ending date"        Format(date)
// @Param        payment_narrative    query     string    false  "partial match by narrative"
// @Success      200  {array}   models.Transaction
// @Failure		 400  {object}  echo.HTTPError
// @Router       /transactions/csv [get]
func (h transactionHandler) GetTransactionsCSV(ctx echo.Context) error {
	filters := extractFilters(ctx)

	fileName, err := h.service.GetTransactionsCSV(filters)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer func() {
		if err := os.Remove(fileName); err != nil {
			log.DefaultLogger.Error().Err(err).Msg("the tmp sent file is not deleted")
		}
	}()
	return ctx.File(fileName)
}

func extractFilters(ctx echo.Context) map[string]string {
	result := make(map[string]string, len(allowedFilters))
	for _, filter := range allowedFilters {
		value := ctx.QueryParams().Get(filter)
		if value != "" {
			result[filter] = value
		}
	}
	return result
}
