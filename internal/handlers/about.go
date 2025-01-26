package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Информация о магазине

// about.go: Команда /about с информацией о магазине.

func CatalogKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Коробочки 🎁"),
			tgbotapi.NewKeyboardButton("Открытки 🧧"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Обёртки 🎀"),
		),
	)
}
