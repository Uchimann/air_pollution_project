package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func Get(key string, value string) string {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println(".env file not found")
        return value
    }

    return os.Getenv(key)
}