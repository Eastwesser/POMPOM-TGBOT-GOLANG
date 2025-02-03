package main

import (
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
	logger = utils.CreateLogger("bot.log")
}

func main() {
	logger.Println("Бот запускается...")
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db := config.ConnectDB(cfg)
	defer db.Close()

	// Инициализация Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		logger.Fatalf("Failed to initialize bot: %v", err)
	}

	bot.Debug = true
	logger.Printf("Authorized on account %s", bot.Self.UserName)

	// Канал для обработки системных сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message != nil {
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

	<-stop
	logger.Println("Бот завершает работу.")
}
