package service

import (
	"fmt"
	"log"
	"reflect"

    "encoding/json"

	"github.com/uchimann/air_pollution_project/analyzer/internal/analyzer"
	"github.com/uchimann/air_pollution_project/analyzer/internal/rabbitmq"
    "github.com/uchimann/air_pollution_project/analyzer/internal/model"
    "github.com/uchimann/air_pollution_project/analyzer/internal/repository"
	"gorm.io/gorm"
)

type AnalyzerService struct {
    db         *gorm.DB
    rabbitClient *rabbitmq.Client
    messageChan chan []byte
}

func NewAnalyzerService(db *gorm.DB, rabbitClient *rabbitmq.Client) *AnalyzerService {
    return &AnalyzerService{
        db:         db,
        rabbitClient: rabbitClient,
    }
}

func (s *AnalyzerService) StartAnalysis() error {

    //
    q, err := s.rabbitClient.ConnectQueue()
    if err != nil {
        return fmt.Errorf("Error while connecting to RabbitMQ queue: %w ", err)
    }

    messageChan, err := s.rabbitClient.ConsumeMessages(q)
    if err != nil {
        return fmt.Errorf("Error consuming RabbitMQ messages: %w", err)
    }
    
    s.messageChan = messageChan
    go s.processMessages()

    log.Println("Analyze service started, messages are prosccessing...")
    return nil
}

func (s *AnalyzerService) processMessages() {

    for data := range s.messageChan {
        log.Printf("Message from RabbitMq ( Data Collector ): %s \n variable type is %s", string(data), reflect.TypeOf(data))
        
        // TODO: Burada mesajı analiz edip, gerekirse veritabanına kaydedip, notifier'a bildirim gönderebilirim
        bool, err := analyzer.AnomalyDetection(data)
        if err != nil {
            log.Printf("Error while analyzing anomaly: %s", err)
            continue
        }

        if bool {
            var rawData model.PollutantData
            if err := json.Unmarshal(data, &rawData); err != nil {
                log.Printf("Error while unmarshalling data: %s", err)
                continue
            }

            AnalyzedData, err := analyzer.AnalyzePollutionData(&rawData)
            if err != nil {
                log.Printf("Error while analyzing data: %s", err)
                continue
            }

            if err := repository.SaveAnalysisResult(*AnalyzedData); err != nil {
                log.Printf("Error while saving analysis result: %s", err)
                continue
            }

            if err := s.rabbitClient.PublishAnalysisResult(*AnalyzedData); err != nil {
                log.Printf("Error while publishing analysis result: %s", err)
            } else {
                log.Printf("Notification sent to notifier service for %s (Level: %s, Risk: %s)",
                    AnalyzedData.Pollutant,
                    AnalyzedData.AnomalyLevel,
                    AnalyzedData.HealthRiskLevel)
            }
        }
    }
    log.Println("ended of message channel")
}