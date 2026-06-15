package llm

import (
	"context"
	"log"
)

type LlmClient struct {
	baseUrl string
}

func NewLLM(baseUrl string) *LlmClient {
	return &LlmClient{
		baseUrl: baseUrl,
	}
}

func (l *LlmClient) Chat(ctx context.Context, prompt []Message) (string, error) {
	log.Println("LLM client received prompt:", prompt)
	response := "This is a response from the LLM based on the prompt."
	return response, nil
}
