package handler

import(

	"github.com/gofiber/fiber/v2"

	"github.com/uchimann/air_pollution_project/data-collector/internal/service"
)

func SetRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
   
	apiRouter.Post("/pollution", service.AddPollutionData)
}