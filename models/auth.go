package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	ID       int    //`gorm:"primary_key" json:"id"`
	Username string //`valid:"Required; MaxSize(50)" json:"username"`
	Password string //`valid:"Required; MaxSize(50)" json:"password"`
}

func ValidateAuthentication(username string, password string) (bool, error) {
	var auth Auth

	println("username: " + username)
	println("password: " + password)

	//bcrypt password
	bcryptPass, bcerr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if bcerr != nil {
		return false, bcerr
	}
	println("bcryptPass: " + string(bcryptPass))

	err := db.Table("auth").Where(Auth{Username: username}).First(&auth).Error

	if err != nil && err.Error() != "record not found" {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
