package model

import (
	"time"
)


type PollutantData struct{
	Timestamp  	time.Time	
	Latitude 	float64		
	Longitude	float64		
	Pollutant 	string		
	Value		float64		
}

type PollutantDataInput struct{
	Timestamp  	time.Time	`gorm:"autoCreateTime" json:"timestamp"`
	Latitude 	float64		`db:"latitude" json:"latitude"`
	Longitude	float64		`db:"longitude" json:"longitude"`
	Pollutant 	string		`db:"pollutant" json:"pollutant"`
	Value		float64		`db:"value" json:"value"`
}
