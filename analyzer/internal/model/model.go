package model

import (
    "time"
)

type PollutantData struct {
    Timestamp  time.Time `json:"timestamp"`
    Latitude   float64   `json:"latitude"`
    Longitude  float64   `json:"longitude"`
    Pollutant  string    `json:"pollutant"`
    Value      float64   `json:"value"`
}


type PollutionAnalysis struct {
    ID             uint      `gorm:"primaryKey" json:"id"`
    PollutionDataID uint      `gorm:"index" json:"pollution_data_id"`
    AnalysisTime    time.Time `gorm:"autoCreateTime" json:"timestamp"`
    Latitude        float64   `json:"latitude"`
    Longitude       float64   `json:"longitude"`
    Pollutant       string    `json:"pollutant"`
    Value           float64   `json:"value"`
    ThresholdValue  float64   `json:"threshold_value"`
    IsAnomalous     bool      `json:"is_anomalous"`
    AnomalyLevel    string    `json:"anomaly_level"` 
    HealthRiskLevel string    `json:"health_risk_level"`
}

const (
    AnomalyLevelLow    = "Low"
    AnomalyLevelMedium = "Medium"
    AnomalyLevelHigh   = "High"
)

const (
    HealthRiskSafe      = "Safe"
    HealthRiskModerate  = "Moderate"
    HealthRiskUnhealthy = "Unhealthy"
    HealthRiskHazardous = "Hazardous"
)

var PollutantThresholds = map[string]struct {
    Moderate  float64
    Unhealthy float64
    Hazardous float64
}{
    "PM2.5": {12.0, 35.4, 55.4},
    "PM10":  {54.0, 154.0, 254.0},
    "NO2":   {100.0, 360.0, 649.0},
    "SO2":   {35.0, 185.0, 304.0},
    "O3":    {54.0, 70.0, 85.0},
}