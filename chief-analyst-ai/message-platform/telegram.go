package messageplatform

import "gopkg.in/telebot.v4"

// TelegramBot is a struct that implements the MessagePlatform interface for Telegram
type TelegramBot struct {
	telebot *telebot.Bot
}

func NewTelegramBot(token string) *TelegramBot {
	tb, err := telebot.NewBot(telebot.Settings{
		Token: token,
	})
	if err != nil {
		panic("Bot cannot be configured")
	}

	return &TelegramBot{
		telebot: tb,
	}
}

// SendMessage sends a message to a Telegram chat
func (tb *TelegramBot) SendMessage(chatID int64, text string) error {
	// Implementation for sending message via Telegram
	return nil
}
