package dto

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=4,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Code     string `json:"code" binding:"required,len=6"`
}

// RegisterResponse 用户注册响应
type RegisterResponse struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// LoginResponse 用户登录响应
type LoginResponse struct {
	UserID     uint   `json:"user_id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ExpireTime string `json:"expire_time"`
}

// UpdateUsernameRequest 更新用户名请求
type UpdateUsernameRequest struct {
	UserID   uint   `json:"userId" binding:"required"`
	Username string `json:"username" binding:"required,min=4,max=20"`
}

// UpdateUsernameResponse 更新用户名响应
type UpdateUsernameResponse struct {
	Username   string `json:"username"`
	UpdateTime string `json:"updateTime"`
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	UserID   uint   `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// ResetPasswordResponse 重置密码响应
type ResetPasswordResponse struct {
	UpdateTime string `json:"updateTime"`
}

// GetUserCreationTimeRequest 获取用户创建时间请求
type GetUserCreationTimeRequest struct {
	UserID uint `json:"userId" binding:"required"`
}

// GetUserCreationTimeResponse 获取用户创建时间响应
type GetUserCreationTimeResponse struct {
	CreateTime string `json:"createTime"`
}
