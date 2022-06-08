package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func ValidateAuthentication(username string, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}