package handler

import(

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
   
	apiRouter.Post("/pollution", AddPollutionDataHandler)
}

