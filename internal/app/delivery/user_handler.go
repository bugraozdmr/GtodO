package delivery

import (
	app "gtodo/internal/app/entity"
	"gtodo/internal/app/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

type UserUserCase interface {
	RegisterUserHandler(c echo.Context) error
	LoginUserHandler(c echo.Context) error
}

func (u *UserHandler) RegisterUserHandler(c echo.Context) error {
	var user app.UserRegister
	if err := c.Bind(&user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad request",
		})

	}
	ctx := c.Request().Context()
	flag, err := u.userUseCase.RegisterUser(ctx, &user)

	if err != nil {
		if flag == 1 {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Something wrong happened",
			})
		} else if flag == 2 {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "created",
	})

}

func (u *UserHandler) LoginUserHandler(c echo.Context) error {
	var userLogin app.UserLogin
	if err := c.Bind(&userLogin); err != nil {

		return c.JSON(400, map[string]string{"Error": "binding error"})

	}

	user, err := u.userUseCase.Login(&userLogin)
	if err != nil {
		return c.JSON(500, map[string]string{"Error": err.Error()})

	}
	response := map[string]interface{}{
		"data": user,
	}
	return c.JSON(200, response)
}

func UserDelivery(userUseCase usecase.UserUseCase) UserUserCase {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}
