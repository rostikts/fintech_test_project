package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/rostikts/fintech_test_project/docs"
	"github.com/rostikts/fintech_test_project/infrastructure/registry"
	"github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, controller registry.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("/api")
	v1 := api.Group("/v1")
	transactionsGroup := v1.Group("/transactions")
	transactionsGroup.POST("/parse", controller.Transaction.ParseDocuments)
	transactionsGroup.GET("", controller.Transaction.GetTransactions)
	return e
}
