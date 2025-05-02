package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/uchimann/air_pollution_project/data-collector/internal/config"
	"github.com/uchimann/air_pollution_project/data-collector/internal/model"
)

var DB *gorm.DB

func StartConnection() {
   var dbUserName = config.Get("DB_USER_NAME","pollution")
   var dbPassword = config.Get("DB_PASSWORD","password")
   var dbHost = config.Get("DB_HOST","localhost")
   var dbName = config.Get("DB_NAME","pollutiondb")

   dsn := fmt.Sprintf(
      "host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
      dbHost, dbUserName, dbPassword, dbName,
  )

   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
   if err != nil {
      log.Fatalf("Database connection error %s", err)
   }

   DB = db

   if err := DB.Exec("CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;").Error; err != nil {
      log.Fatalf("TimescaleDB extension error: %v", err)
  }

   /*DB.AutoMigrate(&model.PollutantDataInput{})

   if err := DB.Exec("SELECT create_hypertable('pollutant_data_inputs', 'timestamp', if_not_exists => true);",
         ).Error; err != nil {
            log.Fatalf("Hypertable create error: %v", err)
  }*/

  if !DB.Migrator().HasTable(&model.PollutantDataInput{}) {
      if err := DB.AutoMigrate(&model.PollutantDataInput{}); err != nil {
         log.Fatalf("AutoMigrate error: %v", err)
      }
      if err := DB.Exec(
         "SELECT create_hypertable('pollutant_data_inputs', 'timestamp', if_not_exists => true);",
      ).Error; err != nil {
         log.Fatalf("Hypertable create error: %v", err)
      }
   }

}