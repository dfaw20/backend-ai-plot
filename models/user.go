package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null" json:"email"`
	DisplayName string `json:"display_name"`
}
