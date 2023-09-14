package models

import (
	"gorm.io/gorm"
)

type SensitiveOption uint

const (
	VIEW_WITH_CURTAIN = SensitiveOption(1)
	VIEW_DIRECT       = SensitiveOption(2)
	HIDE              = SensitiveOption(3)
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique;not null"`
	DisplayName     string
	SensitiveOption uint `gorm:"not null"`
}
