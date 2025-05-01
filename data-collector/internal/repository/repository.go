package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/uchiman/air_pollution_project/data-collector/internal/config"
	"github.com/uchiman/air_pollution_project/data-collector/internal/model"
)

var DB *gorm.DB

func StartConnection() {
   var dbUserName = config.Get("DB_USER_NAME","pollution")
   var dbPassword = config.Get("DB_PASSWORD","password")
   var dbHost = config.Get("DB_HOST","localhost")
   var dbName = config.Get("DB_NAME","pollutiondb")

   var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUserName, dbPassword, dbHost, dbName)

   db, err := gorm.Open(mysql.Open(dsn))
   if err != nil {
      log.Fatalf("Database connection error %s", err)
   }

   DB = db

   // Create products table
   DB.AutoMigrate(&model.PollutantDataInput{})
}