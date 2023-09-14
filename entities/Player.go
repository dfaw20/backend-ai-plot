package entities

import "github.com/dfaw20/backend-ai-plot/models"

type Player struct {
	ID          uint
	DisplayName string
}

func ToPlayer(user models.User) Player {
	return Player{ID: user.ID, DisplayName: user.DisplayName}
}
