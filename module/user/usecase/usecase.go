package user

import (
	"go-clean-arch/model"
	user "go-clean-arch/module/user/interface"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userUsecase struct {
	userHandler user.Handler
}

func NewUserUsecase(us user.Handler) user.Usecase {
	return &userUsecase{us}
}

func (uc *userUsecase) GetUsers(c echo.Context) error {
	var users []*model.User

	u, err := uc.userHandler.Get(users)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userUsecase) GetUsersByEmail(c echo.Context) error {
	var u []*model.User

	u, err := uc.userHandler.GetUserByEmail(u, c.Param("email"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userUsecase) GetUserById(c echo.Context) error {
	var user *model.User
	var id, _ = strconv.Atoi(c.Param("id"))

	u, err := uc.userHandler.GetUserById(user, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *userUsecase) UpdateUser(c echo.Context) error {
	var user *model.User
	var id, _ = strconv.Atoi(c.Param("id"))

	u, err := uc.userHandler.UpdateUserById(user, id, c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
