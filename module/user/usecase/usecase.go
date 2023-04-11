package user

import (
	"go-clean-arch/model"
	user "go-clean-arch/module/user/interface"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userInteractor user.Handler
}

func NewUserUsecase(us user.Handler) user.Usecase {
	return &userController{us}
}

func (uc *userController) GetUsers(c echo.Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) GetUsersByEmail(c echo.Context) error {
	var u []*model.User

	u, err := uc.userInteractor.GetUserByEmail(u, c.Param("email"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) GetUserById(c echo.Context) error {
	var u []*model.User
	var id, _ = strconv.Atoi(c.Param("id"))

	u, err := uc.userInteractor.GetUserById(u, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	var u []*model.User
	var id, _ = strconv.Atoi(c.Param("id"))

	u, err := uc.userInteractor.UpdateUserById(u, id, c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
