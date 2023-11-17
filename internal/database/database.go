package database

import (
	"fmt"
	"go-player-test/internal/api/models"
	"go-player-test/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	fmt.Println(cfg)
	dsn := GetDSN(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDSN(host, port, user, password, dbname string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func Migrate(db *gorm.DB) *gorm.DB {
	// Add your model migrations here
	db.AutoMigrate(
		&models.Player{},
		&models.BankAccount{},
		&models.Wallet{},
	)

	return db
}
