package services

import (
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/repository"

	"github.com/go-redis/redis/v8"
)

type PlayerService struct {
	PlayerRepo *repository.PlayerRepository
	Redis      *redis.Client
}

func NewPlayerService(
	playerRepo *repository.PlayerRepository,
	redis *redis.Client) *PlayerService {
	return &PlayerService{
		PlayerRepo: playerRepo,
		Redis:      redis,
	}
}

func (s *PlayerService) GetPlayers(filter string) ([]models.Player, error) {
	return s.PlayerRepo.FindAll(filter)
}

func (s *PlayerService) GetPlayerDetail(playerID uint) (*models.Player, error) {
	return s.PlayerRepo.FindByID(playerID)
}
