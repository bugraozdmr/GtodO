package routes

import (
	"gtodo/internal/app/delivery"
	"gtodo/internal/server"
)

type UserRoutesstruct struct {
	Server *server.ServerStruct
	User   delivery.UserUserCase
}

func (u *UserRoutesstruct) UserRoutes() {

	u.Server.Engine.POST("/signup", u.User.RegisterUserHandler)
	u.Server.Engine.POST("/login", u.User.LoginUserHandler)
}

func NewUserInit(server *server.ServerStruct, user delivery.UserUserCase) *UserRoutesstruct {
	return &UserRoutesstruct{
		Server: server,
		User:   user,
	}
}
