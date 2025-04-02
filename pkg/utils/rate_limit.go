package utils

import (
	"sync"
	"time"
)

type LoginAttempt struct {
	Count    int
	FirstTry time.Time
	LastTry  time.Time
}

var (
	loginAttempts = make(map[string]*LoginAttempt)
	mutex         sync.RWMutex
)

// IsLoginTooFrequent 检查登录频率是否过高
func IsLoginTooFrequent(ip string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	now := time.Now()
	attempt, exists := loginAttempts[ip]

	if !exists {
		loginAttempts[ip] = &LoginAttempt{
			Count:    1,
			FirstTry: now,
			LastTry:  now,
		}
		return false
	}

	// 重置10分钟前的记录
	if now.Sub(attempt.FirstTry) > 10*time.Minute {
		attempt.Count = 1
		attempt.FirstTry = now
		attempt.LastTry = now
		return false
	}

	// 检查是否超过限制（5次/10分钟）
	if attempt.Count >= 5 {
		return true
	}

	attempt.Count++
	attempt.LastTry = now
	return false
}
