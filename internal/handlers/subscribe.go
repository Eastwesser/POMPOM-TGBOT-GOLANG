package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Подписка на новости

// subscribe.go: Обрабатывает подписку на рассылки.

func HandleSubscribe(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Вы успешно подписались на рассылку! 📩")
	bot.Send(msg)
}
