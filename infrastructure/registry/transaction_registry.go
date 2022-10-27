package registry

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostikts/fintech_test_project/internal/loader"
	"github.com/rostikts/fintech_test_project/internal/loader/handler"
	"github.com/rostikts/fintech_test_project/internal/loader/repository"
	"github.com/rostikts/fintech_test_project/internal/loader/service"
)

func NewTransactionRegistry(db *sqlx.DB) loader.Handler {
	return handler.NewTransactionHandler(service.NewLoaderService(repository.NewTransactionRepository(db)))
}
