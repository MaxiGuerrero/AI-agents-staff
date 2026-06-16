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

func (s *ChatService) ProcessMessage(ctx context.Context, message string, channel string, requestId string, userId int64) (string, error) {
	log.Println("Processing message for channel:", channel, "with request ID:", requestId)
	prompt := llm.Message{
		Role:    "user",
		Content: message,
	}
	response, err := s.llmClient.Chat(ctx, prompt, userId)
	if err != nil {
		return "", err
	}
	log.Println("message processed by LLM for User: ", userId)
	return response, nil
}
