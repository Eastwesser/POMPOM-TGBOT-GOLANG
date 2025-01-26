package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"pompon-bot-golang/internal/config"
	"pompon-bot-golang/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Инициализация Telegram Bot API
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIKey)
	if err != nil {
		log.Fatalf("Failed to initialize bot: %v", err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Канал для обработки системных сигналов (выход по Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Настройка обновлений Telegram
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			// Обработка команд
			switch update.Message.Command() {
			case "about":
				handlers.HandleAbout(bot, update.Message.Chat.ID)
			case "catalog":
				handlers.HandleCatalog(nil, bot, update.Message.Chat.ID) // TODO: Передать контекст базы данных
			case "order":
				handlers.HandleOrder(bot, update.Message.Chat.ID, update.Message.Text)
			case "subscribe":
				handlers.HandleSubscribe(bot, update.Message.Chat.ID)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Попробуйте /about, /catalog, /order или /subscribe.")
				bot.Send(msg)
			}
		}
	}()

	// Ожидание завершения работы
	<-stop
	log.Println("Bot stopped gracefully.")
}
