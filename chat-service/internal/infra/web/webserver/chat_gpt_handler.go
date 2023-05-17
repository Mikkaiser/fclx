package web

import "github.com/Mikkaiser/fclx/chat-service/internal/usecase/chatcompletion"

type WebChatGPTHandler struct {
	CompletionUseCase chatcompletion.ChatCompletionUseCase
	Config            chatcompletion.ChatCompletionConfigInputDTO
	AuthToken         string
}

func NewWebChatGPTHandler(completionUseCase chatcompletion.ChatCompletionUseCase, config chatcompletion.ChatCompletionConfigInputDTO, AuthToken string) *WebChatGPTHandler {
	return &WebChatGPTHandler{
		CompletionUseCase: completionUseCase,
		Config:            config,
		AuthToken:         AuthToken,
	}
}

func ()
