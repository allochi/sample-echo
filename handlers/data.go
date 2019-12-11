package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func GetQuotes(ctx echo.Context) error {
	response := struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}{"allochi", 45}

	return ctx.JSON(http.StatusOK, response)
}

func GetTenders(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "These are tenders!")
}

func GetNames(ctx echo.Context) error {
	users, err := ioutil.ReadFile("data/users.json")
	if err != nil {
		return err
	}
	return ctx.JSONBlob(http.StatusOK, users)
}
