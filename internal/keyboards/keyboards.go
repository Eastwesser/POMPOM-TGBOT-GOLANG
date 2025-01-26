package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// Генерация клавиатур

/*
	Генерация клавиатур для удобного взаимодействия:
		Кнопки для выбора категории.
		Подтверждение заказа.
		Уведомления о подписке.
*/

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

func ConfirmOrderKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Подтвердить ✅"),
			tgbotapi.NewKeyboardButton("Отменить ❌"),
		),
	)
}

func SubscribeKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Подписаться 📩"),
			tgbotapi.NewKeyboardButton("Отписаться ❌"),
		),
	)
}
