package routes

import (
	"blog/controller"
	"blog/middleware"
	"blog/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	r := echo.New()
	r.Validator = &utils.CustomValidator{Validator: validator.New()}

	// authGroup := r.Group("/auth")
	// authGroup.Use(AuthMiddleware)

	// // Rute untuk mendapatkan profil berdasarkan id_user
	// authGroup.GET("/profile/:id", controller.GetSpecProfile)
	// // Kelompok rute yang memerlukan Basic Authentication
	authorGroup := r.Group("/author")
	authorGroup.Use(middleware.AuthorMiddleware)

	authGroup := r.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware)

	r.POST("/login", controller.Login)
	r.POST("/register-reader", controller.RegisterReader)
	r.POST("/register-author", controller.RegisterAuthor)
	//profile users
	// r.GET("/profile/:id", controller.GetSpecProfile)
	// r.PUT("/profile/update/:id", controller.ProfileUpdate)

	// all about contents
	r.GET("/contents", controller.GetAllContent)
	r.GET("/content/:id", controller.GetSpecContent)
	authGroup.POST("/content/create", controller.CreateContent)
	authorGroup.PUT("/content/update/:id", controller.ContentUpdate)
	authorGroup.DELETE("/content/delete/:id", controller.ContentDelete)

	//change password
	authorGroup.PUT("/password/change/:id", controller.PasswordUpdate)

	r.GET("/categories", controller.GetAllCategory)

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AdminMiddleware)
	//all about category
	r.GET("/category/:id", controller.GetSpecCategory)
	adminGroup.POST("/category/create", controller.CategoryAdd)
	adminGroup.PUT("/category/update/:id", controller.CategoryUpdate)
	adminGroup.DELETE("/category/delete/:id", controller.CategoryDelete)

	//about admin permission to manage users
	adminGroup.GET("/users", controller.GetAllUser)
	adminGroup.PUT("/user/update/:id", controller.UserUpdate)
	adminGroup.PUT("user/update-role/:id", controller.UserRoleUpdate)
	adminGroup.DELETE("/user/delete/:id", controller.UserDelete)
	return r
}
