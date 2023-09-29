package controller

import (
	"blog/models"
	"blog/service"
	"database/sql"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}
type ProfileNotFoundError struct {
	Message string
}

func (e *ProfileNotFoundError) Error() string {
	return e.Message
}

func ProfileUpdate(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User

	var UserProfile models.Profile
	//c.Bind(&UserProfile)
	if err := c.Bind(&UserProfile); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Format JSON tidak valid!",
			Status:  false,
		})
	}
	err := c.Validate(&UserProfile)

	if err != nil {
		// Handle error validasi
		log.Println(err)
		// Kesalahan umum
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Gagal memperbarui profil!",
			Status:  false,
		})
	}

	_, updateErr := service.EditProfile(UserProfile, idUser)
	if updateErr != nil {
		log.Println(updateErr)
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server saat memperbarui profil. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, &models.RegisterResp{
		Code:    200,
		Message: "Profil berhasil disimpan!",
		Status:  true,
	})
}

func GetSpecProfile(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User dari konteks

	var userProfile models.Profile

	userProfile, err := service.GetProfile(idUser)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, &models.Response{
				Code:    404,
				Message: "Halaman profil tidak ditemukan",
				Status:  false,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, &models.Response{
				Code:    500,
				Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
				Status:  false,
			})
		}
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
		return c.JSON(http.StatusUnprocessableEntity, &models.Response{
			Code:    422,
			Message: "Invalid data! Password tidak boleh kosong!",
			Status:  false,
		})
	}

	errS := service.EditPassword(passUpdate, idUser)
	if errS != nil {

		if validationErr, ok := errS.(*service.ValidationError); ok {
			if validationErr.Tag == "strong_password" {
				return c.JSON(http.StatusUnprocessableEntity, &models.RegisterResp{
					Code:    422,
					Message: "Password harus memiliki setidaknya 8 karakter",
					Status:  false,
				})
			}
		}

		if passUpdate.OldPassword == passUpdate.NewPassword {
			return c.JSON(http.StatusBadRequest, &models.Response{
				Code:    400,
				Message: "Password baru tidak boleh sama dengan password lama!",
				Status:  false,
			})
		}
		return c.JSON(http.StatusUnauthorized, &models.Response{
			Code:    401,
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
	userID := e.Get("id_user").(int) // Mengambil ID User

	// Menerima berkas yang diunggah
	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid! Data yang dimasukkan harus bertipe image!",
			Status:  false,
		})
	}

	// Simpan gambar dengan nama unik di server
	pathGambar := "E:/golang/blog/picture/" + file.Filename
	if err := saveUploadedFile(file, pathGambar); err != nil {
		return e.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada saat menyimpan gambar. Coba lagi nanti!",
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
