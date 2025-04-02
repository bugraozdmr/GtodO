package delivery

import (
	app "gtodo/internal/app/entity"
	"gtodo/internal/app/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	tagUseCase usecase.TagUseCase
}

type TagUserCase interface {
	CreateTagHandler(c echo.Context) error
	GetAllTagsHandler(c echo.Context) error
	DeleteTagHandler(c echo.Context) error
}

func (t *TagHandler) CreateTagHandler(c echo.Context) error {
	var tag app.Tag
	if err := c.Bind(&tag); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})
	}

	ctx := c.Request().Context()
	err := t.tagUseCase.CreateTag(ctx, &tag)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "tag created successfully",
	})
}

func (t *TagHandler) DeleteTagHandler(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	if err := t.tagUseCase.DeleteTag(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "tag deleted successfully",
	})
}

func (t *TagHandler) GetAllTagsHandler(c echo.Context) error {
	ctx := c.Request().Context()
	tags, err := t.tagUseCase.GetAllTags(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "something went wrong",
		})
	}

	return c.JSON(http.StatusOK, tags)
}

func TagDelivery(tagUseCase usecase.TagUseCase) *TagHandler {
	return &TagHandler{
		tagUseCase: tagUseCase,
	}
}
