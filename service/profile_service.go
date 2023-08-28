package service

import (
	"blog/models"
	"context"
	"errors"
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

// service
func EditPassword(ctx context.Context, changePassword models.ChangePasswordRequest, id int) error {
	idStr := strconv.Itoa(id)

	// Ambil password lama dari database
	var dbPassword string
	err := db.GetContext(ctx, &dbPassword, "SELECT password FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	// Periksa apakah password lama sesuai dengan yang ada di database
	errBycript := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(changePassword.OldPassword))
	if errBycript != nil {
		return errBycript
	}

	// Periksa apakah password baru sama dengan password lama
	if changePassword.OldPassword == changePassword.NewPassword {
		return errors.New("new password must be different from old password")
	}

	// Hash password baru sebelum menyimpannya
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, errP := db.NamedExecContext(ctx, "UPDATE users SET password = :password WHERE id = :id", map[string]interface{}{
		"password": hashedPassword,
		"id":       idStr,
	})

	if errP != nil {
		return errP
	}
	return nil
}

// func EditPassword(ctx context.Context, changePassword models.ChangePasswordRequest, id int) error {
// 	idStr := strconv.Itoa(id)

// 	// Ambil password lama dari database
// 	var dbPassword string
// 	err := db.Get(&dbPassword, "SELECT password FROM users WHERE id = $1", id)
// 	if err != nil {
// 		return err
// 	}

// 	// Periksa apakah password lama sesuai dengan yang ada di database
// 	if errBycript := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(changePassword.OldPassword)); err != nil {
// 		return errBycript
// 	}

// 	// Hash password baru sebelum menyimpannya
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	_, errP := db.NamedExec("UPDATE users SET password = :password WHERE id = :id", map[string]interface{}{
// 		"password": hashedPassword,
// 		"id":       idStr,
// 	})

// 	if errP != nil {
// 		return errors.New("failed to update password")
// 	}
// 	return nil
// }
