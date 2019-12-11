package main

import (
	"github.com/allochi/sample-echo/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	router := echo.New()
	router.GET("/", handlers.Home)

	// add `data` section to API
	// -------------------------
	data := router.Group("/data")
	data.GET("/quotes", handlers.GetQuotes)
	data.GET("/tenders", handlers.GetTenders)

	// add `secrets` section to API with authentication
	// ------------------------------------------------
	secrets := router.Group("/secrets")
	secrets.Use(middleware.BasicAuth(handlers.Authenticate))
	secrets.GET("/names", handlers.GetNames)

	router.Logger.Fatal(router.Start(":3300"))
}
