package routes

import (
	"apirest/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/users", controllers.GetAllUsers) 
	e.GET("/user/:userId", controllers.GetAUser) 
	e.POST("/user", controllers.CreateUser)
	e.PATCH("/user/:userId", controllers.EditAUser)
	e.DELETE("/user/:userId", controllers.DeleteAUser)
}