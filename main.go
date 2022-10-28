package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/config"
	_ "github.com/rostikts/fintech_test_project/docs"
	"github.com/rostikts/fintech_test_project/infrastructure/datastore"
	"github.com/rostikts/fintech_test_project/infrastructure/registry"
	"github.com/rostikts/fintech_test_project/infrastructure/router"
	"net/http"
)

// @title           Swagger API
// @version         1.0
// @description     Documentation for test project for EVO trainee program
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rostik
// @contact.email  rostiktsyapura@gmail.com

// @host      localhost:8000
// @BasePath  /api/v1
func main() {
	e := echo.New()
	db := datastore.NewDB(config.Config.DBConfig)
	if err := db.Ping(); err != nil {
		log.DefaultLogger.Error().Err(err).Msg("Lost connection to db")
	}

	appController := registry.AppController{
		Transaction: registry.NewTransactionRegistry(db),
	}
	e = router.NewRouter(e, appController)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, I'm working!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
