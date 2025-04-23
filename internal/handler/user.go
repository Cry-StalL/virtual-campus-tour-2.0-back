package handler

import (
	"net/http"

	"virtual-campus-tour-2.0-back/internal/dto"
	"virtual-campus-tour-2.0-back/internal/service"
	"virtual-campus-tour-2.0-back/pkg/redis"
	"virtual-campus-tour-2.0-back/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
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

	resp, err := h.userService.Register(&req)
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
		"code":    0,
		"message": "success",
		"data":    resp,
	})
}

// UpdateUsername 处理更新用户名请求
func (h *UserHandler) UpdateUsername(c *gin.Context) {
	// 1. 获取并验证请求参数
	var req dto.UpdateUsernameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    nil,
		})
		return
	}

	// 2. 调用服务层处理请求
	resp, err := h.userService.UpdateUsername(&req)
	if err != nil {
		code := 400
		message := err.Error()
		switch message {
		case "用户不存在":
			code = 1001
		case "用户名已存在":
			code = 1002
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	// 3. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "用户名更新成功",
		"data":    resp,
	})
}

// ResetPassword 处理重置密码请求
func (h *UserHandler) ResetPassword(c *gin.Context) {
	// 1. 获取并验证请求参数
	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    nil,
		})
		return
	}

	// 2. 调用服务层处理请求
	resp, err := h.userService.ResetPassword(&req)
	if err != nil {
		code := 400
		message := err.Error()
		switch message {
		case "用户不存在":
			code = 1001
		case "密码加密失败":
			code = 1002
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": message,
			"data":    nil,
		})
		return
	}

	// 3. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "密码重置成功",
		"data":    resp,
	})
}

// Login 处理用户登录请求
func (h *UserHandler) Login(c *gin.Context) {
	// 1. 获取并验证请求参数
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1003,
			"message": "邮箱格式不正确",
			"data":    nil,
		})
		return
	}

	// 2. 检查登录频率限制
	ip := c.ClientIP()
	if utils.IsLoginTooFrequent(ip) {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"code":    1004,
			"message": "登录尝试次数过多，请稍后再试",
			"data":    nil,
		})
		return
	}

	// 3. 调用服务层处理登录
	resp, err := h.userService.Login(&req)
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

	// 4. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    resp,
	})
}

// GetEmailCode 处理获取邮箱验证码请求
func (h *UserHandler) GetEmailCode(c *gin.Context) {
	// 1. 获取并验证请求参数
	var req dto.GetEmailCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    2002,
			"message": "邮箱格式不正确",
			"data":    nil,
		})
		return
	}

	// 2. 检查发送频率限制
	ip := c.ClientIP()
	if err := utils.CheckIPSendLimit(redis.GetClient(), ip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    2004,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if err := utils.CheckEmailSendInterval(redis.GetClient(), req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    2003,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 3. 调用服务层处理请求
	resp, err := h.userService.GetEmailCode(&req)
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

	// 4. 更新发送记录
	utils.UpdateEmailSendTime(redis.GetClient(), req.Email)
	utils.UpdateIPSendCount(redis.GetClient(), ip)

	// 5. 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    resp,
	})
}
