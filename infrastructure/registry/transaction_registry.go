package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostikts/fintech_test_project/internal/transaction"
	"github.com/rostikts/fintech_test_project/internal/transaction/handler"
	"github.com/rostikts/fintech_test_project/internal/transaction/repository"
	"github.com/rostikts/fintech_test_project/internal/transaction/service"
)

func NewTransactionRegistry(db *sqlx.DB) transaction.Handler {
	return handler.NewTransactionHandler(service.NewTransactionService(repository.NewTransactionRepository(db)))
}
