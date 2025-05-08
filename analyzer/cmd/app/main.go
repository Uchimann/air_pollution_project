package main

import (
	"log"
    "os"
    "os/signal"
    "syscall"

	"github.com/uchimann/air_pollution_project/analyzer/internal/rabbitmq"
	"github.com/uchimann/air_pollution_project/analyzer/internal/repository"
	"github.com/uchimann/air_pollution_project/analyzer/internal/service"
)

func main() {
    log.Println("Analyzer service starting...")
    
    repository.StartConnection()
    
    rabbitClient, err := rabbitmq.NewClient()
    if err != nil {
        log.Fatalf("RabbitMQ connection error: %s", err)
    }
    
    defer func() {
        if err := rabbitClient.Close(); err != nil {
            log.Printf("Error while closing RabbitMQ connection : %v", err)
        }
    }()
    
    service := service.NewAnalyzerService(repository.DB, rabbitClient)
    if err := service.StartAnalysis(); err != nil {
        log.Fatalf("Error while starting analysis: %s", err)
    }


    // Sinyal yakalama için kanal
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    
    log.Println("Press Ctrl+C to exit")
    <-sig
    log.Println("Analyzer service closing...")
    
}