package models

import (
	"github.com/jinzhu/gorm"
)

type Role string

const (
	HERO    = Gender("i")
	PARTNER = Gender("u")
	SUPPORT = Gender("support")
)

type StoryCharacter struct {
	gorm.Model
	StoryID     uint   `json:"story_id"`
	CharacterID uint   `json:"character_id"`
	Role        string `json:"role"`
}

type Story struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	EventID    uint        `json:"event_id"`
	Prompt     string      `json:"prompt"`
	Text       string      `json:"text"`
	Characters []Character `gorm:"many2many:story_characters;" json:"-"`
}
