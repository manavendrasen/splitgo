package repository

import (
	"payment-service/db"
	"payment-service/model"
)

func ViewPayment() (model.Payment, error)  {
	db := db.GetDB()

	sqlStatement := `SELECT * FROM Payment`

	var payment model.Payment
	err := db.QueryRow(sqlStatement).Scan(&payment) 

	if err != nil {
		return model.Payment{}, err
	}

	return payment, nil
}