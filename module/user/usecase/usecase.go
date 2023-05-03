package user

import (
	"go-clean-arch/model"
	user "go-clean-arch/module/user/interface"

	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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

func (uc *userUsecase) CreateUser(c echo.Context) error {
	var user *model.User
	var userNew = new(model.User)

	userNew.Email = c.FormValue("email")
	userNew.Nama = c.FormValue("nama")

	validate := validator.New()

	err := validate.Struct(userNew)
	if err != nil {

		var result = new(model.UserDetail)

		result.Status = http.StatusBadRequest

		if _, ok := err.(*validator.InvalidValidationError); ok {
			result.ErrorMsg = err.Error()
		}

		for _, err := range err.(validator.ValidationErrors) {
			result.ErrorMsg = err.Field() + " is " + err.Tag() + " and must be " + err.Kind().String()
		}

		return c.JSON(http.StatusBadRequest, result)
	}

	u, err := uc.userHandler.CreateUser(user, c)
	if err != nil {

		var result = new(model.UserDetail)

		result.Status = http.StatusBadRequest
		result.ErrorMsg = err.Error()

		return c.JSON(http.StatusBadRequest, result)
	}

	return c.JSON(http.StatusOK, u)
}
