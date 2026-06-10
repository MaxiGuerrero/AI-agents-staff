package handlers

import (
	platform "github.com/MaxiGuerrero/AI-agents-staff/chief-analyst-ai/message-platform"
)

func NewRealTimeChatHandler() platform.Handler {
	return platform.CreateHandler("OnText", realTimeChat)
}

func realTimeChat(ctx platform.CommandContext) error {
	if ctx.Authorized[ctx.UserID] {
		ctx.Send("You said: " + ctx.Text)
		return nil
	}
	return unauthorizedAccess(ctx)
}
