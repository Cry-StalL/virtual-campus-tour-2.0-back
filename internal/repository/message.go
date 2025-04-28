package repository

import (
	"virtual-campus-tour-2.0-back/internal/model"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create 创建新留言
func (r *MessageRepository) Create(message *model.Message) error {
	return r.db.Create(message).Error
}

// GetByPanoramaID 获取指定全景图的留言列表
func (r *MessageRepository) GetByPanoramaID(panoramaID string) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Where("panorama_id = ?", panoramaID).
		Order("created_at DESC").
		Find(&messages).Error
	return messages, err
}

// GetByUserID 获取指定用户的所有留言
func (r *MessageRepository) GetByUserID(userID uint64) ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&messages).Error
	return messages, err
}
