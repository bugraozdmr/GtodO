package helpers

import (
	"errors"
	"strings"
)

func GetTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("Authorization token not provided")
	}
	
	splitToken := strings.Split(authHeader, " ")

	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", errors.New("Invalid token format")
	}

	return splitToken[1], nil
}
