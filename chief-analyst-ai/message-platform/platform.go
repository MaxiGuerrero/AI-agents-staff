package messageplatform

// MessagePlatform defines the interface for sending messages to different platforms
type MessagePlatform interface {
	SendMessage(chatID int64, text string) error
	Start()
	Stop()
}

// CommandContext represents the context of a command received from a user, including the user ID, the text of the command, and any arguments provided
type CommandContext struct {
	Authorized    map[int64]bool
	SecurityToken *string
	UserID        int64
	Text          string
	Args          []string
	Send          func(string) error
}

// Handler represents a command handler for any bot like telegram, whatsapp, etc.
// Containing the command and the function to execute when the command is received
type Handler struct {
	Command  string
	Function func(CommandContext) error
}

func CreateHandler(command string, function func(CommandContext) error) Handler {
	return Handler{
		Command:  command,
		Function: function,
	}
}
