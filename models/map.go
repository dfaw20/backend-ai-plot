package models

import (
	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	UserID uint
	Name   string
}
