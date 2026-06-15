package realTimeChat

import "time"

type ChatRequest struct {
	RequestId string `json:"request_id"`
	UserId    string `json:"user_id"`
	Channel   string `json:"channel"`
	Message   string `json:"message"`
}

type ChatResponse struct {
	RequestId string    `json:"request_id"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}
