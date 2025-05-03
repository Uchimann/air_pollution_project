package service

import (
	"fmt"
	"log"
	"reflect"

	"github.com/uchimann/air_pollution_project/analyzer/internal/analyzer"
	"github.com/uchimann/air_pollution_project/analyzer/internal/rabbitmq"
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
            log.Printf("Error while analyzing data: %s", err)
            continue
        }
        if bool {
            log.Printf("Anomaly detected in data: %s", string(data))
        }
    }
    
    log.Println("ended of message channel")
}