package routes

import (
	"gtodo/internal/app/delivery"
	"gtodo/internal/server"
)

type TodoRoutesstruct struct {
	Server *server.ServerStruct
	Todo   delivery.TodoUseCase
}

func (u *TodoRoutesstruct) TodoRoutes() {
	u.Server.Engine.POST("/todos", u.Todo.CreateTodoHandler)
	u.Server.Engine.GET("/todos", u.Todo.GetAllTodosHandler)
	u.Server.Engine.GET("/todos/:id", u.Todo.GetTodoByIDHandler)
	u.Server.Engine.PUT("/todos/:id", u.Todo.UpdateTodoHandler)
	u.Server.Engine.DELETE("/todos/:id", u.Todo.DeleteTodoHandler)
}

func NewTodoInit(server *server.ServerStruct, todo delivery.TodoUseCase) *TodoRoutesstruct {
	return &TodoRoutesstruct{
		Server: server,
		Todo:   todo,
	}
}
