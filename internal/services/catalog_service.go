package services

import (
	"context"
	"pompon-bot-golang/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// GetCategories возвращает список категорий
func GetCategories(db *pgxpool.Pool) ([]string, error) {
	rows, err := db.Query(context.Background(), "SELECT name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		categories = append(categories, name)
	}
	return categories, nil
}

// GetProductsByCategory возвращает товары по категории
func GetProductsByCategory(ctx context.Context, db *pgxpool.Pool, categoryID string) ([]models.Product, error) {
	rows, err := db.Query(ctx, "SELECT id, name, description, price FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
