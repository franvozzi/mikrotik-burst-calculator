package models

// Parametro de input para calcular burst
type BurstRequest struct {
	MaxLimit   int64 `json:"max_limit"`
	BurstLimit int64 `json:"burst_limit"`
	BurstTime  int   `json:"burst_time"`
	Priority   int   `json:"priority"`
}

// Parametro del output (resultado)
type BurstResponse struct {
	MaxLimit       int64   `json:"max_limit"`
	BurstLimit     int64   `json:"burst_limit"`
	BurstThreshold int64   `json:"burst_threshold"`
	BurstTime      int     `json:"burst_time"`
	AverageRate    float64 `json:"average_rate"`
}

// SimulationRequest para simulación de tráfico con burst
type SimulationRequest struct {
	MaxLimit       int64 `json:"max_limit"`       // Límite máximo sin burst (bps)
	BurstLimit     int64 `json:"burst_limit"`     // Límite durante burst (bps)
	BurstThreshold int64 `json:"burst_threshold"` // Umbral para activar burst (bps)
	BurstTime      int   `json:"burst_time"`      // Tiempo de ventana burst (segundos)
	Duration       int   `json:"duration"`        // Duración de la simulación (segundos)
}

// SimulationResponse contiene los resultados de la simulación
type SimulationResponse struct {
	TimePoints   []int     `json:"time_points"`   // Puntos de tiempo en segundos
	DataRates    []int64   `json:"data_rates"`    // Velocidad actual en cada punto (bps)
	AverageRates []float64 `json:"average_rates"` // Promedio móvil en cada punto (bps)
	BurstActive  []bool    `json:"burst_active"`  // Estado del burst en cada punto
	TotalData    int64     `json:"total_data"`    // Total de datos transferidos (bytes)
	AverageSpeed float64   `json:"average_speed"` // Velocidad promedio general (bps)
}

// DataPoint representa un punto en el tiempo de la simulación
type DataPoint struct {
	Time        int     `json:"time"`
	CurrentRate int64   `json:"current_rate"`
	AverageRate float64 `json:"average_rate"`
	BurstActive bool    `json:"burst_active"`
}
