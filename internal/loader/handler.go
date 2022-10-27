package loader

import "github.com/labstack/echo/v4"

type Handler interface {
	ParseDocuments(ctx echo.Context) error
}
