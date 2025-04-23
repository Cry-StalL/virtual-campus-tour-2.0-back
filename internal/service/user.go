package service

import (
	"errors"
	"time"

	"virtual-campus-tour-2.0-back/internal/dto"
	"virtual-campus-tour-2.0-back/internal/model"
	"virtual-campus-tour-2.0-back/internal/repository"
	"virtual-campus-tour-2.0-back/pkg/redis"
	"virtual-campus-tour-2.0-back/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetEmailCode 获取邮箱验证码
func (s *UserService) GetEmailCode(req *dto.GetEmailCodeRequest) (*dto.GetEmailCodeResponse, error) {
	// 检查邮箱是否已注册
	exists, err := s.repo.CheckEmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("邮箱已被注册")
	}

	// 生成验证码
	code := utils.GenerateCode()

	// 存储验证码到Redis
	if err := utils.StoreCode(redis.GetClient(), req.Email, code); err != nil {
		return nil, errors.New("存储验证码失败")
	}

	// 发送验证码邮件
	if err := utils.SendVerificationCode(req.Email, code); err != nil {
		return nil, errors.New("邮件发送失败")
	}

	return &dto.GetEmailCodeResponse{
		ExpireTime: utils.CodeExpireTime,
	}, nil
}

// Register 用户注册
func (s *UserService) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// 检查邮箱是否已注册
	exists, err := s.repo.CheckEmailExists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("邮箱已注册")
	}

	// 检查用户名是否已存在
	exists, err = s.repo.CheckUsernameExists(req.Username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 验证邮箱验证码
	if err := utils.VerifyCode(redis.GetClient(), req.Email, req.Code); err != nil {
		return nil, errors.New("验证码不正确")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
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
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
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

// CreateUser 创建用户
func (s *UserService) CreateUser(username, email, password string) (*model.User, error) {
	// 检查用户名是否已存在
	exists, err := s.repo.CheckUsernameExists(username)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUsername 更新用户名
func (s *UserService) UpdateUsername(userID, newUsername string) error {
	// 检查用户是否存在
	_, err := s.repo.GetByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 检查新用户名是否已存在
	exists, err := s.repo.CheckUsernameExists(newUsername)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("用户名已被占用")
	}

	// 更新用户名
	return s.repo.UpdateUsername(userID, newUsername)
}

// UpdatePassword 更新密码
func (s *UserService) UpdatePassword(userID, newPassword string) error {
	// 检查用户是否存在
	_, err := s.repo.GetByID(userID)
	if err != nil {
		return errors.New("用户不存在")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	return s.repo.UpdatePassword(userID, string(hashedPassword))
}

// VerifyUser 验证用户
func (s *UserService) VerifyUser(email, password string) (*model.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return user, nil
}
