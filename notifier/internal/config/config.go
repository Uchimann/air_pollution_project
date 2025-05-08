package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func Get(key string, defaultValue string) string {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println(".env file not found")
        return defaultValue
    }

    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}