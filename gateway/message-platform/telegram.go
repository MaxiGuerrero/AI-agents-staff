package messageplatform

import (
	"fmt"
	"time"

	config "github.com/MaxiGuerrero/AI-agents-staff/gateway/config"
	"gopkg.in/telebot.v4"
)

// TelegramBot is a struct that implements the MessagePlatform interface for Telegram
type TelegramBot struct {
	telebot       *telebot.Bot
	securityToken string
}

// NewTelegramBot creates a new instance of TelegramBot with the provided token and security token for requests
func NewTelegramBot(token string, securityToken string) *TelegramBot {
	tb, err := telebot.NewBot(telebot.Settings{
		Token: token,
		Poller: &telebot.LongPoller{
			Timeout: 10 * time.Second,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("cannot configure bot: %v", err))
	}
	return &TelegramBot{
		telebot:       tb,
		securityToken: securityToken,
	}
}

// initHandlers: initializes the handlers for incoming Telegram messages and commands
func (tb *TelegramBot) InitHandlers(handlers []Handler) {
	authorized := initUsersAuthorized() // Initialize the Authorized map to keep track of authorized users
	for _, h := range handlers {
		if h.Command == "" || h.Function == nil {
			continue // Skip invalid handlers
		}
		if h.Command == "OnText" {
			tb.telebot.Handle(telebot.OnText, func(c telebot.Context) error {
				ctx := CommandContext{
					UserID: c.Sender().ID,
					Text:   c.Text(),
					Args:   c.Args(),
					Send: func(msg string) error {
						return c.Send(msg)
					},
					Authorized:    authorized,
					SecurityToken: &tb.securityToken,
				}
				return h.Function(ctx)
			})
		}
		tb.telebot.Handle(h.Command, func(c telebot.Context) error {
			userID := c.Sender().ID

			if h.Command != "/start" && !isAuthorized(userID) {
				return c.Send(
					"Must execute /start before using this bot. UNAUTHORIZED",
				)
			}
			ctx := CommandContext{
				UserID: c.Sender().ID,
				Text:   c.Text(),
				Args:   c.Args(),
				Send: func(msg string) error {
					return c.Send(msg)
				},
				Authorized:    authorized,
				SecurityToken: &tb.securityToken,
			}
			return h.Function(ctx)
		})
	}
}

// SendMessage sends a message to a Telegram chat
func (tb *TelegramBot) SendMessage(chatID int64, text string) error {
	// Implementation for sending message via Telegram
	return nil
}

func (tb *TelegramBot) Start() {
	go tb.telebot.Start() // Start the bot in a separate goroutine to allow it to run concurrently with the main application
	fmt.Println("Bot initialized and running...")
	fmt.Printf("Security token for authentication: %s\n", tb.securityToken)
}

func (tb *TelegramBot) Stop() {
	tb.telebot.Stop()
	fmt.Println("Bot stopped sucsuccessfully.")
}

// IsAuthorized checks if a user is authorized to interact with the bot based on their user ID
// Only the user with the ID specified in the configuration file is authorized to use the bot
// Must register the user ID in the configuration file "config.yaml" from root of project to allow access to the bot
func isAuthorized(userID int64) bool {
	whiteUserId := config.LoadConfig().WhiteUserId
	if userID == whiteUserId {
		return true
	}
	return false
}

func initUsersAuthorized() map[int64]bool {
	return make(map[int64]bool)
}
