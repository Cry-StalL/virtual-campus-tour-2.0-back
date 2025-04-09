package dto

// GetEmailCodeRequest 获取邮箱验证码请求
type GetEmailCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// GetEmailCodeResponse 获取邮箱验证码响应
type GetEmailCodeResponse struct {
	ExpireTime int `json:"expire_time"` // 验证码有效期（秒）
}
