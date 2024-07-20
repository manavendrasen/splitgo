package repository

import (
	"errors"
	"payment-service/src/database"
	"payment-service/src/model"
	"strings"
)

func ViewPayment(userID uint) ([]model.Payment, error) {
	db := database.GetDB()

	var payments []model.Payment
	result := db.Where("from_user_id = ?", userID).Find(&payments)

	if result.Error != nil {
		return nil, result.Error
	}

	return payments, nil
}

func AddPayment(userID uint, to string, amount float32, description string) (model.Payment, error) {
	db := database.GetDB()

	payment := model.Payment{
		FromUserId:  userID,
		To:          to,
		Amount:      amount,
		Description: strings.Trim(description, " "),
	}

	result := db.Create(&payment)

	if result.Error != nil {
		return model.Payment{}, result.Error
	}

	return payment, nil
}

func UpdatePayment(userID uint, paymentID uint, amount float32, to string, description string) (model.Payment, error) {
	db := database.GetDB()

	payment := model.Payment{}
	result := db.First(&payment, "ID = ? and from_user_id = ?", paymentID, userID)

	if result.RowsAffected == 0 {
		return model.Payment{}, errors.New("COULD_NOT_FIND_PAYMENT")
	}

	payment.Amount = amount
	payment.To = to
	payment.Description = strings.Trim(description, " ")

	result = db.Save(&payment)

	if result.Error != nil {
		return model.Payment{}, errors.New("COULD_NOT_SAVE")
	}

	result = db.First(&payment, "ID = ? and from_user_id = ?", paymentID, userID)

	if result.Error != nil {
		return model.Payment{}, errors.New("COULD_NOT_FETCH_UPDATED_PAYMENT")
	}

	return payment, nil
}

func DeletePayment(userID uint, paymentID uint) (uint, error) {
	db := database.GetDB()

	payment := model.Payment{}
	result := db.Where("ID = ? and from_user_id = ?", paymentID, userID).Delete(&payment)

	if result.RowsAffected == 0 {
		return 0, errors.New("COULD_NOT_FIND_PAYMENT")
	}

	return uint(result.RowsAffected), nil
}
