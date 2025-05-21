package model

import (
	"time"
)

// Message 表示用户留言
type Message struct {
	ID         string    `gorm:"primaryKey;type:char(36)"`
	Content    string    `gorm:"type:varchar(50);not null"`
	UserID     uint64    `gorm:"type:bigint unsigned;not null"`
	Username   string    `gorm:"type:varchar(50);not null"`
	PanoramaID string    `gorm:"type:varchar(36);not null"`
	Location   string    `gorm:"type:varchar(255);not null"`  // 位置描述
	Longitude  float64   `gorm:"type:decimal(10,6);not null"` // 经度坐标
	Latitude   float64   `gorm:"type:decimal(10,6);not null"` // 纬度坐标
	CreatedAt  time.Time `gorm:"not null"`
	UpdatedAt  time.Time
}
