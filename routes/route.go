package routes

import (
	"blog/controller"

	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	r := echo.New()

	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.PUT("/profile/update/:id", controller.ProfileUpdate)
	r.GET("/profile/:id", controller.GetSpecProfile)
	r.GET("/contents", controller.GetAllContent)
	r.GET("/content/:id", controller.GetSpecContent)
	r.PUT("/password/change/:id", controller.PasswordUpdate)
	r.GET("/categories", controller.GetAllCategory)
	r.GET("/category/:id", controller.GetSpecCategory)
	r.POST("/category/create", controller.CategoryAdd)
	return r
}
