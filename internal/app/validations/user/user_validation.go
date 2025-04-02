package validation

import (
	"errors"
	"fmt"
	app "gtodo/internal/app/entity"
	"regexp"
	"strings"
)

func ValidateUserRegister(user *app.UserRegister) error {
	fmt.Println(user)
	if len(user.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}

	if len(user.Name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	if len(user.Password) < 8 || !containsUppercase(user.Password) {
		return errors.New("password must be at least 8 characters long and contain at least one uppercase letter")
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func containsUppercase(password string) bool {
	for _, char := range password {
		if strings.ToUpper(string(char)) == string(char) && char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}
