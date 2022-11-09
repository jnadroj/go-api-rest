package main

import (
	"apirest/configs"
	"apirest/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func indexResponse(c echo.Context) error {
	return c.JSON(http.StatusOK, &echo.Map{"message": "Hello, World!"})
}

func main() {
	e := echo.New()

	//run database
	configs.ConnectDB()

	e.GET("/", indexResponse)
	routes.UserRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}