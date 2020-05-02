package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func getUserFromJWT(c echo.Context) (int64, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int64(claims["id"].(float64))

	return userID, nil
}
