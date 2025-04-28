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
<<<<<<< HEAD
	UserID    uint64 `json:"user_id"`
=======
	UserID    uint   `json:"user_id"`
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
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
<<<<<<< HEAD
	UserID     uint64 `json:"user_id"`
=======
	UserID     uint   `json:"user_id"`
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
	Username   string `json:"username"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ExpireTime string `json:"expire_time"`
}

// UpdateUsernameRequest 更新用户名请求
type UpdateUsernameRequest struct {
<<<<<<< HEAD
	UserID   uint64 `json:"userId" binding:"required"`
	Username string `json:"username" binding:"required,min=3,max=20"`
=======
	UserID   uint   `json:"userId" binding:"required"`
	Username string `json:"username" binding:"required,min=4,max=20"`
}

// UpdateUsernameResponse 更新用户名响应
type UpdateUsernameResponse struct {
	Username   string `json:"username"`
	UpdateTime string `json:"updateTime"`
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
<<<<<<< HEAD
	UserID   uint64 `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// UpdatePasswordRequest 更新密码请求
type UpdatePasswordRequest struct {
	UserID   uint64 `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// UserResponse 用户响应
type UserResponse struct {
	UserID    uint64 `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
=======
	UserID   uint   `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// ResetPasswordResponse 重置密码响应
type ResetPasswordResponse struct {
	UpdateTime string `json:"updateTime"`
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
}
