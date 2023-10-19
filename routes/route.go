package routes

import (
	"blog/controller"
	"blog/database"
	"blog/middleware"
	"blog/service"
	"blog/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	r := echo.New()
	r.Validator = &utils.CustomValidator{Validator: validator.New()}

	db := database.Koneksi()

	// // Auto Migrate
	// database.DbGorm().AutoMigrate(&models.User{})

	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AdminMiddleware)

	//kelompok rute authentikasi
	authGroup := r.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware)

	// // Menggunakan middleware untuk menyimpan koneksi database dalam konteks Echo
	// authGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		c.Set("db", database.DbGorm())
	// 		return next(c)
	// 	}
	// })

	contentService := service.NewContentService(db)
	contentController := controller.NewContentController(contentService)

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.SetContentServiceMiddleware(contentService))

	authorGroup := r.Group("/author")
	authorGroup.Use(middleware.AuthorMiddleware)

	r.POST("/login", controller.Login)
	r.POST("/register-reader", controller.RegisterReader)
	r.POST("/register-author", controller.RegisterAuthor)

	authGroup.GET("/profile", controller.GetSpecProfile)
	authGroup.PUT("/profile/update", controller.ProfileUpdate)
	//rute untuk mengunggah foto proflil
	authGroup.PUT("/upload-picture", userController.UploadPicture)

	authGroup.PUT("/password/change", controller.PasswordUpdate)

	authGroup.GET("/user/:id", controller.GetSpecUser)

	// all about contents
	r.GET("/contents", controller.GetAllContent)
	r.GET("/content/:id", controller.GetSpecContent)
	authorGroup.GET("/content/mycontent", controller.GetMyContent)
	authorGroup.POST("/content/create", controller.CreateContent)
	authorGroup.PUT("/content/update/:id", controller.ContentUpdate)
	apiGroup.PUT("/upload-image/:contentID/cover-image", contentController.UploadCoverImage, middleware.AuthorMiddleware)
	authorGroup.DELETE("/content/delete/:id", controller.ContentDelete)
	adminGroup.DELETE("/content/delete/:id", controller.DeleteContent)

	r.GET("/categories", controller.GetAllCategory)
	r.GET("/category/:id", controller.GetSpecCategory)
	r.GET("/category-content/:id", controller.GetContentCategory)
	adminGroup.POST("/category/create", controller.CategoryAdd)
	adminGroup.PUT("/category/update/:id", controller.CategoryUpdate)
	adminGroup.DELETE("/category/delete/:id", controller.CategoryDelete)

	//about admin permission to manage users
	adminGroup.GET("/users", controller.GetAllUser)
	adminGroup.GET("/roles", controller.GetAllRole)
	//adminGroup.PUT("/user/update/:id", controller.UserUpdate)
	adminGroup.PUT("/user/update-role/:id", controller.UserRoleUpdate)
	adminGroup.DELETE("/user/delete/:id", controller.UserDelete)

	r.POST("/logout", controller.EchoHandleLogout, middleware.LogoutMiddleware)

	r.Static("/picture", "E:/golang/blog/picture")
	r.Static("/cover", "E:/golang/blog/cover")
	return r
}
