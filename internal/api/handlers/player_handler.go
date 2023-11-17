package handlers

import (
	"go-player-test/internal/api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	PlayerService *services.PlayerService
}

func NewPlayerHandler(playerService *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{PlayerService: playerService}
}

func (h *PlayerHandler) GetPlayers(c *gin.Context) {
	filter := c.Query("filter")
	players, err := h.PlayerService.GetPlayers(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Failed to get players"})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (h *PlayerHandler) GetPlayerDetail(c *gin.Context) {
	playerIDStr := c.Param("id")
	playerID, err := strconv.ParseUint(playerIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	player, err := h.PlayerService.GetPlayerDetail(uint(playerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found!"})
		return
	}

	c.JSON(http.StatusOK, player)
}
