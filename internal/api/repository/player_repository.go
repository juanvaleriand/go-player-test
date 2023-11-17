package repository

import (
	"go-player-test/internal/api/models"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{DB: db}
}

func (r *PlayerRepository) Create(player *models.Player) error {
	return r.DB.Create(player).Error
}

func (r *PlayerRepository) FindByID(id uint) (*models.Player, error) {
	var player models.Player
	err := r.DB.Preload("BankAccounts").Preload("Wallet").First(&player, id).Error
	return &player, err
}

func (r *PlayerRepository) FindByUsername(username string) (*models.Player, error) {
	var player models.Player
	err := r.DB.Preload("BankAccounts").Preload("Wallet").Where("username = ?", username).First(&player).Error
	return &player, err
}

func (r *PlayerRepository) FindAll(filter string) ([]models.Player, error) {
	var players []models.Player
	query := r.DB.Preload("BankAccounts").Preload("Wallet")

	if filter != "" {
		query = query.Joins("LEFT JOIN bank_accounts ON players.id = bank_accounts.player_id").
			Where("players.username LIKE ? OR bank_accounts.account_name LIKE ? OR bank_accounts.account_number LIKE ? OR bank_accounts.bank_name LIKE ?", "%"+filter+"%", "%"+filter+"%", "%"+filter+"%", "%"+filter+"%")
	}

	err := query.Find(&players).Error
	return players, err
}
