package handlers

import (
	"go-player-test/internal/api/models"
	"go-player-test/internal/api/requests"
	"go-player-test/internal/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var playerRequest requests.RegisterPlayerRequest
	if err := c.ShouldBindJSON(&playerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the registration service
	err := h.AuthService.Register(&models.Player{
		FullName: playerRequest.FullName,
		Email:    playerRequest.Email,
		Username: playerRequest.Username,
		Password: playerRequest.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register player"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Player registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest requests.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var request requests.LogoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// Handle invalid request
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := h.AuthService.Logout(request.Token)
	if err != nil {
		// Handle logout error
		c.JSON(500, gin.H{"error": "Logout failed"})
		return
	}

	c.JSON(200, gin.H{"message": "Logout successful"})
}
