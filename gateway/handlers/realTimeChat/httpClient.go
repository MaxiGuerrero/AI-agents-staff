package realTimeChat

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ChiefAnalystHttpClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewChiefAnalystServiceClient(baseURL string) *ChiefAnalystHttpClient {
	return &ChiefAnalystHttpClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *ChiefAnalystHttpClient) SendMessageChiefAnalyst(message string, channel string, requestId string, userId int64) (ChatResponse, error) {
	request := ChatRequest{
		RequestId: requestId,
		UserId:    userId,
		Channel:   channel,
		Message:   message,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return ChatResponse{}, err
	}
	req, err := http.NewRequest("POST", s.baseURL+"/chat", bytes.NewBuffer(body)) // Placeholder for actual request body
	if err != nil {
		return ChatResponse{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return ChatResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ChatResponse{}, err
	}
	var chatResponse ChatResponse
	err = json.NewDecoder(resp.Body).Decode(&chatResponse)
	if err != nil {
		return ChatResponse{}, err
	}
	return chatResponse, nil
}
