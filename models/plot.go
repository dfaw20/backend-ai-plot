package models

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name string `json:"name"`
}

type PlotEvent struct {
	gorm.Model
	PlotID  uint `json:"plot_id"`
	EventID uint `json:"event_id"`
}

type Support struct {
	gorm.Model
	PlotID      uint `json:"plot_id"`
	CharacterID uint `json:"character_id"`
}

type Plot struct {
	gorm.Model
	UserID       uint    `json:"user_id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Prompt       string  `json:"prompt"`
	Location     string  `json:"location"`
	Season       string  `json:"season"`
	Genre        string  `json:"genre"`
	OutputFormat string  `json:"output_format"`
	ShowWarning  bool    `json:"show_warning"`
	Supports     []Event `gorm:"many2many:supports;" json:"-"`
	Events       []Event `gorm:"many2many:plot_events;" json:"-"`
}
