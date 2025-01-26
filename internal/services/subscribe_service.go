package services

import (
	"context"
	"pompon-bot/internal/models"
)

// Логика подписок

/*
	Сервисы — логика работы приложения:
		subscribe_service.go: Добавление пользователей в список подписчиков.
*/

// Добавление пользователей в список подписчиков
func AddSubscriber(ctx context.Context, telegramID int64) error {
	// Здесь будет логика сохранения подписчика в базе данных
	return nil
}

func RemoveSubscriber(ctx context.Context, telegramID int64) error {
	// Логика удаления подписчика
	return nil
}

func ListSubscribers(ctx context.Context) ([]models.Subscriber, error) {
	// Получение всех подписчиков
	return []models.Subscriber{}, nil
}
