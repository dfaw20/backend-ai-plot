package models

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name string
}

type PlotEvent struct {
	gorm.Model
	PlotID  uint
	EventID uint
}

type Support struct {
	gorm.Model
	PlotID      uint
	CharacterID uint
}

type Plot struct {
	gorm.Model
	UserID       uint
	Title        string
	Description  string
	Prompt       string
	Location     string
	Season       string
	Genre        string
	OutputFormat string
	ShowWarning  bool
	Supports     []Event `gorm:"many2many:supports;" json:"-"`
	Events       []Event `gorm:"many2many:plot_events;" json:"-"`
}
