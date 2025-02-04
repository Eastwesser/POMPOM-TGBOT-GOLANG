package services

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"pompon-bot-golang/internal/models"
)

// OrderService предоставляет методы для работы с заказами.
type OrderService struct {
	db *pgxpool.Pool
}

// NewOrderService создает новый экземпляр OrderService.
func NewOrderService(db *pgxpool.Pool) *OrderService {
	return &OrderService{db: db}
}

// CreateOrder создает новый заказ.
func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Вставляем заказ
	_, err = tx.Exec(ctx, "INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4)",
		order.UserID, order.Product.ID, order.Quantity, order.Status)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// UpdateOrderStatus обновляет статус заказа.
func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID int, status string) error {
	_, err := s.db.Exec(ctx, "UPDATE orders SET status = $1 WHERE id = $2", status, orderID)
	if err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}

	return nil
}

// GetOrdersByUser возвращает заказы пользователя.
func (s *OrderService) GetOrdersByUser(ctx context.Context, userID int) ([]models.Order, error) {
	rows, err := s.db.Query(ctx, "SELECT id, product_id, quantity, status FROM orders WHERE user_id = $1", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.Product.ID, &order.Quantity, &order.Status); err != nil {
			return nil, fmt.Errorf("failed to scan order: %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}
