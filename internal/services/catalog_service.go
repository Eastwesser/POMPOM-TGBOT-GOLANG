package services

// Логика для работы с каталогом

/*
	Сервисы — логика работы приложения:
		catalog_service.go: Получение списка товаров или категорий из базы.
*/

import (
	"context"
)

func GetCategories(ctx context.Context) []string {
	// В будущем будет из базы, пока mock-данные
	return []string{"Коробочки", "Открытки", "Обёртки"}
}
