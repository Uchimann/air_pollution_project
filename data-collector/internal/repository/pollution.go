package repository

func CreatePollution(input *model.PollutantDataInput) error {
    return DB.Create(input).Error
}