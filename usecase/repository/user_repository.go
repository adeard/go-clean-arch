package repository

import (
	"go-clean-arch/domain/model"

	"github.com/labstack/echo"
)

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	GetUserByEmail(u []*model.User, email string) ([]*model.User, error)
	GetUserById(u []*model.User, id int) ([]*model.User, error)
	UpdateUserById(u []*model.User, id int, c echo.Context) ([]*model.User, error)
}
