package service

import (
	"fmt"
	"time"
	"github.com/uchimann/air_pollution_project/data-collector/internal/model"
	"github.com/uchimann/air_pollution_project/data-collector/internal/repository"
)


func AddPollutionData(in *model.PollutantDataInput) error {


	if in.Timestamp.IsZero() {
        in.Timestamp = time.Now()
    }

	if err := validatePollution(in.Pollutant); err != nil{
		return err
	}

    if err := repository.CreatePollution(in); err != nil {
        return fmt.Errorf("db error: %w", err)
    }
    return nil
}

