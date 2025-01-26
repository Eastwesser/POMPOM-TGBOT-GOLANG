package services

import "context"

// Рассылка уведомлений

/*
	Сервисы — логика работы приложения:
		notify_service.go: Рассылка уведомлений всем подписчикам.
*/

func NotifyAllSubscribers(ctx context.Context, message string) error {
	// Получить список подписчиков и разослать им сообщение
	return nil
}
