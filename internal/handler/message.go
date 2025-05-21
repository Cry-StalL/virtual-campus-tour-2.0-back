package handler

import (
	"fmt"
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

// CreateMessage 创建留言
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var req dto.CreateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的请求参数: " + err.Error(),
		})
		return
	}

	response, err := h.service.CreateMessage(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "创建留言失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetMessages 获取留言列表
func (h *MessageHandler) GetMessages(c *gin.Context) {
	var req dto.GetMessagesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的请求参数: " + err.Error(),
		})
		return
	}

	messages, err := h.service.GetMessagesByPanoramaID(req.PanoramaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取留言列表失败: " + err.Error(),
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
	userID := c.GetUint64("userID")
	if userID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "无效的用户ID",
		})
		return
	}

	messages, err := h.service.GetMessagesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": fmt.Sprintf("获取用户留言失败: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "获取用户留言成功",
		"data":    messages,
	})
}

// DeleteMessage 删除留言
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	var req dto.DeleteMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "无效的请求参数: " + err.Error(),
		})
		return
	}

	if err := h.service.DeleteMessage(req.MessageID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}
