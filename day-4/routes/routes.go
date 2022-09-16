package routes

import (
	"github.com/mahendrabp/alterra-agmc/day-4/controllers"
	"github.com/mahendrabp/alterra-agmc/day-4/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte("secret")))
	e.Validator = utils.GetValidator()

	// route users
	e.POST("/v1/login", controllers.LoginUsersController)

	r.GET("/v1/users", controllers.GetUserControllers)
	e.POST("/v1/users", controllers.CreateUserControllers)
	r.GET("/v1/users/:id", controllers.GetUserByIDControllers)
	r.PUT("/v1/users/:id", controllers.UpdateUserControllers)
	r.DELETE("/v1/users/:id", controllers.DeleteUserControllers)

	e.GET("/v1/books", controllers.GetBookControllers)
	r.POST("/v1/books", controllers.CreateBookController)
	e.GET("/v1/books/:id", controllers.GetBookByIDControllers)
	r.PUT("/v1/books/:id", controllers.UpdateBookController)
	r.DELETE("/v1/books/:id", controllers.DeleteBookController)

	return e
}
