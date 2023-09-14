package repositories

import (
	"github.com/dfaw20/backend-ai-plot/entities"
)

type PlayerRepository struct {
	userRepository UserRepository
}

func NewPlayerRepository(userRepository UserRepository) PlayerRepository {
	return PlayerRepository{userRepository: userRepository}
}

func (r *PlayerRepository) FindByPlayerID(playerID uint) (entities.Player, error) {
	user, err := r.userRepository.FindByUserID(playerID)
	return entities.ToPlayer(user), err
}
