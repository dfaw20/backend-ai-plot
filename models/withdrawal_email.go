package models

import (
	"gorm.io/gorm"
)

type WithdrawalEmail struct {
	gorm.Model
	Email string `gorm:"not null"`
}
