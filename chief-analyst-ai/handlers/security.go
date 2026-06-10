package handlers

import (
	"fmt"

	platform "github.com/MaxiGuerrero/AI-agents-staff/chief-analyst-ai/message-platform"
)

func NewStartHandler() platform.Handler {
	return platform.CreateHandler("/start", securityLayer)
}

// Handler that must compare the password sent by the user with the security token stored in the environment variable and return an error if they do not match
func securityLayer(ctx platform.CommandContext) error {
	if len(ctx.Args) == 0 {
		return passwordRequired(ctx)
	}
	password := ctx.Args[0]
	if password == "" {
		return passwordRequired(ctx)
	}
	securityToken := *ctx.SecurityToken
	if password != securityToken {
		return unauthorizedAccess(ctx)
	}
	ctx.Authorized[ctx.UserID] = true // Mark the user as authorized in the context
	ctx.Send("Access granted. Welcome to the Chief Analyst AI.")
	return nil
}

func passwordRequired(ctx platform.CommandContext) error {
	errorMessage := "Password is required. Please provide the security token to access the Chief Analyst AI."
	ctx.Send(errorMessage)
	return fmt.Errorf("%s", errorMessage)
}

func unauthorizedAccess(ctx platform.CommandContext) error {
	errorMessage := "Unauthorized access. Please provide the correct security token to access the Chief Analyst AI."
	ctx.Send(errorMessage)
	return fmt.Errorf("%s", errorMessage)
}
