package repository

import (
	"errors"
	"payment-service/database"
	"payment-service/model"
)

func SignUp(user *model.User) (err error) {
	db := database.GetDB()
	result := db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Login(email string) (*model.User, error) {
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
