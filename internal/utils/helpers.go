package utils

import "log"

// Утилиты и вспомогательные функции (например, форматирование текста, snakeCase, обработка ошибок).

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
