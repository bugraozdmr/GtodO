package di

import (
	"gtodo/internal/app/delivery"
	"gtodo/internal/app/repository"
	"gtodo/internal/app/usecase"
	"gtodo/internal/config"
	db "gtodo/internal/database"
	"gtodo/internal/routes"
	"gtodo/internal/server"
)

func Init(conf config.Config) *server.ServerStruct {
	server := server.NewHTTPServer()
	database := db.ConnectPGDB(conf)

	// ---------------- User Service ----------------
	userRepo := repository.NewUserRepository(database)
	userUseCase := usecase.UseCase(userRepo)
	userDelivery := delivery.UserDelivery(userUseCase)
	userRoutes := routes.NewUserInit(server, userDelivery)
	userRoutes.UserRoutes()

	// ---------------- Todo Service ----------------
	todoRepo := repository.NewTodoRepository(database)
	todoUseCase := usecase.UseCaseTodo(todoRepo)
	todoDelivery := delivery.TodoDelivery(todoUseCase)
	todoRoutes := routes.NewTodoInit(server, todoDelivery)
	todoRoutes.TodoRoutes()

	return server
}
