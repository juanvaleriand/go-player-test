package handlers

import (
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/requests"
	"go-player-test/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WallerService *services.WalletService
}

func NewWalletHandler(walletService *services.WalletService) *WalletHandler {
	return &WalletHandler{WallerService: walletService}
}

func (h *WalletHandler) TopUp(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var walletRequest requests.WalletRequest

	if err := c.ShouldBindJSON(&walletRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the wallet service
	err := h.WallerService.TopUp(token, &models.Wallet{
		Balance: walletRequest.Balance,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Top up balance successfully"})
}
