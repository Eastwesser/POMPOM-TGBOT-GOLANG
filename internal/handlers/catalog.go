package handlers

import (
	"context"
	"fmt"
	"pompon-bot-golang/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Логика работы с каталогом товаров

// catalog.go: Обрабатывает команду /catalog, показывает категории и список товаров.

func HandleCatalog(ctx context.Context, bot *tgbotapi.BotAPI, chatID int64) {
	categories := services.GetCategories(ctx)
	message := "Категории:\n"

	for _, category := range categories {
		message += fmt.Sprintf("- %s\n", category)
	}

	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}
