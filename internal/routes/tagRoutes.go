package routes

import (
	"gtodo/internal/app/delivery"
	"gtodo/internal/app/middleware"
	"gtodo/internal/server"
)

type TagRoutesstruct struct {
	Server *server.ServerStruct
	Tag    delivery.TagUserCase
}

func (u *TagRoutesstruct) TagRoutes() {
	// u.Server.Engine.Use(middleware.AuthMiddleware) --
	u.Server.Engine.POST("/tags", u.Tag.CreateTagHandler, middleware.AuthMiddleware)
	u.Server.Engine.GET("/tags", u.Tag.GetAllTagsHandler)
	u.Server.Engine.DELETE("/tags/:id", u.Tag.DeleteTagHandler, middleware.AuthMiddleware)
}

func NewTagInit(server *server.ServerStruct, tag delivery.TagUserCase) *TagRoutesstruct {
	return &TagRoutesstruct{
		Server: server,
		Tag:    tag,
	}
}
