package llm

// Models for messages

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatSession struct {
	Messages []Message `json:"messages"`
	UserId   int64     `json:"user_id"`
}

// Models for request - response with LLM Server

type LLMRequest struct {
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Timing struct {
	PredictedMs float64 `json:"predicted_ms"`
}

type Usage struct {
	TokensUsage int `json:"completion_tokens"`
}

type LLMResponse struct {
	Choices []Choice `json:"choices"`
	Model   string   `json:"model"`
	Timing  Timing   `json:"timings"`
	Usage   Usage    `json:"usage"`
}
