package auth_service

import "flash_sale/models"

type Auth struct {
	Username string
	Password string
	UserID   uint
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) GetUserID() (uint, error) {
	return models.GetUserID(a.Username)
}
