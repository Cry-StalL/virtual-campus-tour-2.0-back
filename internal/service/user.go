package service

import (
	"errors"
	"time"

	"virtual-campus-tour-2.0-back/internal/dto"
	"virtual-campus-tour-2.0-back/internal/model"
	"virtual-campus-tour-2.0-back/pkg/database"
	"virtual-campus-tour-2.0-back/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// Register 用户注册
func (s *UserService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// 检查邮箱是否已注册
	var existingUser model.User
	if err := database.GetDB().Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("邮箱已注册")
	}

	// 检查用户名是否已存在
	if err := database.GetDB().Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	}

	// TODO: 验证邮箱验证码
	// if !verifyEmailCode(req.Email, req.Code) {
	// 	return nil, errors.New("验证码不正确")
	// }

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := database.GetDB().Create(&user).Error; err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		UserID:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

// Login 用户登录
func (s *UserService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. 验证邮箱格式（通过gin的binding标签已经验证）

	// 2. 查询用户是否存在
	var user model.User
	if err := database.GetDB().Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, errors.New("该邮箱尚未注册，请先注册")
	}

	// 3. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}

	// 4. 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	// 5. 设置token过期时间（24小时后）
	expireTime := time.Now().Add(24 * time.Hour)

	return &dto.LoginResponse{
		UserID:     user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Token:      token,
		ExpireTime: expireTime.Format(time.RFC3339),
	}, nil
}
