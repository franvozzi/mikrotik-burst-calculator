package calculator

import (
	"errors"

	"github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/ports"
)

type burstService struct{}

func NewBurstService() ports.BurstCalculator {
	return &burstService{}
}

func (s *burstService) Calculate(req models.BurstRequest) (*models.BurstResponse, error) {
	if req.MaxLimit <= 0 || req.BurstLimit <= 0 {
		return nil, errors.New("invalid limits")
	}

	threshold := req.BurstLimit - req.MaxLimit
	avgRate := float64(req.MaxLimit) / float64(req.BurstTime)

	return &models.BurstResponse{
		MaxLimit:       req.MaxLimit,
		BurstLimit:     req.BurstLimit,
		BurstThreshold: threshold,
		BurstTime:      req.BurstTime,
		AverageRate:    avgRate,
	}, nil
}
