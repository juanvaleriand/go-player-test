package handlers

import (
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/requests"
	"go-player-test/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankAccountHandler struct {
	BankAccountService *services.BankAccountService
}

func NewBankAccountHandler(bankAccountService *services.BankAccountService) *BankAccountHandler {
	return &BankAccountHandler{BankAccountService: bankAccountService}
}

func (h *BankAccountHandler) RegisterBankAccount(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var bankAccountRequest requests.BankAccountRequest

	if err := c.ShouldBindJSON(&bankAccountRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Call the BankAccount service
	err := h.BankAccountService.RegisterBankAccount(token, &models.BankAccount{
		AccountName:   bankAccountRequest.AccountName,
		AccountNumber: bankAccountRequest.AccountNumber,
		BankName:      bankAccountRequest.BankName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank account registered successfully"})
}
