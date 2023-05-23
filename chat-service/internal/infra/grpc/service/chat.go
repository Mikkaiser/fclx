package service

import (
	"github.com/Mikkaiser/fclx/chat-service/internal/infra/grpc/pb"
	"github.com/Mikkaiser/fclx/chat-service/internal/usecase/chatcompletionstream"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
	ChatConfigStream            chatcompletionstream.ChatCompletionConfigInputDTO
	StreamChannel               chan chatcompletionstream.ChatCompletionOutputDTO
}

func NewChatService(chatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase, chatConfigStream chatcompletionstream.ChatCompletionConfigInputDTO, streamChannel chan chatcompletionstream.ChatCompletionOutputDTO) *ChatService {
	return &ChatService{
		ChatCompletionStreamUseCase: chatCompletionStreamUseCase,
		ChatConfigStream:            chatConfigStream,
		StreamChannel:               streamChannel,
	}
}
