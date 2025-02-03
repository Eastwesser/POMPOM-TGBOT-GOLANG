package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pompon-bot-golang/internal/config"
	"pompon-bot-golang/internal/handlers"
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

	// Пример использования ToSnakeCase
	camelCase := "CamelCaseString"
	snakeCase := utils.ToSnakeCase(camelCase)
	utils.LogInfo(logger, fmt.Sprintf("Snake case: %s", snakeCase)) // Вывод: camel_case_string

	// Пример использования IsEmpty
	emptyStr := ""
	utils.LogInfo(logger, fmt.Sprintf("Is empty: %v", utils.IsEmpty(emptyStr))) // Вывод: true

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
					handlers.HandleCatalog(db, bot, update.Message.Chat.ID)
				case "order":
					handlers.HandleOrder(bot, update.Message.Chat.ID, update.Message.Text)
				case "subscribe":
					handlers.HandleSubscribe(bot, update.Message.Chat.ID)
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
