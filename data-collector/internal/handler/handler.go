package handler

import(

	"github.com/gofiber/fiber/v2"

	"github.com/uchimann/air_pollution_project/data-collector/internal/handler"
)

func SetRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
   
	apiRouter.Post("/pollution", handler.AddPollutionDataHandler)
}

