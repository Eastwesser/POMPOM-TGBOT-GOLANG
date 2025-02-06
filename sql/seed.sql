-- Демо-данные для тестирования

-- Добавление категорий
INSERT INTO categories (name)
VALUES ('Коробочки'),
       ('Открытки'),
       ('Обёртки');

-- Добавление товаров
INSERT INTO products (name, description, price, category_id)
VALUES ('Подарочная коробочка', 'Идеально подходит для небольших подарков', 300.00, 1),
       ('Ручная открытка', 'Уникальная открытка ручной работы', 200.00, 2),
       ('Крафтовая обёртка', 'Красивая обёрточная бумага для подарков', 150.00, 3);

-- Добавление подписчиков
INSERT INTO subscribers (telegram_id)
VALUES (123456789),
       (987654321);

-- Добавление заказов
INSERT INTO orders (user_id, product_id, quantity, status)
VALUES (123456789, 1, 2, 'completed'),
       (987654321, 3, 1, 'pending');
