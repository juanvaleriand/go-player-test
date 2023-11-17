package routes

import (
	"go-player-test/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router
func SetupRouter(
	authHandler *handlers.AuthHandler,
	playerHandler *handlers.PlayerHandler,
	bankAccountHandler *handlers.BankAccountHandler,
	walletHandler *handlers.WalletHandler,
) *gin.Engine {
	router := gin.Default()
	authMiddleware := handlers.AuthMiddleware()

	// Group routes
	api := router.Group("api")
	{

		// Auth routes
		authRoutes := api.Group("auth")
		{
			authRoutes.POST("/register", authHandler.Register)
			authRoutes.POST("/login", authHandler.Login)
			authRoutes.POST("/logout", authHandler.Logout)
		}

		// Player routes
		playerRoutes := api.Group("players")
		{
			playerRoutes.GET("/", playerHandler.GetPlayers)
			playerRoutes.GET("/:id", playerHandler.GetPlayerDetail)
		}

		// Bank account routes
		bankAccountRoutes := api.Group("bank-account")
		{
			bankAccountRoutes.Use(authMiddleware)
			bankAccountRoutes.POST("/", bankAccountHandler.RegisterBankAccount)
		}

		// Wallet routes
		walletRoutes := api.Group("wallet")
		{
			walletRoutes.Use(authMiddleware)
			walletRoutes.POST("/topup", walletHandler.TopUp)
		}
	}

	return router
}
