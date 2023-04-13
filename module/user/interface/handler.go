package user

import (
	"go-clean-arch/model"

	"github.com/labstack/echo/v4"
)

// News Repository
type Handler interface {
	Get(u []*model.User) (*model.UserList, error)
	GetUserByEmail(u []*model.User, email string) ([]*model.User, error)
	GetUserById(u []*model.User, id int) ([]*model.User, error)
	UpdateUserById(u []*model.User, id int, c echo.Context) ([]*model.User, error)
}
