package routes

import (
	"gtodo/internal/app/delivery"
	"gtodo/internal/server"
)

type TagRoutesstruct struct {
	Server *server.ServerStruct
	Tag   delivery.TagUserCase
}

func (u *TagRoutesstruct) TagRoutes() {
	u.Server.Engine.POST("/tags", u.Tag.CreateTagHandler)
	u.Server.Engine.GET("/tags", u.Tag.GetAllTagsHandler)
	u.Server.Engine.DELETE("/tags/:id", u.Tag.DeleteTagHandler)
}

func NewTagInit(server *server.ServerStruct, tag delivery.TagUserCase) *TagRoutesstruct {
	return &TagRoutesstruct{
		Server: server,
		Tag:   tag,
	}
}
