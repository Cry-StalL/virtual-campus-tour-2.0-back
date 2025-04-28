package repository

import (
	"virtual-campus-tour-2.0-back/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建新用户
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// GetByEmail 通过邮箱获取用户
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID 通过ID获取用户
func (r *UserRepository) GetByID(id uint64) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUsername 更新用户名
func (r *UserRepository) UpdateUsername(id uint64, username string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("username", username).Error
}

// UpdatePassword 更新密码
func (r *UserRepository) UpdatePassword(id uint64, password string) error {
	return r.db.Model(&model.User{}).Where("id = ?", id).Update("password", password).Error
}

// CheckUsernameExists 检查用户名是否已存在
func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).
		Where("username = ?", username).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckEmailExists 检查邮箱是否已存在
func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).
		Where("email = ?", email).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
