package repository

import (
	"payment-service/database"
	"payment-service/model"
)

func ViewPayment(user model.User) (*model.Payment, error)  {
	db := database.GetDB()

	var payment model.Payment
	result := db.Where("from = ?", user.ID).Find(&payment)

	if result.Error != nil {
		return nil, result.Error
	}

	return &payment, nil
}