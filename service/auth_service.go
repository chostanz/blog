package service

import (
	"blog/database"
	"blog/models"
	"encoding/base64"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var db *sqlx.DB = database.Koneksi()

type ValidationError struct {
	Message string
	Field   string
	Tag     string
}

func (ve *ValidationError) Error() string {
	return ve.Message
}

func CheckCredential(userLogin models.LoginParam) (int, bool, int, error) {
	var isAuthentication bool
	var id int
	var roleID int

	rows, err := db.Query("SELECT CASE WHEN COUNT(*) > 0 THEN 'true' ELSE 'false' END FROM users WHERE username = $1 AND password = $2", userLogin.Username, userLogin.Password)
	if err != nil {
		return 0, false, 0, err
	}

	defer rows.Close()

	rows, err = db.Query("SELECT id, password from users where username = $1", userLogin.Username)
	if err != nil {
		fmt.Println("Error querying users:", err)
		return 0, false, 0, err
	}

	defer rows.Close()

	var dbPasswordBase64 string
	if rows.Next() {
		err = rows.Scan(&id, &dbPasswordBase64)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return 0, false, 0, err
		}
		dbPassword, errBycript := base64.StdEncoding.DecodeString(dbPasswordBase64)

		if errBycript != nil {
			fmt.Println("Password comparison failed:", errBycript)
			return 0, false, 0, errBycript
		}
		errBycript = bcrypt.CompareHashAndPassword(dbPassword, []byte(userLogin.Password))
		if errBycript != nil {
			fmt.Println("Password comparison failed:", errBycript)
			return 0, false, 0, errBycript
		}
		isAuthentication = true
	}

	if isAuthentication {
		rows, err = db.Query("SELECT role_id FROM users_roles WHERE user_id = $1", id)
		if err != nil {
			fmt.Println("Error querying user roles:", err)
			return 0, false, 0, err
		}
		defer rows.Close()

		if rows.Next() {
			err = rows.Scan(&roleID)
			if err != nil {
				fmt.Println("Error scanning role row:", err)
				return 0, false, 0, err
			}
		}
		return id, isAuthentication, roleID, nil
	}
	return 0, false, 0, nil // Jika tidak ada authentikasi yang berhasil

}

func RegisterReader(userRegister models.RegisterParam) error {
	if len(userRegister.Password) < 8 {
		return &ValidationError{
			Message: "Password should be of 8 characters long",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPasswordStr := base64.StdEncoding.EncodeToString(hashedPassword)
	fmt.Println(hashedPassword)
	fmt.Println(hashedPasswordStr)

	_, errInsert := db.NamedExec("INSERT INTO users (username, email, password) VALUES (:username, :email, :password)", map[string]interface{}{
		"username": userRegister.Username,
		"email":    userRegister.Email,
		"password": string(hashedPasswordStr),
	})

	if errInsert != nil {
		return err
	}

	var userID int
	err = db.Get(&userID, "SELECT id FROM users WHERE username = $1", userRegister.Username)
	if err != nil {
		return err
	}

	// Insert into users_roles table with default role_id
	_, err = db.Exec("INSERT INTO users_roles (role_id, user_id) VALUES ($1, $2)", 3, userID)
	if err != nil {
		return err
	}

	return nil
}

func RegisterAuthor(userRegister models.RegisterParam) error {
	if len(userRegister.Password) < 8 {
		return &ValidationError{
			Message: "Password should be of 8 characters long",
			Field:   "password",
			Tag:     "strong_password",
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashedPasswordStr := base64.StdEncoding.EncodeToString(hashedPassword)
	fmt.Println(hashedPassword)
	fmt.Println(hashedPasswordStr)

	_, errInsert := db.NamedExec("INSERT INTO users (username, email, password) VALUES (:username, :email, :password)", map[string]interface{}{
		"username": userRegister.Username,
		"email":    userRegister.Email,
		"password": hashedPasswordStr,
	})

	if errInsert != nil {
		return errInsert
	}

	var userID int
	err = db.Get(&userID, "SELECT id FROM users WHERE username = $1", userRegister.Username)
	if err != nil {
		return err
	}

	// Insert into users_roles table with default role_id
	_, err = db.Exec("INSERT INTO users_roles (role_id, user_id) VALUES ($1, $2)", 2, userID)
	if err != nil {
		return err
	}

	return nil
}
