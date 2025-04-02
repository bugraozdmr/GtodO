package repository

import (
	todo "gtodo/internal/app/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo *todo.Todo) error
	FindTodoById(id string) (*todo.Todo, error)
	UpdateTodo(todo *todo.Todo) error
	DeleteTodo(id string) error
	GetAllTodos() ([]todo.Todo, error)
}

type TodoDataBaseInteraction struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoDataBaseInteraction{
		DB: db,
	}
}

func (r *TodoDataBaseInteraction) CreateTodo(todo *todo.Todo) error {
	if err := r.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoDataBaseInteraction) FindTodoById(id string) (*todo.Todo, error) {
	var todo todo.Todo
	if err := r.DB.First(&todo, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoDataBaseInteraction) UpdateTodo(todo *todo.Todo) error {
	if err := r.DB.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoDataBaseInteraction) DeleteTodo(id string) error {
	if err := r.DB.Delete(&todo.Todo{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *TodoDataBaseInteraction) GetAllTodos() ([]todo.Todo, error) {
	var todos []todo.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
