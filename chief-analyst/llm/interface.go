package llm

import "context"

type ILlm interface {
	Chat(
		ctx context.Context,
		prompt []Message,
	) (string, error)
}
