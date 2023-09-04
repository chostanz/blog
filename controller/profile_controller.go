package controller

import (
	"blog/models"
	"blog/service"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func ProfileUpdate(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User

	var UserProfile models.Profile
	c.Bind(&UserProfile)
	err := c.Validate(&UserProfile)

	if err == nil {
		_, registerErr := service.EditProfile(UserProfile, idUser)
		if registerErr != nil {

			return echo.NewHTTPError(http.StatusBadRequest, "raiso")
		}
		return c.JSON(http.StatusCreated, &models.RegisterResp{
			Message: "Berhasil update",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func GetSpecProfile(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User dari konteks

	var userProfile models.Profile

	userProfile, err := service.GetProfile(idUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "tidak ada")
	}
	return c.JSON(http.StatusOK, userProfile)
}

func PasswordUpdate(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User dari konteks

	var passUpdate models.ChangePasswordRequest
	if err := c.Bind(&passUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if err := c.Validate(&passUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error")
	}

	errS := service.EditPassword(passUpdate, idUser)
	if errS != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "An error occurred")
	}

	return c.JSON(http.StatusOK, &models.Response{
		Message: "Password updated successfully",
		Status:  true,
	})

}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func saveUploadedFile(file *multipart.FileHeader, path string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

// func (c *UserController) UploadPicture(e echo.Context) error {
// 	userID := uint(1) // Ganti dengan pengambilan ID dari sesi atau permintaan

// 	// Menerima berkas yang diunggah
// 	file, err := e.FormFile("image")
// 	if err != nil {
// 		return e.String(http.StatusBadRequest, "Error uploading image")
// 	}
// 	db := c.userService.db.Begin() // Mulai transaksi
// 	defer db.Close()               // Tutup transaksi setelah selesai

// 	// Simpan gambar dengan nama unik di server
// 	pathGambar := "E:/golang/blog/picture/" + file.Filename
// 	if err := saveUploadedFile(file, pathGambar); err != nil {
// 		return e.String(http.StatusInternalServerError, "Error saving image")
// 	}

// 	// Membuat URL gambar berdasarkan base URL situs dan path gambar
// 	baseURL := "https://example.com" // Ganti dengan base URL situs Anda
// 	pictureURL := baseURL + "/" + pathGambar

// 	// Memperbarui URL gambar dalam database
// 	if err := c.userService.UpdatePictureURL(db, userID, pictureURL); err != nil {
// 		return e.String(http.StatusInternalServerError, "Failed to update picture URL")
// 	}
// 	db.Commit() // Konfirmasi transaksi jika sukses

// 	return e.String(http.StatusOK, "Picture uploaded and URL updated")
// }

func (c *UserController) UploadPicture(e echo.Context) error {
	userID := e.Get("id_user").(int) // Mengambil ID U

	// Menerima berkas yang diunggah
	file, err := e.FormFile("image")
	if err != nil {
		return e.String(http.StatusBadRequest, "Error uploading image")
	}

	// Simpan gambar dengan nama unik di server
	pathGambar := "E:/golang/blog/picture/" + file.Filename
	if err := saveUploadedFile(file, pathGambar); err != nil {
		return e.String(http.StatusInternalServerError, "Error saving image")
	}
	baseURL := "http://localhost:8080"
	pictureURL := baseURL + "/picture/" + file.Filename

	// Mulai transaksi database
	db := e.Get("db").(*gorm.DB)
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Memperbarui URL gambar dalam database
	if err := c.userService.UpdatePictureURL(userID, pictureURL); err != nil {
		tx.Rollback()
		return e.String(http.StatusInternalServerError, "Failed to update picture URL")
	}

	tx.Commit()

	return e.String(http.StatusOK, "Picture uploaded and URL updated")
}
