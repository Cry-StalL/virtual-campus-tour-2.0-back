package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
<<<<<<< HEAD
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // 密码不返回给前端
	Email     string    `json:"email" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
=======
	ID        uint           `gorm:"primarykey" json:"user_id"`
	Username  string         `gorm:"size:20;not null;unique" json:"username"`
	Email     string         `gorm:"size:100;not null;unique" json:"email"`
	Password  string         `gorm:"size:100;not null" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
}
