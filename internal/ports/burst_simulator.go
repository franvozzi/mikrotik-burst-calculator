package ports

import "github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"

// BurstSimulator define el contrato para simular el comportamiento del burst
type BurstSimulator interface {
	Simulate(req models.SimulationRequest) (*models.SimulationResponse, error)
}
