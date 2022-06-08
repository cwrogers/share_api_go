package services

import "share/share-api/models"

type Auth struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func (a *Auth) Validate() (bool, error) {
	return models.ValidateAuthentication(a.Username, a.Password)
}
