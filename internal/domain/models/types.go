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

// Simulación Request para simular tráfico
type SimulationRequest struct {
	MaxLimit   int64 `json:"max_limit"`
	BurstLimit int64 `json:"burst_limit"`
	BurstTime  int   `json:"burst_time"`
	Duration   int   `json:"duration"`
}
