package controller

import (
	"blog/models"
	"blog/service"
	"io"
	"mime/multipart"
	"net/http"
	"os"

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

			return echo.NewHTTPError(http.StatusBadRequest, &models.Response{
				Code:    400,
				Message: "Data invalid!",
				Status:  false,
			})
		}
		return c.JSON(http.StatusCreated, &models.RegisterResp{
			Code:    201,
			Message: "Profil berhasil disimpan!",
			Status:  true,
		})
	}

	return c.JSON(http.StatusBadRequest, &models.Response{
		Code:    400,
		Message: "Gagal mengupdate profil!",
		Status:  false,
	})
}

func GetSpecProfile(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User dari konteks

	var userProfile models.Profile

	userProfile, err := service.GetProfile(idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}
	return c.JSON(http.StatusOK, userProfile)
}

func PasswordUpdate(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User dari konteks

	var passUpdate models.ChangePasswordRequest
	if err := c.Bind(&passUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Invalid request data",
			Status:  false,
		})
	}

	if err := c.Validate(&passUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Validasi error",
			Status:  false,
		})
	}

	errS := service.EditPassword(passUpdate, idUser)
	if errS != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Password lama salah!",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Password berhasil diubah!",
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

func (c *UserController) UploadPicture(e echo.Context) error {
	userID := e.Get("id_user").(int) // Mengambil ID U

	// Menerima berkas yang diunggah
	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid!",
			Status:  false,
		})
	}

	// Simpan gambar dengan nama unik di server
	pathGambar := "E:/golang/blog/picture/" + file.Filename
	if err := saveUploadedFile(file, pathGambar); err != nil {
		return e.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Error menyimpan gambar",
			Status:  false,
		})
	}
	baseURL := "http://localhost:8080"
	pictureURL := baseURL + "/picture/" + file.Filename

	// Memperbarui URL gambar dalam database
	if err := c.userService.UpdatePictureURL(userID, pictureURL); err != nil {
		return e.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Gagal mengupload cover image url",
			Status:  false,
		})

	}

	return e.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Profil berhasil diperbarui!",
		Status:  true,
	})
}
