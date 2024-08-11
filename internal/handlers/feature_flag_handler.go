package handlers

import (
	"net/http"
	"strconv"

	"github.com/besyuzkirk/feature-flag-management/internal/domain/services"
	"github.com/gin-gonic/gin"
)

type FeatureFlagHandler struct {
	service services.FeatureFlagService
}

func NewFeatureFlagHandler(service services.FeatureFlagService) *FeatureFlagHandler {
	return &FeatureFlagHandler{service}
}

func (h *FeatureFlagHandler) CreateFeatureFlag(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag, err := h.service.CreateFeatureFlag(input.Name, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flag)
}

func (h *FeatureFlagHandler) UpdateFeatureFlag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		IsActive    bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flag, err := h.service.UpdateFeatureFlag(uint(id), input.Name, input.Description, input.IsActive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flag)
}

func (h *FeatureFlagHandler) DeleteFeatureFlag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteFeatureFlag(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *FeatureFlagHandler) GetFeatureFlag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	flag, err := h.service.GetFeatureFlag(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flag)
}

func (h *FeatureFlagHandler) ListFeatureFlags(c *gin.Context) {
	flags, err := h.service.ListFeatureFlags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flags)
}

func (h *FeatureFlagHandler) RollbackFeatureFlag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	flag, err := h.service.RollbackFeatureFlag(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flag)
}
