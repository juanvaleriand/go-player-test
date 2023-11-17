package repository

import (
	"go-player-test/internal/api/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{DB: db}
}

func (r *WalletRepository) UpdateOrCreate(wallet *models.Wallet) error {
	var existingWallet models.Wallet
	err := r.DB.Where("player_id = ?", wallet.PlayerID).First(&existingWallet).Error

	if err == nil {
		return r.DB.Model(&existingWallet).Update("balance", gorm.Expr("balance + ?", wallet.Balance)).Error
	} else if err == gorm.ErrRecordNotFound {
		return r.DB.Create(&models.Wallet{PlayerID: wallet.PlayerID, Balance: wallet.Balance}).Error
	}

	return err
}
