package main

import (
	"log"
	"net/http"

	chat "github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/chat"
	"github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/config"
	"github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/llm"
	"github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/middlewares"
)

func main() {
	// Load configuration
	var conf = config.LoadConfig()
	// Initialize services and handlers
	llmClient := llm.NewLLM(conf.LlmBaseUrl)
	chatService := chat.NewChatService(llmClient)
	chatHandler := chat.NewHandler(chatService)
	// Set up HTTP server and routes
	mux := http.NewServeMux()
	mux.HandleFunc("POST /chat", chatHandler.HandleChat)
	log.Print("Server listening on :8080")

	err := http.ListenAndServe(
		":8080",
		middlewares.LoggingMiddleware(mux),
	)

	if err != nil {
		log.Fatal(err)
	}
}
