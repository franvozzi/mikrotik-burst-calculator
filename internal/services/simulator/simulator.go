package simulator

import (
	"errors"
	"math"

	"github.com/franvozzi/mikrotik-burst-calculator/internal/domain/models"
	"github.com/franvozzi/mikrotik-burst-calculator/internal/ports"
)

type burstSimulator struct{}

// NewBurstSimulator crea una nueva instancia del simulador
func NewBurstSimulator() ports.BurstSimulator {
	return &burstSimulator{}
}

// Simulate ejecuta la simulación del comportamiento del burst según el algoritmo de Mikrotik
// Cada 1/16 del burst-time, el router calcula el promedio de datos sobre el burst-time completo
func (s *burstSimulator) Simulate(req models.SimulationRequest) (*models.SimulationResponse, error) {
	// Validaciones
	if err := s.validate(req); err != nil {
		return nil, err
	}

	// Inicializar respuesta
	response := &models.SimulationResponse{
		TimePoints:   make([]int, 0, req.Duration),
		DataRates:    make([]int64, 0, req.Duration),
		AverageRates: make([]float64, 0, req.Duration),
		BurstActive:  make([]bool, 0, req.Duration),
	}

	// Buffer circular para calcular el promedio móvil sobre burst-time
	burstTimeWindow := req.BurstTime
	dataHistory := make([]int64, burstTimeWindow)

	var totalData int64 = 0

	// Simular cada segundo
	for t := 0; t < req.Duration; t++ {
		// Calcular el promedio de los últimos burst-time segundos
		averageRate := s.calculateAverageRate(dataHistory, burstTimeWindow)

		// Determinar si el burst está activo
		// Burst se activa si el promedio está por debajo del threshold
		burstActive := averageRate < float64(req.BurstThreshold)

		// Establecer la velocidad actual según el estado del burst
		var currentRate int64
		if burstActive {
			currentRate = req.BurstLimit
		} else {
			currentRate = req.MaxLimit
		}

		// Actualizar el historial circular (índice módulo para circular)
		dataHistory[t%burstTimeWindow] = currentRate

		// Acumular datos transferidos (convertir bps a bytes)
		totalData += currentRate / 8

		// Guardar resultados
		response.TimePoints = append(response.TimePoints, t)
		response.DataRates = append(response.DataRates, currentRate)
		response.AverageRates = append(response.AverageRates, averageRate)
		response.BurstActive = append(response.BurstActive, burstActive)
	}

	// Calcular estadísticas finales
	response.TotalData = totalData
	response.AverageSpeed = s.calculateOverallAverage(response.DataRates)

	return response, nil
}

// validate verifica que los parámetros de entrada sean válidos
func (s *burstSimulator) validate(req models.SimulationRequest) error {
	if req.MaxLimit <= 0 {
		return errors.New("max_limit debe ser mayor a 0")
	}
	if req.BurstLimit <= 0 {
		return errors.New("burst_limit debe ser mayor a 0")
	}
	if req.BurstLimit <= req.MaxLimit {
		return errors.New("burst_limit debe ser mayor que max_limit")
	}
	if req.BurstThreshold <= 0 {
		return errors.New("burst_threshold debe ser mayor a 0")
	}
	if req.BurstThreshold >= req.BurstLimit {
		return errors.New("burst_threshold debe ser menor que burst_limit")
	}
	if req.BurstTime <= 0 {
		return errors.New("burst_time debe ser mayor a 0")
	}
	if req.Duration <= 0 {
		return errors.New("duration debe ser mayor a 0")
	}
	if req.Duration > 300 {
		return errors.New("duration no puede ser mayor a 300 segundos")
	}

	return nil
}

// calculateAverageRate calcula el promedio de velocidad sobre la ventana de tiempo
func (s *burstSimulator) calculateAverageRate(dataHistory []int64, burstTime int) float64 {
	var sum int64 = 0
	for _, rate := range dataHistory {
		sum += rate
	}
	return float64(sum) / float64(burstTime)
}

// calculateOverallAverage calcula el promedio general de todas las velocidades
func (s *burstSimulator) calculateOverallAverage(rates []int64) float64 {
	if len(rates) == 0 {
		return 0
	}

	var sum int64 = 0
	for _, rate := range rates {
		sum += rate
	}

	return math.Round(float64(sum) / float64(len(rates)))
}
