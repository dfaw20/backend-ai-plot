package models

import (
	"gorm.io/gorm"
)

type SensitiveOption string

const (
	HIDE              = Gender("HIDE")
	VIEW_WITH_CURTAIN = SensitiveOption("VIEW_WITH_CURTAIN")
	VIEW_DIRECT       = Gender("VIEW_DIRECT")
)

type User struct {
	gorm.Model
	Email           string `gorm:"unique;not null"`
	DisplayName     string
	SensitiveOption string `gorm:"not null"`
}
