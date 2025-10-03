package usecases

import (
	"github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/ports"
)

type SimulateBurstUseCase struct {
	simulator ports.BurstSimulator
}

// NewSimulateBurstUseCase crea una nueva instancia del caso de uso de simulación
func NewSimulateBurstUseCase(sim ports.BurstSimulator) *SimulateBurstUseCase {
	return &SimulateBurstUseCase{
		simulator: sim,
	}
}

// Execute ejecuta la simulación de burst
func (uc *SimulateBurstUseCase) Execute(req models.SimulationRequest) (*models.SimulationResponse, error) {
	return uc.simulator.Simulate(req)
}
