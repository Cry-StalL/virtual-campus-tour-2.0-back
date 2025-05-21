package service

import (
	"time"

	"virtual-campus-tour-2.0-back/internal/dto"
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
func (s *MessageService) CreateMessage(req *dto.CreateMessageRequest) (*dto.MessageResponse, error) {
	message := &model.Message{
		ID:         uuid.New().String(),
		Content:    req.Content,
		UserID:     req.UserID,
		Username:   req.Username,
		PanoramaID: req.PanoramaID,
		Location:   "未设置位置", // 设置默认位置描述
		Longitude:  req.Position.Longitude,
		Latitude:   req.Position.Latitude,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.repo.Create(message); err != nil {
		return nil, err
	}

	return &dto.MessageResponse{
		MessageID:  message.ID,
		Content:    message.Content,
		UserID:     message.UserID,
		Username:   message.Username,
		PanoramaID: message.PanoramaID,
		Location:   message.Location,
		Position: dto.Position{
			Longitude: message.Longitude,
			Latitude:  message.Latitude,
		},
		CreatedAt: message.CreatedAt,
	}, nil
}

// GetMessagesByPanoramaID 获取指定全景图的留言列表
func (s *MessageService) GetMessagesByPanoramaID(panoramaID string) ([]*dto.MessageResponse, error) {
	messages, err := s.repo.GetByPanoramaID(panoramaID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.MessageResponse, 0, len(messages))
	for _, msg := range messages {
		responses = append(responses, &dto.MessageResponse{
			MessageID:  msg.ID,
			Content:    msg.Content,
			UserID:     msg.UserID,
			Username:   msg.Username,
			PanoramaID: msg.PanoramaID,
			Location:   msg.Location,
			Position: dto.Position{
				Longitude: msg.Longitude,
				Latitude:  msg.Latitude,
			},
			CreatedAt: msg.CreatedAt,
		})
	}

	return responses, nil
}

// GetMessagesByUserID 获取指定用户的所有留言
func (s *MessageService) GetMessagesByUserID(userID uint64) ([]*dto.UserMessageResponse, error) {
	messages, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.UserMessageResponse, 0, len(messages))
	for _, msg := range messages {
		responses = append(responses, &dto.UserMessageResponse{
			MessageID:  msg.ID,
			Content:    msg.Content,
			UserID:     msg.UserID,
			Username:   msg.Username,
			PanoramaID: msg.PanoramaID,
			Location:   msg.Location,
			CreateTime: msg.CreatedAt,
		})
	}

	return responses, nil
}

// DeleteMessage 删除留言
func (s *MessageService) DeleteMessage(messageID string) error {
	return s.repo.Delete(messageID)
}
