package main

import (
	"log"
	"net/http"

	chat "github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/chat"
	"github.com/MaxiGuerrero/AI-agents-staff/chief-analyst/middlewares"
)

func main() {
	chatService := chat.NewChatService()
	chatHandler := chat.NewHandler(chatService)
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
