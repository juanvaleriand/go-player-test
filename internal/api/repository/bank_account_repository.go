package repository

import (
	"errors"
	"go-player-test/internal/api/models"

	"gorm.io/gorm"
)

type BankAccountRepository struct {
	DB *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) *BankAccountRepository {
	return &BankAccountRepository{DB: db}
}

func (r *BankAccountRepository) UpdateOrCreate(bankAccount *models.BankAccount) error {
	if err := r.checkDuplicateBankAccount(bankAccount.AccountNumber, bankAccount.BankName, bankAccount.PlayerID); err != nil {
		return err
	}

	var existingBankAccount models.BankAccount
	err := r.DB.Where("player_id = ?", bankAccount.PlayerID).First(&existingBankAccount).Error

	if err == nil {
		return r.DB.Model(&existingBankAccount).Updates(bankAccount).Error
	} else if err == gorm.ErrRecordNotFound {
		return r.DB.Create(bankAccount).Error
	}

	return err
}

func (r *BankAccountRepository) checkDuplicateBankAccount(accountNumber string, bankName string, playerID uint) error {
	var count int64
	query := r.DB.Model(&models.BankAccount{}).Where("account_number = ? AND bank_name = ?", accountNumber, bankName)

	if playerID != 0 {
		query = query.Not("player_id = ?", playerID)
	}

	query.Count(&count)

	if count > 0 {
		return errors.New("cannot have the same account number in the same bank")
	}

	return nil
}
