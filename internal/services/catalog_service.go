package services

// Логика для работы с каталогом

/*
	Сервисы — логика работы приложения:
		catalog_service.go: Предоставляет доступ к списку категорий и товаров.
		Используется обработчиками Telegram-команд, такими как /catalog.
*/

import (
	"context"
)

func GetCategories(ctx context.Context) []string {
	// В будущем данные будут подтягиваться из базы, сейчас — заглушка
	return []string{"Коробочки", "Открытки", "Обёртки"}
}
