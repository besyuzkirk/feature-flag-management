package handlers

import (
	"net/http"
	"strconv"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	"github.com/gin-gonic/gin"
)

type RolloutStrategyHandler struct {
	service services.RolloutStrategyService
}

func NewRolloutStrategyHandler(service services.RolloutStrategyService) *RolloutStrategyHandler {
	return &RolloutStrategyHandler{service}
}

func (h *RolloutStrategyHandler) CreateRolloutStrategy(c *gin.Context) {
	var input struct {
		FeatureFlagID uint   `json:"feature_flag_id" binding:"required"`
		Percentage    int    `json:"percentage" binding:"required"`
		Description   string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	strategy, err := h.service.CreateRolloutStrategy(input.FeatureFlagID, input.Percentage, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, strategy)
}

func (h *RolloutStrategyHandler) UpdateRolloutStrategy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Percentage  int    `json:"percentage" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	strategy, err := h.service.UpdateRolloutStrategy(uint(id), input.Percentage, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, strategy)
}

func (h *RolloutStrategyHandler) DeleteRolloutStrategy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteRolloutStrategy(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *RolloutStrategyHandler) GetRolloutStrategy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	strategy, err := h.service.GetRolloutStrategy(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, strategy)
}

func (h *RolloutStrategyHandler) ListRolloutStrategiesByFlag(c *gin.Context) {
	flagID, err := strconv.Atoi(c.Param("flag_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Feature Flag ID"})
		return
	}

	strategies, err := h.service.ListRolloutStrategiesByFlag(uint(flagID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, strategies)
}

func (h *RolloutStrategyHandler) TrackRolloutProgress(c *gin.Context) {
	flagID, err := strconv.Atoi(c.Param("flag_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Feature Flag ID"})
		return
	}

	percentage, err := h.service.TrackRolloutProgress(uint(flagID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"rollout_percentage": percentage})
}
