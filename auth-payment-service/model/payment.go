package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Amount     float32
	FromUserId int
	ToUserId   int
}
