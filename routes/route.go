package routes

import (
	"blog/controller"
	"blog/database"
	"blog/middleware"
	"blog/models"
	"blog/service"
	"blog/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	r := echo.New()
	r.Validator = &utils.CustomValidator{Validator: validator.New()}

	authGroup := r.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware)
	// defer db.Close()
	// Auto Migrate
	database.DbGorm().AutoMigrate(&models.User{})

	userService := service.NewUserService(database.DbGorm())
	userController := controller.NewUserController(userService)

	// Menggunakan middleware untuk menyimpan koneksi database dalam konteks Echo
	authGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", database.DbGorm())
			return next(c)
		}
	})

	authGroup.POST("/upload-picture", userController.UploadPicture)

	// authGroup := r.Group("/auth")
	// authGroup.Use(AuthMiddleware)

	// // Rute untuk mendapatkan profil berdasarkan id_user
	// // Kelompok rute yang memerlukan Basic Authentication
	authorGroup := r.Group("/author")
	authorGroup.Use(middleware.AuthorMiddleware)

	authGroup.GET("/profile", controller.GetSpecProfile)
	authGroup.PUT("/profile/update", controller.ProfileUpdate)

	r.POST("/login", controller.Login)
	r.POST("/register-reader", controller.RegisterReader)
	r.POST("/register-author", controller.RegisterAuthor)
	//profile users
	// r.GET("/profile/:id", controller.GetSpecProfile)
	authGroup.PUT("/password/change", controller.PasswordUpdate)

	// all about contents
	r.GET("/contents", controller.GetAllContent)
	r.GET("/content/:id", controller.GetSpecContent)
	authorGroup.GET("/content/mycontent", controller.GetMyContent)
	authorGroup.POST("/content/create", controller.CreateContent)
	authorGroup.PUT("/content/update/:id", controller.ContentUpdate)
	authorGroup.DELETE("/content/delete/:id", controller.ContentDelete)

	//change password

	r.GET("/categories", controller.GetAllCategory)

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AdminMiddleware)

	//all about category
	r.GET("/category/:id", controller.GetSpecCategory)
	r.GET("/category-content/:id", controller.GetContentCategory)
	adminGroup.POST("/category/create", controller.CategoryAdd)
	adminGroup.PUT("/category/update/:id", controller.CategoryUpdate)
	adminGroup.DELETE("/category/delete/:id", controller.CategoryDelete)

	//about admin permission to manage users
	adminGroup.GET("/users", controller.GetAllUser)
	adminGroup.PUT("/user/update/:id", controller.UserUpdate)
	adminGroup.PUT("user/update-role/:id", controller.UserRoleUpdate)
	adminGroup.DELETE("/user/delete/:id", controller.UserDelete)

	adminGroup.POST("/logout", controller.EchoHandleLogout)
	authorGroup.POST("/logout", controller.EchoHandleLogout)
	r.Static("/picture", "E:/golang/blog/picture")
	return r
}
