package utils

import (
	"strings"

	"github.com/labstack/echo/v4"
)

var c echo.Context

// this func does not any checkings beacause already has middleware
func GetUsername() string {
	authHeader := c.Request().Header.Get("Authorization")

	tokenParts := strings.Split(authHeader, " ")

	username, _ := VerifyToken(tokenParts[1])
	
	return username
}
