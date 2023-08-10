package routes

import (
	"blog/controller"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	r := echo.New()

	r.POST("/login", controller.Login)
	return r
}
