package utils

import (
	"io"
	"log"
	"os"
)

const (
	LevelInfo  = "INFO"
	LevelError = "ERROR"
	LevelDebug = "DEBUG"
)

// CreateLogger создает логгер с уровнями
func CreateLogger(logFilePath string) *log.Logger {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Ошибка при создании лог-файла: %v", err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	logger := log.New(multiWriter, "POMPON-BOT: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

func LogInfo(logger *log.Logger, message string) {
	logger.Printf("[%s] %s", LevelInfo, message)
}

func LogError(logger *log.Logger, message string) {
	logger.Printf("[%s] %s", LevelError, message)
}

func LogDebug(logger *log.Logger, message string) {
	logger.Printf("[%s] %s", LevelDebug, message)
}
