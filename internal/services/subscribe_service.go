package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"pompon-bot-golang/internal/models"
)

// SubscribeService предоставляет методы для работы с подписками.
type SubscribeService struct {
	db *pgxpool.Pool
}

// NewSubscribeService создает новый экземпляр SubscribeService.
func NewSubscribeService(db *pgxpool.Pool) *SubscribeService {
	return &SubscribeService{db: db}
}

// AddSubscriber добавляет подписчика.
func (s *SubscribeService) AddSubscriber(ctx context.Context, telegramID int64) error {
	_, err := s.db.Exec(ctx, "INSERT INTO subscribers (telegram_id) VALUES ($1) ON CONFLICT DO NOTHING", telegramID)
	if err != nil {
		return fmt.Errorf("failed to add subscriber: %w", err)
	}

	return nil
}

// RemoveSubscriber удаляет подписчика.
func (s *SubscribeService) RemoveSubscriber(ctx context.Context, telegramID int64) error {
	_, err := s.db.Exec(ctx, "DELETE FROM subscribers WHERE telegram_id = $1", telegramID)
	if err != nil {
		return fmt.Errorf("failed to remove subscriber: %w", err)
	}

	return nil
}

// ListSubscribers возвращает список всех подписчиков.
func (s *SubscribeService) ListSubscribers(ctx context.Context) ([]models.Subscriber, error) {
	rows, err := s.db.Query(ctx, "SELECT id, telegram_id FROM subscribers")
	if err != nil {
		return nil, fmt.Errorf("failed to get subscribers: %w", err)
	}
	defer rows.Close()

	var subscribers []models.Subscriber
	for rows.Next() {
		var subscriber models.Subscriber
		if err := rows.Scan(&subscriber.ID, &subscriber.TelegramID); err != nil {
			return nil, fmt.Errorf("failed to scan subscriber: %w", err)
		}
		subscribers = append(subscribers, subscriber)
	}

	return subscribers, nil
}
