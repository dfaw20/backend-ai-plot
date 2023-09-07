package models

import (
	"github.com/jinzhu/gorm"
)

type Character struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Gender      string `json:"gender"`
	Outfit      string `json:"outfit"`
	Hairstyle   string `json:"hairstyle"`
	Personality string `json:"personality"`
	Tone        string `json:"tone"`
	Profile     string `json:"profile"`
}