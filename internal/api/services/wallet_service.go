package services

import (
	"context"
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/repository"
	"go-player-test/internal/utils"

	"github.com/go-redis/redis/v8"
)

type WalletService struct {
	WalletRepo *repository.WalletRepository
	Redis      *redis.Client
}

func NewWalletService(
	walletRepo *repository.WalletRepository,
	redis *redis.Client) *WalletService {
	return &WalletService{
		WalletRepo: walletRepo,
		Redis:      redis,
	}
}

func (s *WalletService) TopUp(token string, wallet *models.Wallet) error {
	result, err := s.Redis.Get(context.Background(), token).Result()
	if err != nil {
		return err
	}

	playerId, err := utils.StringToUint(result)
	if err != nil {
		return err
	}

	wallet.PlayerID = playerId
	return s.WalletRepo.UpdateOrCreate(wallet)
}
