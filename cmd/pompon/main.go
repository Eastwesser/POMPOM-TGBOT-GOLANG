package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pompon-bot-golang/internal/config"
	"pompon-bot-golang/internal/handlers"
	"pompon-bot-golang/internal/models"
	"pompon-bot-golang/internal/services"
	"pompon-bot-golang/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var logger *log.Logger

func init() {
	// Создаем логгер один раз при старте приложения
	logger = utils.CreateLogger("bot.log", utils.LoggerConfig{
		Prefix: "POMPON-BOT",
		Level:  utils.LevelInfo,
	})
}

func main() {
	utils.LogInfo(logger, "Бот запускается...")

	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db := config.ConnectDB(cfg)
	defer db.Close()

	// Инициализация Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		utils.LogError(logger, fmt.Sprintf("Failed to initialize bot: %v", err))
		os.Exit(1)
	}

	bot.Debug = true
	utils.LogInfo(logger, fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	// Инициализация сервисов
	catalogService := services.NewCatalogService(db)
	orderService := services.NewOrderService(db)
	notifyService := services.NewNotifyService(db, bot) // Передаем bot в NotifyService
	subscribeService := services.NewSubscribeService(db)

	// Пример использования сервисов
	categories, err := catalogService.GetCategories(context.Background())
	if err != nil {
		utils.LogError(logger, fmt.Sprintf("Failed to get categories: %v", err))
	} else {
		for _, category := range categories {
			utils.LogInfo(logger, fmt.Sprintf("Category: %s", category))
		}
	}

	order := models.Order{
		UserID:   1,
		Product:  models.Product{ID: 1},
		Quantity: 2,
		Status:   "pending",
	}
	if err := orderService.CreateOrder(context.Background(), order); err != nil {
		utils.LogError(logger, fmt.Sprintf("Failed to create order: %v", err))
	}

	// Канал для обработки системных сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message != nil {
				// Логируем входящее сообщение с полями
				fields := map[string]interface{}{
					"chat_id": update.Message.Chat.ID,
					"text":    update.Message.Text,
				}
				utils.LogWithFields(logger, utils.LevelInfo, "Received message", fields)

				switch update.Message.Command() {
				case "start":
					handlers.HandleStart(update, bot)
				case "about":
					handlers.HandleAbout(update, bot)
				case "catalog":
					handlers.HandleCatalog(catalogService, bot, update.Message.Chat.ID)
				case "order":
					handlers.HandleOrder(orderService, bot, update.Message.Chat.ID, update.Message.Text)
				case "subscribe":
					handlers.HandleSubscribe(subscribeService, bot, update.Message.Chat.ID)
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /start для начала.")
					bot.Send(msg)
				}
			} else if update.CallbackQuery != nil {
				handlers.HandleCallbackQuery(db, bot, update.CallbackQuery)
			}
		}
	}()

	// Ожидание сигнала завершения
	<-stop
	utils.LogInfo(logger, "Бот завершает работу.")
}
