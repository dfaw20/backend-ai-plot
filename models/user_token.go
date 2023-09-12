package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model

	AccessToken  string    `gorm:"unique;not null" json:"access_token"`
	TokenType    string    `gorm:"not null" json:"token_type"`
	RefreshToken string    `gorm:"not null" json:"refresh_token"`
	Expiry       time.Time `gorm:"not null" json:"expiry"`
}
