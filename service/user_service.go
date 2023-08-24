package service

import (
	"blog/models"
	"fmt"
	"strconv"
)

func UsersAll() ([]models.Users, error) {
	usersGet := []models.Users{}

	rows, err := db.Queryx("SELECT u.id, u.username, u.email, r.role FROM users u JOIN users_roles ur ON u.id = ur.user_id JOIN roles r ON ur.role_id = r.id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		place := models.Users{}
		rows.StructScan(&place)
		usersGet = append(usersGet, place)
	}
	return usersGet, nil
}

// Service untuk mengedit data pengguna (users)
func EditUser(editUser models.Users, id int) (models.Users, error) {
	idStr := strconv.Itoa(id)

	_, err := db.NamedExec("UPDATE users SET username = :username, email = :email, password = :password WHERE id = :id", map[string]interface{}{
		"username": editUser.Username,
		"password": editUser.Password,
		"email":    editUser.Email,
		"id":       idStr,
	})
	if err != nil {
		fmt.Println("Error updating user:", err)
		return models.Users{}, err
	}

	return editUser, nil
}

// Service untuk mengedit peran pengguna (roles)
func EditUserRole(userID int, roleID int) error {
	idStr := strconv.Itoa(userID)
	_, err := db.Exec("UPDATE users_roles SET role_id = $1 WHERE user_id = $2", roleID, idStr)
	if err != nil {
		fmt.Println("Error updating user role:", err)
		return err
	}

	return nil
}

// func EditUser(editUser models.Users, id int) (models.Users, error) {
// 	idStr := strconv.Itoa(id)

// 	_, err := db.NamedExec("UPDATE users SET username = :username, email = :email, password = :password WHERE id = :id", map[string]interface{}{
// 		"username": editUser.Username,
// 		"password": editUser.Password,
// 		"email":    editUser.Email,
// 		"id":       idStr,
// 	})
// 	if err != nil {
// 		fmt.Println("Error updating user:", err)
// 		return models.Users{}, err
// 	}
// 	_, errRole := db.Exec("UPDATE users_roles SET role_id = ? WHERE user_id = ?", editUser.Role, id)

//		if errRole != nil {
//			return models.Users{}, errRole
//		}
//		return editUser, nil
//	}
// func EditUser(editUser models.Users, id int, roleID int) (models.Users, error) {
// 	idStr := strconv.Itoa(id)

// 	_, err := db.NamedExec("UPDATE users SET username = :username, email = :email, password = :password WHERE id = :id", map[string]interface{}{
// 		"username": editUser.Username,
// 		"password": editUser.Password,
// 		"email":    editUser.Email,
// 		"id":       idStr,
// 	})
// 	if err != nil {
// 		fmt.Println("Error updating user:", err)
// 		return models.Users{}, err
// 	}

// 	_, errRole := db.Exec("UPDATE users_roles SET role_id = $1 WHERE user_id = $2", roleID, id) // Menggunakan parameter $1 dan $2
// 	if errRole != nil {
// 		fmt.Println("Error updating user role:", errRole)
// 		return models.Users{}, errRole
// 	}

// 	return editUser, nil
// }

// func getRoleIDByName(roleName string) (int, error) {
// 	// Lakukan query ke tabel roles berdasarkan nama peran
// 	var roleID int
// 	err := db.Get(&roleID, "SELECT id FROM roles WHERE role = ?", roleName)
// 	return roleID, err
// }

func DeleteUser(deleteUser models.Users, id int) (models.Users, error) {
	idStr := strconv.Itoa(id)

	_, err := db.Exec("DELETE FROM users WHERE id = $1", idStr)
	if err != nil {
		return deleteUser, err
	}
	return deleteUser, nil
}
