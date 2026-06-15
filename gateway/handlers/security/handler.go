package security

import (
	errors "github.com/MaxiGuerrero/AI-agents-staff/gateway/errors"
	platform "github.com/MaxiGuerrero/AI-agents-staff/gateway/message-platform"
)

func NewStartHandler() platform.Handler {
	return platform.CreateHandler("/start", securityLayer)
}

// Handler that must compare the password sent by the user with the security token stored in the environment variable and return an error if they do not match
func securityLayer(ctx platform.CommandContext) error {
	if len(ctx.Args) == 0 {
		return errors.PasswordRequired(ctx)
	}
	password := ctx.Args[0]
	if password == "" {
		return errors.PasswordRequired(ctx)
	}
	securityToken := *ctx.SecurityToken
	if password != securityToken {
		return errors.UnauthorizedAccess(ctx)
	}
	ctx.Authorized[ctx.UserID] = true // Mark the user as authorized in the context
	ctx.Send("Access granted. Welcome to the Chief Analyst AI.")
	return nil
}
