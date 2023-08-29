package service

import (
	"blog/models"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func EditProfile(UserProfile models.Profile, id int) (models.Profile, error) {
	idStr := strconv.Itoa(id)

	_, err := db.NamedExec("UPDATE users SET name = :name, bio = :bio, picture_url = :picture_url WHERE id = :id", map[string]interface{}{
		"name":        UserProfile.Name,
		"bio":         UserProfile.Bio,
		"picture_url": UserProfile.Pictureurl,
		"id":          idStr,
	})

	if err != nil {
		return models.Profile{}, err
	}
	return UserProfile, nil
}

func GetProfile(id int) (models.Profile, error) {
	var userProfile models.Profile
	idStr := strconv.Itoa(id)

	err := db.Get(&userProfile, "SELECT id, name, bio, picture_url FROM users WHERE id = $1", idStr)
	if err != nil {
		return models.Profile{}, err
	}
	return userProfile, nil

}

func EditPassword(changePassword models.ChangePasswordRequest, id int) error {
	idStr := strconv.Itoa(id)

	// Ambil password lama dari database
	var dbPassword string
	err := db.Get(&dbPassword, "SELECT password FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	decodedPassword, err := base64.StdEncoding.DecodeString(dbPassword)
	if err != nil {
		return err
	}

	errBycript := bcrypt.CompareHashAndPassword(decodedPassword, []byte(changePassword.OldPassword))
	if errBycript != nil {
		fmt.Println("Error comparing old passwords:", errBycript)
		return errBycript
	}

	if changePassword.OldPassword == changePassword.NewPassword {
		return errors.New("new password must be different from old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hashed password:", err)
		return err
	}

	hashedPasswordStr := base64.StdEncoding.EncodeToString(hashedPassword)
	_, err = db.NamedExec("UPDATE users SET password = :password WHERE id = :id", map[string]interface{}{
		"password": hashedPasswordStr,
		"id":       idStr,
	})

	if err != nil {
		fmt.Println("Error updating password in database:", err)
		return err
	}
	return nil
}
