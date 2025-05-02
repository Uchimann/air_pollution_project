package service

import (
	"fmt"

	"github.com/uchimann/air_pollution_project/data-collector/internal/model"
	"github.com/uchimann/air_pollution_project/data-collector/internal/repository"
)

/*func AddPollutionData(ctx *fiber.Ctx) error {
	var pollution model.PollutantDataInput
	var err = ctx.BodyParser(&pollution)

	if err != nil {
	 return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	  "data":  nil,
	  "error": "Invalid request",
	 })
	}

	//fonksiyon ekle burada aşağıda tanımla analizlerini yapacak fonksiyon olsun


	//kayıt işlemini repositoryde yapacaksın parametre olarak içeriye aldığını ve modelini göncereceksin

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
}/***/

//fonksiyon



func AddPollutionData(in *model.PollutantDataInput) error {

	/*if err := validateTimestamp(in.Timestamp); err != nil {
        return err  
    }*/

	if err := validatePollution(in.Pollutant); err != nil{
		return err
	}

    // 2) Persist
    if err := repository.CreatePollution(in); err != nil {
        return fmt.Errorf("db error: %w", err)
    }
    return nil
}

