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

    repository.StartConnection()

    rabbitClient, err := rabbitmq.NewClient()
    if err != nil {
        log.Fatalf("Error establishing abbitMQ connection: %v", err)
    }
    defer rabbitClient.Close()

    analyzerService := service.NewAnalyzerService(repository.DB, rabbitClient)

    consumer, err := rabbitmq.NewConsumer(rabbitClient, analyzerService)
    if err != nil {
        log.Fatalf("RabbitMQ consumer create error: %v", err)
    }
    
    go func() {
        if err := consumer.StartConsuming(); err != nil {
            log.Fatalf("Data listen proccees cannot start: %v", err)
        }
    }()
    
    log.Println("Analyzer service started. Data listening...")

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Analyzer service closing...")
}