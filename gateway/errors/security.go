package errors

import (
	"fmt"

	platform "github.com/MaxiGuerrero/AI-agents-staff/gateway/message-platform"
)

func PasswordRequired(ctx platform.CommandContext) error {
	errorMessage := "Password is required. Please provide the security token to access the Chief Analyst AI."
	ctx.Send(errorMessage)
	return fmt.Errorf("%s", errorMessage)
}

func UnauthorizedAccess(ctx platform.CommandContext) error {
	errorMessage := "Unauthorized access. Please provide the correct security token to access the Chief Analyst AI."
	ctx.Send(errorMessage)
	return fmt.Errorf("%s", errorMessage)
}
