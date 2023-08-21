package controller

import (
	"blog/models"
	"blog/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, err := service.UsersAll()

	if err != nil {
		response := models.Response{
			Message: "Halaman tidak ditemukan atau url salah",
			Status:  false,
		}
		return c.JSON(http.StatusNotFound, response)
	}
	return c.JSON(http.StatusOK, users)
}

// func UserUpdate(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	var editUser models.Users
// 	c.Bind(&editUser)
// 	err := c.Validate(&editUser)

// 	// 	if err == nil {
// 	// 		_, editErr := service.EditUser(editUser, id)
// 	// 		if editErr != nil {
// 	// 			fmt.Println("Error editing user:", editErr)
// 	// 			return c.JSON(http.StatusBadRequest, editErr.Error())

// 	// 			//return c.JSON(http.StatusBadRequest, err.Error())
// 	// 		}
// 	// 		return c.JSON(http.StatusOK, &models.Response{
// 	// 			Message: "Berhasil update",
// 	// 			Status:  true,
// 	// 		})
// 	// 	}
// 	// 	fmt.Println("Validation error:", err)

// 	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

// 	// }
// 	if err != nil {
// 		fmt.Println("Validation error:", err)
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	roleID, _ := strconv.Atoi(editUser.Role)             // Konversi role menjadi integer
// 	_, editErr := service.EditUser(editUser, id, roleID) // Mengirim roleID ke service
// 	if editErr != nil {
// 		fmt.Println("Error editing user:", editErr)
// 		return c.JSON(http.StatusBadRequest, editErr.Error())
// 	}

// 	return c.JSON(http.StatusOK, &models.Response{
// 		Message: "Berhasil update",
// 		Status:  true,
// 	})
// }

// Controller untuk mengedit data pengguna
func UserUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editUser models.Users
	c.Bind(&editUser)
	err := c.Validate(&editUser)

	if err != nil {
		fmt.Println("Validation error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, editErr := service.EditUser(editUser, id)
	if editErr != nil {
		fmt.Println("Error editing user:", editErr)
		return c.JSON(http.StatusBadRequest, editErr.Error())
	}

	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil update",
		Status:  true,
	})
}

// Controller untuk mengedit peran pengguna
func UserRoleUpdate(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	roleID, _ := strconv.Atoi(c.Param("role_id")) // Anda bisa menggunakan c.FormValue jika role_id diteruskan melalui form data

	err := service.EditUserRole(userID, roleID)
	if err != nil {
		fmt.Println("Error editing user role:", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil update role",
		Status:  true,
	})
}
