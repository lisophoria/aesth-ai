package handlers

import (
	"aesth-api/dto"
	"aesth-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelevancyHandler struct {
	relevancyService *services.RelevancyService
}

func NewRelevancyHandler(relevancyService *services.RelevancyService) *RelevancyHandler {
	return &RelevancyHandler{ relevancyService }
}

func (h *RelevancyHandler) GetByUserID(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	relevancy, err := h.relevancyService.GetRelevancy(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, relevancy)
}

func (h *RelevancyHandler) AdjustRelevancy(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	var input dto.RelevancyDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	relevancy, err := h.relevancyService.AdjustRelevancy(uint(userID), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update"})
		return
	}

	c.JSON(http.StatusOK, relevancy)
}

func (h *RelevancyHandler) GetRelevancyPair(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	var diff dto.RelevancyDTO
	if err := c.ShouldBindJSON(&diff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid diff format"})
		return
	}

	relevancyPair, err := h.relevancyService.GetRelevancyPair(uint(userID), diff)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get diff"})
	}

	c.JSON(http.StatusOK, relevancyPair)
}