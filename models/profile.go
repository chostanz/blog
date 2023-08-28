package models

type Profile struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name" validate:"required"`
	Bio        string `json:"bio" db:"bio"`
	Pictureurl string `json:"picture_url" db:"picture_url"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}
