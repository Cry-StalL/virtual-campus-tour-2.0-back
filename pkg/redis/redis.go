package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var client *redis.Client

// InitRedis 初始化Redis客户端
func InitRedis() error {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("连接Redis失败：%v", err)
	}

	return nil
}

// GetClient 获取Redis客户端
func GetClient() *redis.Client {
	return client
}
