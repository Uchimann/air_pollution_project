package handler

import(
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/uchimann/air_pollution_project/data-collector/internal/service"
	"github.com/uchimann/air_pollution_project/data-collector/internal/model"
	"github.com/uchimann/air_pollution_project/data-collector/internal/repository"
)


func AddPollutionDataHandler(ctx *fiber.Ctx) error {
    var model model.PollutantDataInput
    if err := ctx.BodyParser(&model); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid payload"})
    }
    if err := service.AddPollutionData(&model); err != nil {
        return ctx.Status(determineStatus(err)).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": model})
}

// determineStatus verilen hataya göre HTTP status kodunu döner.
func determineStatus(err error) int {

    if errors.Is(err, service.ErrUnsupportedPollutant) {
        return fiber.StatusBadRequest
    }

	if errors.Is(err, service.ErrUnsupportedPollutant){
		return fiber.StatusForbidden
	}

    return fiber.StatusInternalServerError
}