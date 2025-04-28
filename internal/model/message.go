package model

import (
	"time"
)

// Message 表示用户留言
type Message struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	UserID     uint64    `json:"userId" gorm:"type:bigint unsigned;not null"`
	Username   string    `json:"username" gorm:"not null"`
	PanoramaID string    `json:"panoramaId" gorm:"not null"`
	Location   string    `json:"location" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
