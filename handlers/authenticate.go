package handlers

import (
	"github.com/labstack/echo"
)

func Authenticate(username, password string, ctx echo.Context) (bool, error) {
	if username == "allochi" && password == "nopass" {
		return true, nil
	}
	return false, nil
}
