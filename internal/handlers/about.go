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

// HandleAbout отвечает на команду /about
func HandleAbout(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// Пример логики для обработки команды /about
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Информация о магазине")
	_, _ = bot.Send(message)
}
