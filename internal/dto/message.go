package dto

// CreateMessageRequest 创建留言请求
type CreateMessageRequest struct {
	Content    string `json:"content" binding:"required"`
	UserID     uint64 `json:"userId" binding:"required"`
	Username   string `json:"username" binding:"required"`
	PanoramaID string `json:"panoramaId" binding:"required"`
	Location   string `json:"location" `
}

// MessageResponse 留言响应
type MessageResponse struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	UserID     int    `json:"userId"`
	Username   string `json:"username"`
	PanoramaID string `json:"panoramaId"`
	CreatedAt  string `json:"createdAt"`
}

// GetMessagesRequest 获取留言列表请求
type GetMessagesRequest struct {
	PanoramaID string `form:"panoramaId" binding:"required"`
}

// GetUserMessagesRequest 获取用户留言请求
type GetUserMessagesRequest struct {
	UserID uint64 `json:"userId" binding:"required"`
}
