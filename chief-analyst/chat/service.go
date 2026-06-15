package chat

import (
	"context"
	"log"

	"github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/llm"
)

type ChatService struct {
	llmClient llm.ILlm
}

func NewChatService(llmClient llm.ILlm) *ChatService {
	return &ChatService{
		llmClient: llmClient,
	}
}

func (s *ChatService) ProcessMessage(ctx context.Context, message string, channel string, requestId string) (string, error) {
	// Placeholder for actual message processing logic
	log.Println("Processing message for channel:", channel, "with request ID:", requestId)
	prompt := []llm.Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: message},
	}
	response, err := s.llmClient.Chat(ctx, prompt)
	if err != nil {
		return "", err
	}
	return response, nil
}
