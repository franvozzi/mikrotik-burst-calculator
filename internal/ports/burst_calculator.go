package ports

import "github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"

type BurstCalculator interface {
	Calculate(req models.BurstRequest) (*models.BurstResponse, error)
}
