package realTimeChat

import "github.com/MaxiGuerrero/AI-agents-staff/gateway/utils"

type RealTimeChatService struct {
	httpClient   *ChiefAnalystHttpClient
	generateUUID func() string
	channel      string
}

func NewRealTimeChatService(baseUrl string, channel string) *RealTimeChatService {
	return &RealTimeChatService{
		httpClient:   NewChiefAnalystServiceClient(baseUrl),
		generateUUID: utils.GenerateUUID,
		channel:      channel,
	}
}

func (s *RealTimeChatService) ProcessMessage(message string, userId int64) (string, error) {
	requestId := s.generateUUID()
	response, err := s.httpClient.SendMessageChiefAnalyst(message, s.channel, requestId, userId)
	if err != nil {
		return "", err
	}
	return response.Message, nil
}
