package usecase

import (
	"context"
	"errors"
	todo "gtodo/internal/app/entity"
	"gtodo/internal/app/helpers"
	"gtodo/internal/app/repository"
	validation "gtodo/internal/app/validations/todo"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoUseCase interface {
	RegisterTodo(ctx context.Context, todo *todo.Todo, c echo.Context) (int, error)
	GetTodoByID(ctx context.Context, id string) (*todo.Todo, error)
	GetAllTodos(ctx context.Context) ([]todo.Todo, error)
	UpdateTodo(ctx context.Context, todo *todo.Todo, c echo.Context) error
	DeleteTodo(ctx context.Context, id string) error
}

type TodoInteraction struct {
	TodoRepository repository.TodoRepository
	Context        echo.Context
	UserUseCase    UserUseCase
}

func (t *TodoInteraction) RegisterTodo(ctx context.Context, todo *todo.Todo, c echo.Context) (int, error) {
	/*authHeader := c.Request().Header.Get("Authorization")

	token, err := utils.GetTokenFromHeader(authHeader)
	if err != nil {
		return 0, err
	}

	username, err := utils.GetUsername(token)
	if err != nil {
		fmt.Println("Error:", err)
	}


	userId, err := t.UserUseCase.GetUserId(username)
	if err != nil {
		return 0, err
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return 0, errors.New("Something went wrong")
	}*/
	userId, err := getUserIdFromContext(c, t.UserUseCase)
	if err != nil {
		return 0, err
	}

	todo.UserID = userId

	if err := validation.ValidateTodo(todo); err != nil {
		return 2, err
	}

	err = t.TodoRepository.CreateTodo(todo)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (t *TodoInteraction) GetTodoByID(ctx context.Context, id string) (*todo.Todo, error) {
	todo, err := t.TodoRepository.FindTodoById(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (t *TodoInteraction) GetAllTodos(ctx context.Context) ([]todo.Todo, error) {
	todos, err := t.TodoRepository.GetAllTodos()
	if err != nil {
		return nil, errors.New("could not retrieve todos")
	}

	return todos, nil
}

func (t *TodoInteraction) UpdateTodo(ctx context.Context, todo *todo.Todo ,c echo.Context) error {
	// TODO: hala düzeltilebilecek şeyler mevcut ama katmanlar güzel ayrıldı mantığı belli
	userId, err := getUserIdFromContext(c, t.UserUseCase)
	if err != nil {
		return err
	}

	todo.UserID = userId

	if err := validation.ValidateTodo(todo); err != nil {
		return err
	}

	tempTodo, err := t.GetTodoByID(ctx, todo.ID.String())
	if err != nil {
		return errors.New("Something Went Wrong")
	}

	if tempTodo.UserID != userId {
		return errors.New("You have to be the owner of the todo to update the todo :D")
	}

	// needs to stay same
	todo.CreatedAt = tempTodo.CreatedAt

	err = t.TodoRepository.UpdateTodo(todo)
	if err != nil {
		return errors.New("failed to update todo")
	}

	return nil
}

func (t *TodoInteraction) DeleteTodo(ctx context.Context, id string) error {
	err := t.TodoRepository.DeleteTodo(id)
	if err != nil {
		return errors.New("failed to delete todo")
	}

	return nil
}

func UseCaseTodo(repo repository.TodoRepository, userUseCase UserUseCase) TodoUseCase {
	return &TodoInteraction{
		TodoRepository: repo,
		UserUseCase:    userUseCase,
	}
}

func getUserIdFromContext(c echo.Context, userUseCase UserUseCase) (uuid.UUID, error) {
	username, err := helpers.GetUsername(c)
	if err != nil {
		return uuid.UUID{}, errors.New("failed to extract username from token")
	}

	userId, err := userUseCase.GetUserId(username)
	if err != nil {
		return uuid.UUID{}, errors.New("error retrieving userId")
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid userId format")
	}

	return parsedUserId, nil
}
