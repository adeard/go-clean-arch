package router

import (
	"go-clean-arch/interface/controller"
	"go-clean-arch/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(log.MiddlewareLogging)

	e.GET("/users", func(context echo.Context) error { return c.GetUsers(context) })

	return e
}
