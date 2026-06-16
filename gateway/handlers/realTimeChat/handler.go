package realTimeChat

import (
	"log"

	"github.com/MaxiGuerrero/AI-agents-staff/gateway/config"
	errors "github.com/MaxiGuerrero/AI-agents-staff/gateway/errors"
	platform "github.com/MaxiGuerrero/AI-agents-staff/gateway/message-platform"
)

func NewRealTimeChatHandler() platform.Handler {
	baseUrl := config.LoadConfig().BaseUrl
	channel := config.LoadConfig().Channel
	service := NewRealTimeChatService(baseUrl, channel)
	return platform.CreateHandler("OnText", func(ctx platform.CommandContext) error {
		return realTimeChat(service, ctx)
	})
}

func realTimeChat(service *RealTimeChatService, ctx platform.CommandContext) error {
	if ctx.Authorized[ctx.UserID] {
		userId := ctx.UserID
		response, err := service.ProcessMessage(ctx.Text, userId)
		if err != nil {
			log.Println("Failed to process message:", err)
			return err
		}
		ctx.Send(response)
		return nil
	}
	return errors.UnauthorizedAccess(ctx)
}
