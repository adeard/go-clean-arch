package user

import (
	"go-clean-arch/log"
	"go-clean-arch/module"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoutes(e *echo.Echo, c module.AppUsecase) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(log.MiddlewareLogging)

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })
	e.GET("/users/:id", func(context echo.Context) error { return c.GetUserById(context) })
	e.PUT("/users/:id", func(context echo.Context) error { return c.UpdateUser(context) })

	return e
}
