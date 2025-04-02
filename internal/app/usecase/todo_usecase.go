package usecase

import (
	"context"
	"errors"
	todo "gtodo/internal/app/entity"
	validation "gtodo/internal/app/validations/todo"
	"gtodo/internal/app/repository"
)

type TodoUseCase interface {
	RegisterTodo(ctx context.Context, todo *todo.Todo) (int, error)
	GetTodoByID(ctx context.Context, id string) (*todo.Todo, error)
	GetAllTodos(ctx context.Context) ([]todo.Todo, error)
	UpdateTodo(ctx context.Context, todo *todo.Todo) error
	DeleteTodo(ctx context.Context, id string) error
}

type TodoInteraction struct {
	TodoRepository repository.TodoRepository
}

func (t *TodoInteraction) RegisterTodo(ctx context.Context, todo *todo.Todo) (int, error) {
	if err := validation.ValidateTodo(todo); err != nil {
		return 2, err
	}

	err := t.TodoRepository.CreateTodo(todo)
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

func (t *TodoInteraction) UpdateTodo(ctx context.Context, todo *todo.Todo) error {
	if err := validation.ValidateTodo(todo); err != nil {
		return err
	}

	err := t.TodoRepository.UpdateTodo(todo)
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

func UseCaseTodo(repo repository.TodoRepository) TodoUseCase {
	return &TodoInteraction{
		TodoRepository: repo,
	}
}
