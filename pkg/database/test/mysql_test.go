package test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"virtual-campus-tour-2.0-back/pkg/database"
)

func TestDatabaseConnection(t *testing.T) {
	// 加载配置文件
	viper.SetConfigFile("../../../config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("读取配置文件失败: %v", err)
	}

	// 从配置文件获取测试配置
	config := &database.Config{
		Driver:    viper.GetString("database.driver"),
		Host:      viper.GetString("database.host"),
		Port:      viper.GetInt("database.port"),
		Username:  viper.GetString("database.username"),
		Password:  viper.GetString("database.password"),
		DBName:    viper.GetString("test.database.dbname"),
		Charset:   viper.GetString("database.charset"),
		ParseTime: viper.GetBool("database.parseTime"),
		Loc:       viper.GetString("database.loc"),
	}

	// 注意：在实际项目中，应该使用如下方式从配置文件加载：
	// viper.SetConfigFile("../../config/config.yaml")
	// viper.ReadInConfig()
	// var dbConfig database.Config
	// viper.UnmarshalKey("database", &dbConfig)
	// dbConfig.DBName = "virtual_campus_tour_test" // 确保使用测试数据库

	// 测试数据库连接
	t.Run("测试数据库连接初始化", func(t *testing.T) {
		err := database.InitDB(config)
		assert.NoError(t, err, "数据库连接初始化失败")
		if err == nil {
			t.Log("数据库连接初始化成功")
		}

		// 验证全局DB实例不为空
		db := database.GetDB()
		assert.NotNil(t, db, "数据库连接实例为空")
		if db != nil {
			t.Log("数据库连接实例创建成功")
		}

		// 测试数据库连接是否可用
		sqlDB, err := db.DB()
		assert.NoError(t, err, "获取底层数据库连接失败")
		if err == nil {
			t.Log("获取底层数据库连接成功")
		}

		err = sqlDB.Ping()
		assert.NoError(t, err, "数据库连接测试失败")
		if err == nil {
			t.Log("数据库连接测试成功")
		}
	})

	// 测试数据库连接关闭
	t.Run("测试数据库连接关闭", func(t *testing.T) {
		err := database.CloseDB()
		assert.NoError(t, err, "关闭数据库连接失败")
		if err == nil {
			t.Log("数据库连接关闭成功")
		}
	})
}
