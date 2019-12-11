package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world!")
}
