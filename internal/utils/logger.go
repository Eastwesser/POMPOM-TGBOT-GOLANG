package utils

import (
	"io"
	"log"
	"os"
)

// CreateLogger создает логгер с записью в файл и консоль
func CreateLogger(logFilePath string) *log.Logger {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Ошибка при создании лог-файла: %v", err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	logger := log.New(multiWriter, "POMPON-BOT: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
