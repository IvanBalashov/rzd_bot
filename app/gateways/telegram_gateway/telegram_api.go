package telegram_gateway

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	)
type TelegramAPI struct {
	APIKey string
}

func NewTelegramAPI(apiKey string) TelegramAPI {
	return TelegramAPI{
		APIKey:apiKey,
	}
}

func (t *TelegramAPI) ReadMessages() {
	fmt.Printf("test")
}