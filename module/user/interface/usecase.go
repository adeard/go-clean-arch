package user

import (
	"github.com/labstack/echo/v4"
)

// News Repository
type Usecase interface {
	GetUsers(c echo.Context) error
	GetUsersByEmail(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateUser(c echo.Context) error
	CreateUser(c echo.Context) error
}
