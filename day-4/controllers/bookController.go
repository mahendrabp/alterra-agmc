package controllers

import (
	"fmt"
	"net/http"

	"github.com/mahendrabp/alterra-agmc/day-4/lib/database"
	"github.com/mahendrabp/alterra-agmc/day-4/middlewares"
	"github.com/mahendrabp/alterra-agmc/day-4/models"

	"github.com/labstack/echo/v4"
)

func GetBookControllers(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func GetBookByIDControllers(c echo.Context) error {
	books, err := database.GetBookByID(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func CreateBookController(c echo.Context) error {

	book := models.Book{}

	userId := middlewares.GetUserIdFromToken(c)

	if err := c.Bind(&book); err != nil {
		fmt.Printf("error : %v", err)
	}

	if err := c.Validate(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book.CreatedBy = userId
	books, err := database.CreateBook(book)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})

}

func UpdateBookController(c echo.Context) error {

	book := models.Book{}

	if err := c.Bind(&book); err != nil {
		fmt.Printf("error : %v", err)
	}

	books, err := database.UpdateBook(book, c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func DeleteBookController(c echo.Context) error {
	_, err := database.DeleteBook(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   "{}",
	})
}
