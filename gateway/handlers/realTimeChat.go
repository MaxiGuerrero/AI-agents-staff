package handlers

import (
	errors "github.com/MaxiGuerrero/AI-agents-staff/gateway/errors"
	platform "github.com/MaxiGuerrero/AI-agents-staff/gateway/message-platform"
)

func NewRealTimeChatHandler() platform.Handler {
	return platform.CreateHandler("OnText", realTimeChat)
}

func realTimeChat(ctx platform.CommandContext) error {
	if ctx.Authorized[ctx.UserID] {
		ctx.Send("You said: " + ctx.Text)
		return nil
	}
	return errors.UnauthorizedAccess(ctx)
}
