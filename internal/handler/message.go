package handler

import (
	"log"
	"net/http"

	"virtual-campus-tour-2.0-back/internal/dto"
	"virtual-campus-tour-2.0-back/internal/service"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service *service.MessageService
}

func NewMessageHandler(service *service.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

// CreateMessage 创建新留言
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var req dto.CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	message, err := h.service.CreateMessage(req.Content, req.UserID, req.Username, req.PanoramaID, req.Location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建留言失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建成功",
		"data":    message,
	})
}

// GetMessages 获取留言列表
func (h *MessageHandler) GetMessages(c *gin.Context) {
	var req dto.GetMessagesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误",
		})
		return
	}

	messages, err := h.service.GetMessagesByPanoramaID(req.PanoramaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取留言列表失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    messages,
	})
}

// GetUserMessages 获取用户的所有留言
func (h *MessageHandler) GetUserMessages(c *gin.Context) {
	var req dto.GetUserMessagesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 添加日志记录
	log.Printf("获取用户留言请求: userId=%d", req.UserID)

	messages, err := h.service.GetMessagesByUserID(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取留言失败",
			"data":    nil,
		})
		return
	}

	// 转换消息格式
	var responseMessages []dto.MessageResponse
	for _, msg := range messages {
		responseMessages = append(responseMessages, dto.MessageResponse{
			ID:         msg.ID,
			Content:    msg.Content,
			UserID:     int(msg.UserID),
			Username:   msg.Username,
			PanoramaID: msg.PanoramaID,
			Location:   msg.Location,
			CreateTime: msg.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取成功",
		"data":    responseMessages,
	})
}
