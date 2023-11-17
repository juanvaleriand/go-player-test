// main.go
package main

import (
	"fmt"
	"go-player-test/internal/api/handlers"
	"go-player-test/internal/api/repository"
	"go-player-test/internal/api/routes"
	"go-player-test/internal/api/services"
	"go-player-test/internal/config"
	"go-player-test/internal/database"
	"go-player-test/internal/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		RedisAddr:  os.Getenv("REDIS_ADDR"),
		RedisPass:  "",
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	rdb := utils.InitRedis(cfg)
	defer rdb.Close()

	// Migrate Database
	database.Migrate(db)

	playerRepo := repository.NewPlayerRepository(db)

	// Auth service
	authService := services.NewAuthService(playerRepo, rdb)
	authHandler := handlers.NewAuthHandler(authService)

	// Player service
	playerService := services.NewPlayerService(playerRepo, rdb)
	playerHandler := handlers.NewPlayerHandler(playerService)

	// Bank account service
	bankAccountRepo := repository.NewBankAccountRepository(db)
	bankAccountService := services.NewBankAccountService(bankAccountRepo, rdb)
	bankAccountHandler := handlers.NewBankAccountHandler(bankAccountService)

	// Wallet service
	walletRepo := repository.NewWalletRepository(db)
	walletService := services.NewWalletService(walletRepo, rdb)
	walletHandler := handlers.NewWalletHandler(walletService)

	router := routes.SetupRouter(authHandler, playerHandler, bankAccountHandler, walletHandler)
	router.Run(":8080")
}
