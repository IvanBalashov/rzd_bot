package telegram_gateway

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)
type TelegramAPI struct {
	Bot tgbotapi.BotAPI
	APIKey string
}

func NewTelegramAPI(apiKey string) TelegramAPI {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Printf("Telegram_API->NewTelegramAPI: Can't create bot instance, - %s\n", err)
		return TelegramAPI{}
	}
	return TelegramAPI{
		APIKey:apiKey,
		Bot: *bot,
	}
}

func (t *TelegramAPI) ReadMessages() {
	tgbotapi.NewUpdate(0)
	fmt.Printf("test")
}