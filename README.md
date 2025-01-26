# POMPON 🍎

Это MVP проект для телеграм-бота команды POMPON!

POMPON занимается подарками и упаковками, оформлением различных индивидуальных и корпоративных подарочных упаковок,
а также проводит мастер-классы! 🎨

## Структура проекта:

```
pompon-bot/
├── cmd/                    // Точка входа в приложение
│   └── pompon/             // Основной бинарник бота
│       └── main.go         // Точка запуска
├── internal/               // Вся внутренняя логика приложения
│   ├── config/             // Конфигурация и база данных
│   │   ├── config.go       // Настройки приложения
│   │   └── database.go     // Подключение к PostgreSQL
│   ├── handlers/           // Обработчики Telegram команд
│   │   ├── about.go        // Команда /about
│   │   ├── catalog.go      // Команда /catalog
│   │   ├── order.go        // Команда /order
│   │   └── subscribe.go    // Команда /subscribe
│   ├── keyboards/          // Клавиатуры для бота
│   │   └── keyboards.go    // Генерация клавиатур
│   ├── models/             // Модели данных и работа с PostgreSQL
│   │   ├── order.go        // Модель заказов
│   │   ├── product.go      // Модель товаров
│   │   └── subscriber.go   // Модель подписчиков
│   ├── services/           // Логика работы приложения
│   │   ├── catalog_service.go   // Логика каталога
│   │   ├── notify_service.go    // Логика уведомлений
│   │   ├── order_service.go     // Логика заказов
│   │   └── subscribe_service.go // Логика подписок
│   └── utils/              // Утилиты и вспомогательные функции
│       └── helpers.go      // Разные утилиты
├── sql/                    // SQL-скрипты для базы данных
│   ├── schema.sql          // Схема базы данных
│   ├── seed.sql            // Тестовые данные
│   └── README.md           // Описание структуры базы
├── .env-template           // Шаблон для переменных окружения
├── .gitignore              // Игнорируемые файлы
├── docker-compose.yml      // Docker Compose для запуска PostgreSQL и бота
├── Dockerfile              // Dockerfile для сборки приложения
├── go.mod                  // Go модули
├── README.md               // Документация проекта
└── .env                    // Настройки окружения (сейчас в .gitignore, используйте шаблон .env-template)
```#   P O M P O M - T G B O T - G O L A N G 
 
 