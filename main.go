package main

import (
	"log"

	"virtual-campus-tour-2.0-back/internal/handler"
	"virtual-campus-tour-2.0-back/internal/model"
	"virtual-campus-tour-2.0-back/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 使用viper加载配置
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	dbConfig := &database.Config{
		Driver:    viper.GetString("database.driver"),
		Host:      viper.GetString("database.host"),
		Port:      viper.GetInt("database.port"),
		Username:  viper.GetString("database.username"),
		Password:  viper.GetString("database.password"),
		DBName:    viper.GetString("database.dbname"),
		Charset:   viper.GetString("database.charset"),
		ParseTime: viper.GetBool("database.parseTime"),
		Loc:       viper.GetString("database.loc"),
	}

	// 1. 初始化数据库
	if err := database.InitDB(dbConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库表
	if err := database.GetDB().AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 2. 创建 Gin 引擎
	r := gin.Default()

	// 3. 初始化处理器
	userHandler := handler.NewUserHandler()

	// 4. 注册路由
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.Register)
		}
	}

	// 5. 启动服务器
	r.Run(":8080")
}
