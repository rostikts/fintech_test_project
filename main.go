package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/infrastructure/datastore"
	"net/http"
)

func main() {
	e := echo.New()
	db := datastore.NewDB(config.Config.DBConfig)
	if err := db.Ping(); err != nil {
		log.DefaultLogger.Error().Err(err).Msg("Lost connection to db")
	}
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, I'm working!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
