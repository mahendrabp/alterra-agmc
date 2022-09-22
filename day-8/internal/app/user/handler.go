package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mahendrabp/alterra-agmc/day-8/internal/entity"
)

type userHandler struct {
	userUC entity.UserService
}

func NewUserHandler(e *echo.Echo, userUC entity.UserService) {
	handler := &userHandler{
		userUC: userUC,
	}

	api := e.Group("/v1")
	// api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)
}

func (h *userHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	user := entity.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Printf("error : %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	if err := c.Validate(user); err != nil {
		fmt.Printf("error : %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userLogin, err := h.userUC.LoginUsers(ctx, &user)

	if err != nil {
		fmt.Printf("error : %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"data":   userLogin,
	})

}
