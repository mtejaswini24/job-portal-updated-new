package repository

import (
	"job-portal-api/internal/models"
	"job-portal-api/internal/pkg"

	"errors"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateUser(userData models.User) (models.User, error) {
	result := r.db.Create(&userData).Error
	//calling default create method
	if result != nil {
		log.Info().Err(result).Send()
		return models.User{}, errors.New("unable to create new user")
	}
	// Successfully created the record, return the user.
	return userData, nil
}
func (r *Repo) UserLogin(email string) (models.User, error) {
	// We attempt to find the User record where the email
	// matches the provided email.
	var user models.User
	result := r.db.Where("email = ?", email).First(&user).Error
	if result != nil {
		log.Info().Err(result).Send()
		return models.User{}, errors.New("email not found")
	}
	return user, nil
}
func (r *Repo) CheckUserEmail(email string) (bool, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return false, errors.New("email not found")
	}
	return true, nil
	// || result.RowsAffected == 0
}

func (r *Repo) UpdatePassword(email, password string) (bool, error) {

	var user models.User
	result := r.db.Where("email=?", email).First(&user)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return false, result.Error
	}
	hashedPass, err := pkg.HashPassword(password)
	if err != nil {
		return false, err
	}
	user.PasswordHash = hashedPass
	if err := r.db.Save(&user).Error; err != nil {
		return false, err
	}

	return false, nil
}
