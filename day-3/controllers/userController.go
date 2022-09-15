package controllers

import (
	"fmt"
	"net/http"

	"github.com/mahendrabp/alterra-agmc/day-3/config"
	"github.com/mahendrabp/alterra-agmc/day-3/lib/database"
	"github.com/mahendrabp/alterra-agmc/day-3/middlewares"
	"github.com/mahendrabp/alterra-agmc/day-3/models"

	"github.com/labstack/echo/v4"
)

func loginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Printf("error : %v", err)
	}

	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, e := loginUsers(&user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"data":   users,
	})
}

func GetUserControllers(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func GetUserByIDControllers(c echo.Context) error {
	userID := c.Param("id")
	users, err := database.GetUserByID(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func CreateUserControllers(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Printf("error : %v", err)
	}

	users, err := database.CreateUsers(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func UpdateUserControllers(c echo.Context) error {
	userID := c.Param("id")
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Printf("error : %v", err)
	}
	users, err := database.UpdateUsers(user, userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func DeleteUserControllers(c echo.Context) error {
	userID := c.Param("id")

	_, err := database.DeleteUserByID(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user",
		"data":   "{}",
	})
}
