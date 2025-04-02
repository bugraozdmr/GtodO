package delivery

import (
	app "gtodo/internal/app/entity"
	"gtodo/internal/app/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

type TodoUseCase interface {
	CreateTodoHandler(c echo.Context) error
	GetAllTodosHandler(c echo.Context) error
	GetTodoByIDHandler(c echo.Context) error
	UpdateTodoHandler(c echo.Context) error
	DeleteTodoHandler(c echo.Context) error
}

func (t *TodoHandler) CreateTodoHandler(c echo.Context) error {
	var todo app.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}

	ctx := c.Request().Context()
	flag, err := t.todoUseCase.RegisterTodo(ctx, &todo)

	if err != nil {
		if flag == 0 {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Something went wrong",
			})
		} else if flag == 2 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Todo created successfully",
	})
}

func (t *TodoHandler) GetAllTodosHandler(c echo.Context) error {
	ctx := c.Request().Context()
	todos, err := t.todoUseCase.GetAllTodos(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) GetTodoByIDHandler(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	todo, err := t.todoUseCase.GetTodoByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, todo)
}

func (t *TodoHandler) UpdateTodoHandler(c echo.Context) error {
	id := c.Param("id")

	todoID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID format",
		})
	}

	var todo app.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}

	todo.ID = todoID

	ctx := c.Request().Context()
	if err := t.todoUseCase.UpdateTodo(ctx, &todo); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Todo updated successfully",
	})
}

func (t *TodoHandler) DeleteTodoHandler(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	if err := t.todoUseCase.DeleteTodo(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Todo deleted successfully",
	})
}

func TodoDelivery(todoUseCase usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}
