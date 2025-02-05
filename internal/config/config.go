package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Настройки приложения

type Config struct {
	TelegramAPIKey string
	DBUser         string
	DBPassword     string
	DBName         string
	DBHost         string
	DBPort         string
}

func LoadConfig() *Config {
	// Загрузка из .env
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		TelegramAPIKey: os.Getenv("TELEGRAM_API_KEY"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
	}
}
