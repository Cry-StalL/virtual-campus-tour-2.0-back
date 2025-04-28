package handler

import (
	"net/http"

	"virtual-campus-tour-2.0-back/internal/dto"
	"virtual-campus-tour-2.0-back/internal/service"

	"log"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register 处理用户注册请求
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    nil,
		})
		return
	}

	resp, err := h.service.Register(&req)
	if err != nil {
		code := 400
		message := err.Error()
		switch message {
		case "邮箱已注册":
			code = 1001
		case "用户名已存在":
			code = 1002
		case "验证码不正确":
			code = 1003
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
		"data":    resp,
	})
}

// Login 处理用户登录请求
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    nil,
		})
		return
	}

	resp, err := h.service.Login(&req)
	if err != nil {
		code := 400
		message := err.Error()
		switch message {
		case "该邮箱尚未注册，请先注册":
			code = 1001
		case "密码错误":
			code = 1002
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data":    resp,
	})
}

// GetEmailCode 处理获取邮箱验证码请求
func (h *UserHandler) GetEmailCode(c *gin.Context) {
	var req dto.GetEmailCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    nil,
		})
		return
	}

	resp, err := h.service.GetEmailCode(&req)
	if err != nil {
		code := 400
		message := err.Error()
		switch message {
		case "邮箱已被注册":
			code = 2001
		case "邮件发送失败":
			code = 2005
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证码发送成功",
		"data":    resp,
	})
}

// UpdateUsername 更新用户名
func (h *UserHandler) UpdateUsername(c *gin.Context) {
	var req dto.UpdateUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误: " + err.Error(),
			"data":    nil,
		})
		return
	}

	// 添加日志记录
	log.Printf("更新用户名请求: UserID=%d, Username=%s", req.UserID, req.Username)

	if req.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "用户ID不能为0",
			"data":    nil,
		})
		return
	}

	err := h.service.UpdateUsername(req.UserID, req.Username)
	if err != nil {
		code := 500
		message := err.Error()
		if message == "用户不存在" {
			code = 404
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "用户名更新成功",
		"data":    nil,
	})
}

// ResetPassword 重置密码
func (h *UserHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    nil,
		})
		return
	}

	err := h.service.UpdatePassword(req.UserID, req.Password)
	if err != nil {
		code := 500
		message := err.Error()
		if message == "用户不存在" {
			code = 404
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "密码重置成功",
		"data":    nil,
	})
}

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	var req dto.UpdatePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdatePassword(req.UserID, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "密码更新成功"})
}
