package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/uchimann/air_pollution_project/data-collector/internal/handler"
	"github.com/uchimann/air_pollution_project/data-collector/internal/repository"
	"github.com/uchimann/air_pollution_project/data-collector/internal/rabbitmq"
)


func main(){

	app := fiber.New()
	repository.StartConnection()
	rabbitClient, err :=rabbitmq.NewClient()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer func() {
        if err := rabbitClient.Close(); err != nil {
            log.Printf("Error while closing RabbitMQ connection: %v", err)
        }
    }()
	
	handler.SetupDependencies(rabbitClient)
	handler.SetRoutes(app)
	
	errr := app.Listen(":8080")
	if errr != nil {
	 log.Fatalf("Error listen server %s", err)
	}
}
