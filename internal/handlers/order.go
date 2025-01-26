package handlers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обработка заказов

// order.go: Реализует оформление заказа через Telegram (например, с вводом количества товаров).

func HandleOrder(bot *tgbotapi.BotAPI, chatID int64, messageText string) {
	// Логика заказа (пока упрощенная)
	if messageText == "/order" {
		msg := tgbotapi.NewMessage(chatID, "Введите количество товаров для заказа:")
		bot.Send(msg)
		return
	}

	// Здесь можно проверить, что пользователь ввел цифру
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ваш заказ на %s товаров принят! 🎉", messageText))
	bot.Send(msg)
}
