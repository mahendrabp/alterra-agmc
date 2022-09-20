package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mahendrabp/alterra-agmc/day-6/database"
	"github.com/mahendrabp/alterra-agmc/day-6/internal/app/user"
	"github.com/mahendrabp/alterra-agmc/day-6/internal/repository"
	"github.com/mahendrabp/alterra-agmc/day-6/internal/utils"
)

func main() {

	dbInstance := database.GetConnection()

	// Setup repository
	userRepo := repository.NewUser(dbInstance)

	// Setup usecase
	userSv := user.NewUserService(userRepo)

	e := echo.New()
	e.Validator = utils.GetValidator()

	user.NewUserHandler(e, userSv)

	e.Logger.Fatal(e.Start(":8000"))
}
