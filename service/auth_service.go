package service

import (
	"blog/database"
	"blog/models"
	"log"
)

func CheckCredential(userLogin models.LoginParam) (int, bool) {
	var isAuthentication bool
	var id int

	database.DB.QueryRow("SELECT IF(COUNT(*), 'true', 'false') FROM users WHERE username = :username AND password = :password", userLogin.Username, userLogin.Password).Scan(&isAuthentication)
	database.DB.QueryRow("SELECT id from users where username = :username AND password = :passowrd", userLogin.Username, userLogin.Password).Scan(&id)
	return id, isAuthentication
}

func RegisterUser(userRegister models.RegisterParam) bool {
	_, err := database.DB.Exec("INSERT INTO users (username, password, email) VALUES (:username, :password, :email)", userRegister)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
