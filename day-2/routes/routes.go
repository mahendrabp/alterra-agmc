package routes

import (
	"github.com/mahendrabp/alterra-agmc/day-2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/v1/users", controllers.GetUserControllers)
	e.POST("/v1/users", controllers.CreateUserControllers)
	e.GET("/v1/users/:id", controllers.GetUserByIDControllers)
	e.PUT("/v1/users/:id", controllers.UpdateUserControllers)
	e.DELETE("/v1/users/:id", controllers.DeleteUserControllers)

	e.GET("/v1/books", controllers.GetBookControllers)
	e.POST("/v1/books", controllers.CreateBookController)
	e.GET("/v1/books/:id", controllers.GetBookByIDControllers)
	e.PUT("/v1/books/:id", controllers.UpdateBookController)
	e.DELETE("/v1/books/:id", controllers.DeleteBookController)

	return e
}
