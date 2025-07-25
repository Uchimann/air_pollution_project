package analyzer

import (
	"log"
	"fmt"
	"encoding/json"
	"github.com/uchimann/air_pollution_project/analyzer/internal/model"
)

func AnomalyDetection(data []byte) (bool, error) {
	log.Printf("Anomaly detection started for data: %s\n", string(data))
	
	var pollutantData model.PollutantData
	if err := json.Unmarshal(data,&pollutantData); err != nil {
		log.Printf("Error while unmarshalling data: %s", err)
		return false, err
	}

	thresholds := model.PollutantThresholds[pollutantData.Pollutant]

	isAnomaly := false

	if pollutantData.Value > thresholds.Hazardous {
		isAnomaly = true
	} else if pollutantData.Value > thresholds.Unhealthy {
		isAnomaly = true
	} else if pollutantData.Value > thresholds.Moderate {
		isAnomaly = true
	} else {

	}		

	return isAnomaly, nil
}

func AnalyzePollutionData(data *model.PollutantData) (*model.PollutionAnalysis, error) {
    
    
    thresholds, exists := model.PollutantThresholds[data.Pollutant]
    if !exists {
        return nil, fmt.Errorf("desteklenmeyen kirletici türü: %s", data.Pollutant) 
    }
    
    result := &model.PollutionAnalysis{
        Pollutant:      data.Pollutant,
        Value:          data.Value,
        ThresholdValue: thresholds.Moderate,
        AnomalyLevel:   model.AnomalyLevelLow,
        HealthRiskLevel: model.HealthRiskSafe,
		AnalysisTime: data.Timestamp,
		Latitude: data.Latitude,
		Longitude: data.Longitude,
    }
    
    if data.Value > thresholds.Hazardous {
        result.IsAnomalous = true
        result.AnomalyLevel = model.AnomalyLevelHigh
        result.HealthRiskLevel = model.HealthRiskHazardous
        result.ThresholdValue = thresholds.Hazardous
    } else if data.Value > thresholds.Unhealthy {
        result.IsAnomalous = true
        result.AnomalyLevel = model.AnomalyLevelMedium
        result.HealthRiskLevel = model.HealthRiskUnhealthy
        result.ThresholdValue = thresholds.Unhealthy
    } else if data.Value > thresholds.Moderate {
        result.IsAnomalous = true
        result.AnomalyLevel = model.AnomalyLevelLow
        result.HealthRiskLevel = model.HealthRiskModerate
        result.ThresholdValue = thresholds.Moderate
    }
    
    return result, nil
}