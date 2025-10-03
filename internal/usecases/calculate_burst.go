package usecases

import (
	"github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/ports"
)

type CalculateBurstUseCase struct {
	calculator ports.BurstCalculator
}

// NewCalculateBurstUseCase crea una nueva instancia del caso de uso de cálculo
func NewCalculateBurstUseCase(calc ports.BurstCalculator) *CalculateBurstUseCase {
	return &CalculateBurstUseCase{
		calculator: calc,
	}
}

// Execute ejecuta el cálculo de burst
func (uc *CalculateBurstUseCase) Execute(req models.BurstRequest) (*models.BurstResponse, error) {
	return uc.calculator.Calculate(req)
}
