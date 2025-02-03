package services

import (
	"context"
	"pompon-bot-golang/internal/models"
)

// Логика подписок

/*
	Сервисы — логика работы приложения:
		subscribe_service.go: Управляет списком подписчиков.
		Используется для обработки команд, таких как /subscribe и /unsubscribe.
*/

// AddSubscriber добавляет подписчика
func AddSubscriber(ctx context.Context, db *pgxpool.Pool, telegramID int64) error {
	_, err := db.Exec(ctx, "INSERT INTO subscribers (telegram_id) VALUES ($1) ON CONFLICT DO NOTHING", telegramID)
	return err
}

// RemoveSubscriber удаляет подписчика
func RemoveSubscriber(ctx context.Context, db *pgxpool.Pool, telegramID int64) error {
	_, err := db.Exec(ctx, "DELETE FROM subscribers WHERE telegram_id = $1", telegramID)
	return err
}}

// Получение всех подписчиков
func ListSubscribers(ctx context.Context) ([]models.Subscriber, error) {
	// Получение всех подписчиков
	return []models.Subscriber{}, nil
}
