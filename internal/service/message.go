package service

import (
	"time"

	"virtual-campus-tour-2.0-back/internal/model"
	"virtual-campus-tour-2.0-back/internal/repository"

	"github.com/google/uuid"
)

type MessageService struct {
	repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// CreateMessage 创建新留言
func (s *MessageService) CreateMessage(content, userID, username, panoramaID string) (*model.Message, error) {
	message := &model.Message{
		ID:         uuid.New().String(),
		Content:    content,
		UserID:     userID,
		Username:   username,
		PanoramaID: panoramaID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := s.repo.Create(message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

// GetMessagesByPanoramaID 获取指定全景图的留言列表
func (s *MessageService) GetMessagesByPanoramaID(panoramaID string) ([]model.Message, error) {
	return s.repo.GetByPanoramaID(panoramaID)
}
