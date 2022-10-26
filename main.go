package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, I'm working!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
