package dto

import (
	"time"
)

// Position 表示留言在全景图中的位置
type Position struct {
	Longitude float64 `json:"longitude" binding:"required"` // 经度坐标
	Latitude  float64 `json:"latitude" binding:"required"`  // 纬度坐标
}

// CreateMessageRequest 创建留言请求
type CreateMessageRequest struct {
	Content    string   `json:"content" binding:"required,max=50"`
	UserID     uint64   `json:"userId" binding:"required"`
	Username   string   `json:"username" binding:"required"`
	PanoramaID string   `json:"panoramaId" binding:"required"`
	Position   Position `json:"position" binding:"required"` // 经纬度坐标
}

// MessageResponse 留言响应
type MessageResponse struct {
	MessageID  string    `json:"messageId"`
	Content    string    `json:"content"`
	UserID     uint64    `json:"userId"`
	Username   string    `json:"username"`
	PanoramaID string    `json:"panoramaId"`
	Location   string    `json:"location"` // 位置描述（可选）
	Position   Position  `json:"position"` // 经纬度坐标
	CreatedAt  time.Time `json:"createdAt"`
}

// GetMessagesRequest 获取留言列表请求
type GetMessagesRequest struct {
	PanoramaID string `form:"panoramaId" binding:"required"`
}

// GetUserMessagesRequest 获取用户留言请求
type GetUserMessagesRequest struct {
	UserID uint64 `json:"userId" binding:"required"`
}

// DeleteMessageRequest 删除留言请求
type DeleteMessageRequest struct {
	MessageID string `json:"messageId" binding:"required"` // 要删除的留言ID
}
