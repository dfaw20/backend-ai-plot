package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	HERO    = Gender("i")
	PARTNER = Gender("u")
	SUPPORT = Gender("support")
)

type StoryCharacter struct {
	gorm.Model
	StoryID     uint
	CharacterID uint
	Role        string
}

type Story struct {
	gorm.Model
	UserID     uint
	PlotID     uint
	Prompt     string
	Text       string
	Characters []Character `gorm:"many2many:story_characters;" json:"-"`
}
