package service

import (
	"blog/models"
	"strconv"
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

	err := db.Get(&userProfile, "SELECT name, bio, picture_url FROM users WHERE id = $1", idStr)
	if err != nil {
		return models.Profile{}, err
	}
	return userProfile, nil

}

func EditPassword(changePassword models.Password, id int) (models.Password, error) {
	idStr := strconv.Itoa(id)

	_, err := db.NamedExec("UPDATE users SET password = :password WHERE id = :id", map[string]interface{}{
		"password": changePassword.Password,
		"id":       idStr,
	})

	if err != nil {
		return models.Password{}, err
	}
	return changePassword, nil
}
