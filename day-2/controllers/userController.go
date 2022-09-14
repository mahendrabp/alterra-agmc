package controllers

import (
	"fmt"
	"net/http"

	"github.com/mahendrabp/alterra-agmc/day-2/lib/database"
	"github.com/mahendrabp/alterra-agmc/day-2/models"

	"github.com/labstack/echo/v4"
)

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
