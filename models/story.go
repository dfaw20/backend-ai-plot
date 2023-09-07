package models

import (
	"github.com/jinzhu/gorm"
)

type StoryCharacter struct {
	gorm.Model
	StoryID     uint `json:"story_id"`
	CharacterID uint `json:"character_id"`
}

type Story struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	EventID    uint        `json:"event_id"`
	Prompt     string      `json:"prompt"`
	Text       string      `json:"text"`
	Characters []Character `gorm:"many2many:story_characters;" json:"-"`
}
