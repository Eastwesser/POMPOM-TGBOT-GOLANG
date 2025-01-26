package utils

// Разные утилитарные функции

// Утилиты и вспомогательные функции (например, форматирование текста, snakeCase, обработка ошибок).

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
