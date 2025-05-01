package handler

func AddPollutionDataHandler(ctx *fiber.Ctx) error {
    var model model.PollutantDataInput
    if err := ctx.BodyParser(&model); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid payload"})
    }
    if err := service.AddPollutionData(&model); err != nil {
        // validation veya DB hatasına göre uygun status dönülecek
        return ctx.Status(determineStatus(err)).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"data": model})
}