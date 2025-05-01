package repository

import(

	"github.com/uchimann/air_pollution_project/data-collector/internal/model"
)

func CreatePollution(input *model.PollutantDataInput) error {
    return DB.Create(input).Error
}