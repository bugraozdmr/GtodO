package helpers

// seperated package for solving cycle errors

import (
	"fmt"
	"gtodo/internal/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetUsernameFromToken(tokenString string) (string, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	username, err := utils.VerifyToken(tokenString)
	if err != nil {
		return "", err
	}

	return username, nil
}

func GetUsername(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	token, err := GetTokenFromHeader(authHeader)
	if err != nil {
		return "", err
	}

	username, err := GetUsernameFromToken(token)
	if err != nil {
		return "", fmt.Errorf("error extracting username from token: %v", err)
	}

	return username, nil
}

