package models

import (
	"github.com/dfaw20/backend-ai-plot/entities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `gorm:"unique;not null"`
	DisplayName string
}

func (m *User) ToPlayer() entities.Player {
	return entities.Player{ID: m.ID, DisplayName: m.DisplayName}
}
