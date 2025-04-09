package utils

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

const (
	// 验证码有效期（秒）
	CodeExpireTime = 300
	// 同一邮箱发送间隔（秒）
	EmailSendInterval = 60
	// 同一IP发送次数限制
	IPSendLimit = 3
	// IP发送限制时间（秒）
	IPSendLimitTime = 60
)

// GenerateCode 生成6位数字验证码
func GenerateCode() string {
	b := make([]byte, 3)
	rand.Read(b)
	return fmt.Sprintf("%06d", int(b[0])*10000+int(b[1])*100+int(b[2])%100)
}

// StoreCode 存储验证码到Redis
func StoreCode(redisClient *redis.Client, email, code string) error {
	ctx := context.Background()
	key := fmt.Sprintf("email_code:%s", email)
	return redisClient.Set(ctx, key, code, CodeExpireTime*time.Second).Err()
}

// GetCode 从Redis获取验证码
func GetCode(redisClient *redis.Client, email string) (string, error) {
	ctx := context.Background()
	key := fmt.Sprintf("email_code:%s", email)
	return redisClient.Get(ctx, key).Result()
}

// CheckEmailSendInterval 检查邮箱发送间隔
func CheckEmailSendInterval(redisClient *redis.Client, email string) error {
	ctx := context.Background()
	key := fmt.Sprintf("email_send_time:%s", email)
	lastSendTime, err := redisClient.Get(ctx, key).Int64()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return err
	}
	if time.Now().Unix()-lastSendTime < EmailSendInterval {
		return fmt.Errorf("发送过于频繁，请稍后再试")
	}
	return nil
}

// UpdateEmailSendTime 更新邮箱发送时间
func UpdateEmailSendTime(redisClient *redis.Client, email string) error {
	ctx := context.Background()
	key := fmt.Sprintf("email_send_time:%s", email)
	return redisClient.Set(ctx, key, time.Now().Unix(), EmailSendInterval*time.Second).Err()
}

// CheckIPSendLimit 检查IP发送次数限制
func CheckIPSendLimit(redisClient *redis.Client, ip string) error {
	ctx := context.Background()
	key := fmt.Sprintf("ip_send_count:%s", ip)
	count, err := redisClient.Get(ctx, key).Int()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return err
	}
	if count >= IPSendLimit {
		return fmt.Errorf("IP发送次数过多，请稍后再试")
	}
	return nil
}

// UpdateIPSendCount 更新IP发送次数
func UpdateIPSendCount(redisClient *redis.Client, ip string) error {
	ctx := context.Background()
	key := fmt.Sprintf("ip_send_count:%s", ip)
	_, err := redisClient.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	return redisClient.Expire(ctx, key, IPSendLimitTime*time.Second).Err()
}
