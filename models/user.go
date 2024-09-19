package models

import (
	"errors"
	"flash_sale/pkg/logging"
	"gorm.io/gorm"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logging.Error("[CheckAuth]User not found")
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetUserID get UserID via username
func GetUserID(username string) (uint, error) {
	var user User
	err := db.Select("id").Where(User{Username: username}).First(&user).Error
	if err != nil {
		logging.Error("[GetUserID]User not found")
		return 0, err
	}

	return user.ID, nil
}
