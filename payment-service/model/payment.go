package model

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Id         int `gorm:"primarykey"`
	Amount     float32
	FromUserId int
	ToUserId   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
