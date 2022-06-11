package models

import (
	"golang.org/x/crypto/bcrypt"
	"share/share-api/common/errors"
)

type Auth struct {
	ID       int    //`gorm:"primary_key" json:"id"`
	Username string //`valid:"Required; MaxSize(50)" json:"username"`
	Password string //`valid:"Required; MaxSize(50)" json:"password"`
}

func ValidateAuthentication(username string, password string) (bool, error) {
	var auth Auth

	err := db.Table("auth").Where(Auth{Username: username}).First(&auth).Error

	if err != nil && err.Error() != "record not found" {
		return false, err
	}

	if auth.ID > 0 {
		//bcrypt check password
		if bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password)) != nil {
			return false, nil
		}
		return true, nil
	}

	return false, nil
}

func CreateUser(username string, password string) (bool, error) {
	var auth Auth

	err := db.Table("auth").Where(Auth{Username: username}).First(&auth).Error
	if err != nil && err.Error() != "record not found" {
		return false, err
	}

	if auth.ID > 0 {
		return false, &errors.UserAlreadyExistsError{}
	}

	bcryptPass, bcerr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if bcerr != nil {
		return false, bcerr
	}

	auth = Auth{Username: username, Password: string(bcryptPass)}

	err = db.Table("auth").Create(&auth).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
