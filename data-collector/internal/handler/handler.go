package handler

import(

	"github.com/gofiber/fiber/v2"
	"github.com/uchimann/air_pollution_project/data-collector/internal/repository"
	"github.com/uchimann/air_pollution_project/data-collector/internal/model"
)

func SetRoutes(app *fiber.App) {
	// Add "/api" prefix for all routes
	apiRouter := app.Group("/api")
   
	// Products routes
	apiRouter.Post("/pollution", AddPollutionData)
}

   
func AddPollutionData(ctx *fiber.Ctx) error {
	var pollution model.PollutantDataInput
	var err = ctx.BodyParser(&pollution)

	if err != nil {
	 return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Invalid request",
	 })
	}

	err = repository.DB.Create(&pollution).Error
	if err != nil {
	 return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Create operation failed",
	 })
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
	 "data":  pollution,
	 "error": nil,
	})
}
   
   