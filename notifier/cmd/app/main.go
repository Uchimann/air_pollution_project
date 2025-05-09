package main

import (
    "log"
    "net/http"
    "github.com/uchimann/air_pollution_project/notifier/internal/rabbitmq"
    "github.com/uchimann/air_pollution_project/notifier/internal/sse"
)

func main() {
    eventServer := sse.NewEventServer()
    rabbitClient, err := rabbitmq.NewClient()
    if err != nil {
        log.Fatalf("RabbitMQ connection error: %v", err)
    }
    defer rabbitClient.Close()

    err = rabbitClient.ConsumeNotifications(eventServer)
    if err != nil {
        log.Fatalf("RabbitMQ consumer couldnt start: %v", err)
    }

    http.HandleFunc("/events", eventServer.ServeHTTP)
    log.Println("Notifier SSE service started on port 8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatalf("HTTP server error: %v", err)
    }
}