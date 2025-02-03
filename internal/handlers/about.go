package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// HandleStart обрабатывает команду /start
func HandleStart(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в магазин POMPON! 🎁\nВыберите действие:")
	msg.ReplyMarkup = MainMenuKeyboard()
	bot.Send(msg)
}

// MainMenuKeyboard создает клавиатуру для главного меню
func MainMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔹 О нас (/about)"),
			tgbotapi.NewKeyboardButton("📦 Каталог (/catalog)"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🛒 Заказать (/order)"),
			tgbotapi.NewKeyboardButton("🔔 Подписка (/subscribe)"),
		),
	)
}

// HandleAbout отвечает на команду /about
func HandleAbout(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Мы — магазин POMPON! 🎁\nЗдесь вы найдете лучшие коробочки, открытки и обёртки для ваших подарков.")
	bot.Send(msg)
}
