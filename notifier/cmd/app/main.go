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
        log.Fatalf("RabbitMQ bağlantı hatası: %v", err)
    }
    defer rabbitClient.Close()

    err = rabbitClient.ConsumeNotifications(eventServer)
    if err != nil {
        log.Fatalf("RabbitMQ consumer başlatılamadı: %v", err)
    }

    http.HandleFunc("/events", eventServer.ServeHTTP)
    log.Println("Notifier SSE servisi 8080 portunda başladı")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("HTTP sunucu hatası: %v", err)
    }
}