package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Description string
	Amount      float32
	FromUserId  uint
	// ToUserId    uint
	To string `gorm:"not null"`
}
