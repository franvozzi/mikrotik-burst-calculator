package main

import (
	"github.com/franvozzi/mikrotik-burst-calculator/internal/adapters/http"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/services/calculator"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/services/simulator"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/usecases"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inyección de dependencias para el calculador
	burstService := calculator.NewBurstService()
	calculateUC := usecases.NewCalculateBurstUseCase(burstService)

	// Inyección de dependencias para el simulador
	simulatorService := simulator.NewBurstSimulator()
	simulateUC := usecases.NewSimulateBurstUseCase(simulatorService)

	// Crear handlers
	handler := http.NewBurstHandler(calculateUC, simulateUC)

	r := gin.Default()

	// Rutas
	r.POST("/api/calculate", handler.Calculate)
	r.POST("/api/simulate", handler.Simulate)

	r.Run(":8080")
}
