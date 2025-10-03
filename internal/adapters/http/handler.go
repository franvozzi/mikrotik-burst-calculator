package http

import (
	"net/http"

	"github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/usecases"
	"github.com/gin-gonic/gin"
)

type BurstHandler struct {
	calculateUseCase *usecases.CalculateBurstUseCase
	simulateUseCase  *usecases.SimulateBurstUseCase
}

// NewBurstHandler crea un nuevo handler con ambos usecases
func NewBurstHandler(calcUC *usecases.CalculateBurstUseCase, simUC *usecases.SimulateBurstUseCase) *BurstHandler {
	return &BurstHandler{
		calculateUseCase: calcUC,
		simulateUseCase:  simUC,
	}
}

// Calculate maneja las peticiones de cálculo de burst
func (h *BurstHandler) Calculate(c *gin.Context) {
	var req models.BurstRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.calculateUseCase.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Simulate maneja las peticiones de simulación de burst
func (h *BurstHandler) Simulate(c *gin.Context) {
	var req models.SimulationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.simulateUseCase.Execute(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
