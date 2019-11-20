package main

import "github.com/labstack/echo/v4"

import "net/http"

func main() {
	router := echo.New()
	router.GET("/", home)
	data := router.Group("/data")
	data.GET("/quotes", getQuotes)

	router.Logger.Fatal(router.Start(":3300"))
}

func home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world!")
}

func getQuotes(ctx echo.Context) error {
	response := struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}{"allochi", 45}

	return ctx.JSON(http.StatusOK, response)
}
