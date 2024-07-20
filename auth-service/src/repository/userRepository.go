package repository

import (
	"auth-service/src/database"
	"auth-service/src/model"
	"errors"
)

func SignUp(user *model.User) (err error) {
	db := database.GetDB()

	// check if the user is in the db already
	user, error := GetUserByEmail(user.Email)

	if error != nil {
		return error	
	}

	// user exists - email is already used
	if (user.ID != 0) {
		return errors.New("EMAIL_ALREADY_TAKEN")
	}

	// user does not exists
	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Login(email string) (*model.User, error) {
	return GetUserByEmail(email)
}

func GetUserByEmail(email string) (*model.User, error) {
	db := database.GetDB()
	var user model.User
	result := db.Find(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	if user.ID == 0 {
		return nil, errors.New("USER_NOT_FOUND")
	}
	return &user, nil
}