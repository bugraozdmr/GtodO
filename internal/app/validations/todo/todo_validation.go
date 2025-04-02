package validation

import (
	"errors"
	"fmt"
	app "gtodo/internal/app/entity"
	"strings"
	"time"

	"github.com/google/uuid"
)

func ValidateTodo(todo *app.Todo) error {
	if strings.TrimSpace(todo.Title) == "" {
		return errors.New("title cannot be empty")
	}

	// maybe we dont need this one at all !
	if len(todo.Description) > 1000 {
		return errors.New("description cannot be longer than 500 characters")
	}

	if todo.UserID == uuid.Nil {
		fmt.Println(todo.UserID)
		return errors.New("invalid user ID")
	}

	if todo.DueDate != nil {
		if todo.DueDate.Before(time.Now()) {
			return errors.New("due date cannot be in the past")
		}
	}

	return nil
}
