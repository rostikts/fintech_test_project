package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rostikts/fintech_test_project/infrastructure/registry"
)

func NewRouter(e *echo.Echo, controller registry.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/parse", controller.Transaction.ParseDocuments)
	transactionsGroup.GET("", controller.Transaction.GetTransactions)
	return e
}
