package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/uchimann/air_pollution_project/data-collector/internal/handler"
	"github.com/uchimann/air_pollution_project/data-collector/internal/repository"
)


func main(){

	app := fiber.New()
	repository.StartConnection()
	handler.SetRoutes(app)
	err := app.Listen(":8080")
	if err != nil {
	 log.Fatalf("Error listen server %s", err)
	}
}
