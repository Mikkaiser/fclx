package chatcompletionstream

import (
	"context"
	"errors"

	"github.com/Mikkaiser/fclx/chat-service/internal/domain/gateway"
	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
}

type ChatCompletionConfigInputDTO struct {
	Model                string
	ModelMaxTokens       int
	Temperature          float32
	TopP                 float32
	N                    int
	Stop                 []string
	MaxTokens            int
	PresencePenalty      float32
	FrequencyPenalty     float32
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
	ChatID      string
	UserID      string
	UserMessage string
	Config      ChatCompletionConfigInputDTO
}

type ChatCompletionOutputDTO struct {
	ChatID  string
	UserID  string
	Content string
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiCLient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		ChatGateway:  chatGateway,
		OpenAiClient: openAiCLient,
	}
}

func (uc *ChatCompletionUseCase) Execute(ctx context.Context, input ChatCompletionInputDTO) (*ChatCompletionOutputDTO, error) {
	chat, err := uc.ChatGateway.FindChatByID(ctx, input.ChatID)

	if err != nil {
		if err.Error() == "Chat not found" {

			//create new chat (entity)
			chat, err = createNewChat(input)
			if err != nil {
				return nil, errors.New("Error creating new chat: " + err.Error())
			}

			//save on database
			err = uc.ChatGateway.CreateChat(ctx, chat)
			if err != nil {
				return nil, errors.New("Error persisting new chat: " + err.Error())
			}
		}
	}
}
