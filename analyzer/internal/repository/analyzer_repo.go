package repository

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/uchimann/air_pollution_project/analyzer/internal/config"
    "github.com/uchimann/air_pollution_project/analyzer/internal/model"
)

var DB *gorm.DB

func StartConnection() {
    var dbUserName = config.Get("DB_USER_NAME", "pollution")
    var dbPassword = config.Get("DB_PASSWORD", "password")
    var dbHost = config.Get("DB_HOST", "localhost")
    var dbName = config.Get("DB_NAME", "pollutiondb")

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
        dbHost, dbUserName, dbPassword, dbName,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Database connection error %s", err)
    }

    DB = db

    if !DB.Migrator().HasTable(&model.PollutionAnalysis{}) {
        if err := DB.AutoMigrate(&model.PollutionAnalysis{}); err != nil {
            log.Fatalf("AutoMigrate error: %v", err)
        }
    }

    log.Println("Database connection established")
}