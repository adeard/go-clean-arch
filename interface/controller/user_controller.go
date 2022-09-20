package controller

import (
	"go-clean-arch/domain/model"
	"go-clean-arch/usecase/interactor"
	"net/http"

	"github.com/labstack/echo"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c echo.Context) error
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
