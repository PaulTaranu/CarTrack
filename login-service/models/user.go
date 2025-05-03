package models

import (
	"github.com/PaulTaranu/CarTrack/login-service/config"
	"github.com/labstack/gommon/log"
)

type User struct {
	ID       string
	Email    string
	Password string
}

func FindUsersByEmail(email string) (*User, error) {
	var user User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Error("User with email {} not found", email)
	}
	return &user, err
}
