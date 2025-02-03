package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"pompon-bot-golang/internal/services"
)

// HandleSubscribe обрабатывает команду /subscribe
func HandleSubscribe(bot *tgbotapi.BotAPI, chatID int64) {
	ctx := context.Background()
	err := services.AddSubscriber(ctx, chatID)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Ошибка при подписке. Попробуйте позже.")
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, "Вы успешно подписались на рассылку! 📩")
	bot.Send(msg)
}

// HandleUnsubscribe обрабатывает отписку
func HandleUnsubscribe(bot *tgbotapi.BotAPI, chatID int64) {
	ctx := context.Background()
	err := services.RemoveSubscriber(ctx, chatID)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Ошибка при отписке. Попробуйте позже.")
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, "Вы успешно отписались от рассылки. ❌")
	bot.Send(msg)
}
