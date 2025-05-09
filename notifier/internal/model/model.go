package model

import (
    "time"
)

type PollutionAnalysis struct {
    ID             uint      `json:"id"`
    Latitude       float64   `json:"latitude"`
    Longitude      float64   `json:"longitude"`
    PollutionDataID uint      `json:"pollution_data_id"`
    Timestamp       time.Time `json:"timestamp"`
    Pollutant       string    `json:"pollutant"`
    Value           float64   `json:"value"`
    ThresholdValue  float64   `json:"threshold_value"`
    IsAnomalous     bool      `json:"is_anomalous"`
    AnomalyLevel    string    `json:"anomaly_level"`
    HealthRiskLevel string    `json:"health_risk_level"`
}
