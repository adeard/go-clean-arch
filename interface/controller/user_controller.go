package controller

import (
	"go-clean-arch/domain/model"
	"go-clean-arch/usecase/interactor"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c echo.Context) error
	GetUsersByEmail(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateUser(c echo.Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
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
