package routes

import (
	"blog/controller"
	"blog/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	r := echo.New()
	r.Validator = &utils.CustomValidator{Validator: validator.New()}

	//authentication login and register
	r.POST("/login", controller.Login)
	r.POST("/register-reader", controller.RegisterReader)
	r.POST("/register-author", controller.RegisterAuthor)

	//profile users
	r.GET("/profile/:id", controller.GetSpecProfile)
	r.PUT("/profile/update/:id", controller.ProfileUpdate)

	// all about contents
	r.GET("/contents", controller.GetAllContent)
	r.GET("/content/:id", controller.GetSpecContent)

	//change password
	r.PUT("/password/change/:id", controller.PasswordUpdate)
	r.GET("/categories", controller.GetAllCategory)

	//all about category
	r.GET("/category/:id", controller.GetSpecCategory)
	r.POST("/category/create", controller.CategoryAdd)
	r.PUT("/category/update/:id", controller.CategoryUpdate)
	r.DELETE("/category/delete/:id", controller.CategoryDelete)

	//about admin permission to manage users
	r.GET("/users", controller.GetAllUser)
	r.PUT("/user/update/:id", controller.UserUpdate)
	r.PUT("user/update-role/:id", controller.UserRoleUpdate)
	return r
}
