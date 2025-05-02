package service

import (
    "time"

    "github.com/uchimann/air_pollution_project/analyzer/internal/model"
    "github.com/uchimann/air_pollution_project/analyzer/internal/repository"
)

func (s *AnalyzerService) AnalyzePollutionData(data *model.PollutantData) (*model.PollutionAnalysis, error) {
    analysis := &model.PollutionAnalysis{
        AnalysisTime: time.Now(),
        Pollutant:    data.Pollutant,
        Value:        data.Value,
    }

    thresholds, exists := model.PollutantThresholds[data.Pollutant]
    if !exists {
        return nil, ErrUnsupportedPollutant
    }

    analysis.ThresholdValue = thresholds.Moderate

    if data.Value > thresholds.Moderate {
        analysis.IsAnomalous = true
        
        if data.Value > thresholds.Hazardous {
            analysis.AnomalyLevel = model.AnomalyLevelHigh
        } else if data.Value > thresholds.Unhealthy {
            analysis.AnomalyLevel = model.AnomalyLevelMedium
        } else {
            analysis.AnomalyLevel = model.AnomalyLevelLow
        }
        
        if data.Value > thresholds.Hazardous {
            analysis.HealthRiskLevel = model.HealthRiskHazardous
        } else if data.Value > thresholds.Unhealthy {
            analysis.HealthRiskLevel = model.HealthRiskUnhealthy
        } else {
            analysis.HealthRiskLevel = model.HealthRiskModerate
        }
    } else {
        analysis.IsAnomalous = false
        analysis.AnomalyLevel = model.AnomalyLevelLow
        analysis.HealthRiskLevel = model.HealthRiskSafe
    }

    if err := repository.SaveAnalysis(analysis); err != nil {
        return nil, err
    }

    return analysis, nil
}