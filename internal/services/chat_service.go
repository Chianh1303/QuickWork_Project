package services

import (
	"QuickWork/internal/models"
	"QuickWork/internal/repositories"
)

type ChatService interface {
	GetChatHistory(appID uint) ([]models.Message, error)
}

type chatService struct {
	chatRepo repositories.ChatRepository
}

func NewChatService(chatRepo repositories.ChatRepository) ChatService {
	return &chatService{chatRepo: chatRepo}
}

func (s *chatService) GetChatHistory(appID uint) ([]models.Message, error) {
	return s.chatRepo.GetMessagesByApplicationID(appID)
}
