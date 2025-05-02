package service

import (
    "github.com/uchimann/air_pollution_project/analyzer/internal/model"
    "github.com/uchimann/air_pollution_project/analyzer/internal/rabbitmq"
    "github.com/uchimann/air_pollution_project/analyzer/internal/repository"
    "gorm.io/gorm"
)

type AnalyzerService struct {
    db     *gorm.DB
    rabbit *rabbitmq.Client
}

func NewAnalyzerService(db *gorm.DB, rabbit *rabbitmq.Client) *AnalyzerService {
    return &AnalyzerService{
        db:     db,
        rabbit: rabbit,
    }
}

func (s *AnalyzerService) NotifyAnomaly(analysis *model.PollutionAnalysis) error {
    return s.rabbit.PublishAnomaly(analysis)
}