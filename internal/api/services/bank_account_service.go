package services

import (
	"context"
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/repository"
	"go-player-test/internal/utils"

	"github.com/go-redis/redis/v8"
)

type BankAccountService struct {
	BankAccountRepo *repository.BankAccountRepository
	Redis           *redis.Client
}

func NewBankAccountService(
	bankAccountRepo *repository.BankAccountRepository,
	redis *redis.Client) *BankAccountService {
	return &BankAccountService{
		BankAccountRepo: bankAccountRepo,
		Redis:           redis,
	}
}

func (s *BankAccountService) RegisterBankAccount(token string, bankAccount *models.BankAccount) error {
	result, err := s.Redis.Get(context.Background(), token).Result()
	if err != nil {
		return err
	}

	playerId, err := utils.StringToUint(result)
	if err != nil {
		return err
	}

	bankAccount.PlayerID = playerId
	return s.BankAccountRepo.UpdateOrCreate(bankAccount)
}
