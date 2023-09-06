package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserToken struct {
	gorm.Model

	AccessToken  string    `gorm:"unique;not null"`
	TokenType    string    `gorm:"not null"`
	RefreshToken string    `gorm:"unique;not null"`
	Expiry       time.Time `gorm:"not null"`
}
