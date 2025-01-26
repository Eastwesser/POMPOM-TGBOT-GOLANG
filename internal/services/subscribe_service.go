package services

import (
	"context"
	"pompon-bot/internal/models"
)

// Логика подписок

/*
	Сервисы — логика работы приложения:
		subscribe_service.go: Управляет списком подписчиков.
		Используется для обработки команд, таких как /subscribe и /unsubscribe.
*/

// Добавление пользователя в список подписчиков
func AddSubscriber(ctx context.Context, telegramID int64) error {
	// Здесь будет логика сохранения подписчика в базе данных
	return nil
}

// Удаление пользователя из списка подписчиков
func RemoveSubscriber(ctx context.Context, telegramID int64) error {
	// Логика удаления подписчика
	return nil
}

// Получение всех подписчиков
func ListSubscribers(ctx context.Context) ([]models.Subscriber, error) {
	// Получение всех подписчиков
	return []models.Subscriber{}, nil
}
