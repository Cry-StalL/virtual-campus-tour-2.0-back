package main

import (
	"log"
	"time"

	"virtual-campus-tour-2.0-back/internal/handler"
	"virtual-campus-tour-2.0-back/internal/model"
	"virtual-campus-tour-2.0-back/internal/repository"
	"virtual-campus-tour-2.0-back/internal/service"
	"virtual-campus-tour-2.0-back/pkg/database"
	"virtual-campus-tour-2.0-back/pkg/redis"
	"virtual-campus-tour-2.0-back/pkg/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 使用viper加载配置
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// 1. 初始化数据库
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

	if err := database.InitDB(dbConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库表
	db := database.GetDB()

	// 自动迁移表结构
	if err := db.AutoMigrate(&model.User{}, &model.Message{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 2. 初始化Redis
	if err := redis.InitRedis(); err != nil {
		log.Printf("警告：Redis连接失败: %v，程序将继续运行", err)
	} else {
		log.Println("Redis连接成功")
	}

	// 3. 初始化邮件配置
	utils.InitEmailConfig(
		viper.GetString("email.host"),
		viper.GetInt("email.port"),
		viper.GetString("email.username"),
		viper.GetString("email.password"),
		viper.GetString("email.from"),
	)

	// 4. 创建 Gin 引擎
	r := gin.Default()

	// 添加CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源，实际生产环境建议指定具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 5. 初始化依赖
	db = database.GetDB()

	// 初始化用户相关依赖
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 初始化消息相关依赖
	messageRepo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepo)
	messageHandler := handler.NewMessageHandler(messageService)

	// 6. 注册路由
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			// 用户认证相关
			users.POST("/email-code", userHandler.GetEmailCode)
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)

			// 用户信息修改相关
			users.POST("/updateUsername", userHandler.UpdateUsername)
			users.POST("/resetPassword", userHandler.ResetPassword)

			// 消息相关
			users.POST("/messages", messageHandler.CreateMessage)
			users.GET("/messages", messageHandler.GetMessages)
			users.POST("/getUserMessages", messageHandler.GetUserMessages)
		}
	}

	// 7. 启动服务器
	r.Run(":8080")
}
