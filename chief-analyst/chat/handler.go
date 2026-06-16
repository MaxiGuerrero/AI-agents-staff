package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	chatService *ChatService
}

func NewHandler(chatService *ChatService) *Handler {
	return &Handler{
		chatService: chatService,
	}
}

func (h *Handler) HandleChat(w http.ResponseWriter, r *http.Request) {
	var req ChatRequest
	var res ChatResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Failed to decode request body:", err)
		return
	}
	message, err := h.chatService.ProcessMessage(r.Context(), req.Message, req.Channel, req.RequestId, req.UserId)
	if err != nil {
		http.Error(w, "Failed to process message", http.StatusInternalServerError)
		log.Println("Failed to process message:", err)
		return
	}
	res = ChatResponse{
		RequestId: req.RequestId,
		Message:   message,
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
