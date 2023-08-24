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

	authGroup := r.Group("/auth")
	authGroup.Use(AuthMiddleware)

	// Rute untuk mendapatkan profil berdasarkan id_user
	authGroup.GET("/profile/:id", controller.GetSpecProfile)
	// Kelompok rute yang memerlukan Basic Authentication
	authorGroup := r.Group("/author")
	authorGroup.Use(AuthorMiddleware)

	r.POST("/login", controller.Login)
	r.POST("/register-reader", controller.RegisterReader)
	r.POST("/register-author", controller.RegisterAuthor)
	//profile users
	// r.GET("/profile/:id", controller.GetSpecProfile)
	// r.PUT("/profile/update/:id", controller.ProfileUpdate)

	// all about contents
	r.GET("/contents", controller.GetAllContent)
	r.GET("/content/:id", controller.GetSpecContent)
	r.POST("/content/create", controller.CreateContent)

	//change password
	r.PUT("/password/change/:id", controller.PasswordUpdate)
	r.GET("/categories", controller.GetAllCategory)

	//all about category
	r.GET("/category/:id", controller.GetSpecCategory)
	authorGroup.POST("/category/create", controller.CategoryAdd)
	authorGroup.PUT("/category/update/:id", controller.CategoryUpdate)
	authorGroup.DELETE("/category/delete/:id", controller.CategoryDelete)

	adminGroup := r.Group("/admin")
	adminGroup.Use(AdminMiddleware)

	//about admin permission to manage users
	adminGroup.GET("/users", controller.GetAllUser)
	adminGroup.PUT("/user/update/:id", controller.UserUpdate)
	adminGroup.PUT("user/update-role/:id", controller.UserRoleUpdate)
	adminGroup.DELETE("/user/delete/:id", controller.UserDelete)
	return r
}
