package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your-secret-key") // 实际应用中应该从配置文件读取

type Claims struct {
<<<<<<< HEAD
	UserID uint64 `json:"user_id"`
=======
	UserID uint   `json:"user_id"`
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT token
<<<<<<< HEAD
func GenerateToken(userID uint64, email string) (string, error) {
	now := time.Now()
	expireTime := now.Add(24 * time.Hour)

=======
func GenerateToken(userID uint, email string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
>>>>>>> 8c80ad534dc0a99b86c3b040525f0238dd8298a4
	claims := Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
