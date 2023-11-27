package models

import "gorm.io/gorm"

type NewUser struct {
	Name     string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	gorm.Model
	Name         string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
type Reset struct {
	Otp             string `json:"otp" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	NewPassword     string `json:"newpassword" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required"`
}
