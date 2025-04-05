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

	userRepo := repository.NewUserRepository(database)
	userUseCase := usecase.UseCase(userRepo)
	userDelivery := delivery.UserDelivery(userUseCase)
	userRoutes := routes.NewUserInit(server, userDelivery)
	userRoutes.UserRoutes()

	todoRepo := repository.NewTodoRepository(database)
	// injection yaptık user'ı todoya
	todoUseCase := usecase.UseCaseTodo(todoRepo,userUseCase)
	todoDelivery := delivery.TodoDelivery(todoUseCase)
	todoRoutes := routes.NewTodoInit(server, todoDelivery)
	todoRoutes.TodoRoutes()

	tagRepo := repository.NewTagRepository(database)
	tagUseCase := usecase.UseCaseTag(tagRepo)
	tagDelivery := delivery.TagDelivery(tagUseCase)
	tagRoutes := routes.NewTagInit(server, tagDelivery)
	tagRoutes.TagRoutes()

	return server
}
