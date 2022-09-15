package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func GetUserIdFromToken(e echo.Context) int {
	authToken := e.Request().Header.Get("Authorization")

	if authToken == "" {
		fmt.Println("error on this")
		return 0
	}

	tokenString := strings.Split(authToken, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error token")
		}
		return []byte("secret"), nil
	})

	if !token.Valid || err != nil {
		return 0
	}

	userId := token.Claims.(jwt.MapClaims)["userId"].(float64)

	if userId != 0 {
		return int(userId)
	}
	return 0
}
