package handlers

import (
	platform "github.com/MaxiGuerrero/AI-agents-staff/chief-analyst-ai/message-platform"
)

// Registers here all GetHandlers function that will return the handlers to be used in the main application,.
// This way we can keep all the handlers organized in one place and easily add new handlers as needed without modifying the main application code.
func GetHandlers() []platform.Handler {
	return []platform.Handler{
		NewStartHandler(),
		NewRealTimeChatHandler(),
	}
}
