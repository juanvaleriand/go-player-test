package services

import (
	"context"
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/repository"
	"go-player-test/internal/api/requests"
	"go-player-test/internal/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	PlayerRepo *repository.PlayerRepository
	Redis      *redis.Client
}

func NewAuthService(
	playerRepo *repository.PlayerRepository,
	redis *redis.Client) *AuthService {
	return &AuthService{
		PlayerRepo: playerRepo,
		Redis:      redis,
	}
}

func (s *AuthService) Register(player *models.Player) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(player.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	player.Password = string(hashedPassword)
	player.CreatedAt = utils.GetCurrentTime()
	player.UpdatedAt = utils.GetCurrentTime()

	return s.PlayerRepo.Create(player)
}

func (s *AuthService) Login(request requests.LoginRequest) (string, error) {
	// Generate a token
	token, err := utils.GenerateToken(request.Username)
	if err != nil {
		return "", err
	}

	// Check if the token is already in Redis
	err = s.Redis.Get(context.Background(), token).Err()
	if err == redis.Nil {
		// Token not found in Redis, proceed with database check

		// Find player by username
		player, err := s.PlayerRepo.FindByUsername(request.Username)
		if err != nil {
			return "", err
		}

		// Compare hashed password
		err = bcrypt.CompareHashAndPassword([]byte(player.Password), []byte(request.Password))
		if err != nil {
			return "", err
		}

		// Set expiration time to 6 hours (6 * 60 * 60 seconds)
		expireTime := 6 * time.Hour

		// Set the token in Redis with an expiration time
		err = s.Redis.Set(context.Background(), token, player.ID, expireTime).Err()
		if err != nil {
			return "", err
		}
		return token, nil
	} else if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Logout(tokenString string) error {
	utils.InvalidateToken(tokenString)
	err := s.Redis.Del(context.Background(), tokenString).Err()
	if err != nil {
		return err
	}

	return nil
}
