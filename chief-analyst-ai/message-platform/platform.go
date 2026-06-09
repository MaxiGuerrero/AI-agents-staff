package messageplatform

// MessagePlatform defines the interface for sending messages to different platforms
type MessagePlatform interface {
	SendMessage(chatID int64, text string) error
}
