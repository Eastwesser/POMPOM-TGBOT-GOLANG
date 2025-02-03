package handlers

import (
	"fmt"
	"pompon-bot-golang/internal/keyboards"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userState = make(map[int64]string) // Хранение состояния пользователя

// HandleOrder обрабатывает команду /order
func HandleOrder(bot *tgbotapi.BotAPI, chatID int64, messageText string) {
	if messageText == "/order" {
		msg := tgbotapi.NewMessage(chatID, "Выберите категорию:")
		msg.ReplyMarkup = keyboards.CatalogKeyboard()
		bot.Send(msg)
		userState[chatID] = "waiting_for_category"
		return
	}

	switch userState[chatID] {
	case "waiting_for_category":
		msg := tgbotapi.NewMessage(chatID, "Введите количество товаров:")
		bot.Send(msg)
		userState[chatID] = "waiting_for_quantity"
	case "waiting_for_quantity":
		quantity, err := strconv.Atoi(messageText)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите число.")
			bot.Send(msg)
			return
		}
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Ваш заказ на %d товаров принят! 🎉", quantity))
		bot.Send(msg)
		delete(userState, chatID)
	}
}
