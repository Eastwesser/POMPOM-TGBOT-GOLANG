package models

// Модель товара

/*
	Модели базы данных для работы с PostgreSQL:
		product.go: Модель для товаров (id, название, цена, описание, категория).
*/

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CategoryID  int
}
