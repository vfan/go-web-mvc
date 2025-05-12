package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"mvc-demo/config"
)

var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	var err error
	appConfig := config.GetConfig()
	dbConfig := appConfig.DB

	// 构建DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Charset)

	// 创建自定义日志配置
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录未找到错误
			Colorful:                  true,        // 彩色打印
		},
	)

	// 打开数据库连接
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})

	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}

	log.Println("数据库连接成功")

	// 获取底层SQL数据库连接并设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("无法获取数据库连接池: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生存时间

	// 注意：我们不再自动迁移数据库结构
	// 在现代开发环境中，数据库结构变更应该由专门的数据库管理员或通过正式的数据库迁移工具进行
	// 这确保了数据库结构变更的可追踪性、安全性和一致性
}

// 以下是数据库迁移代码，仅供参考，不再自动执行
// 在实际生产环境中，应使用专门的数据库迁移工具，如 golang-migrate、atlas 等

// migrateModels 自动迁移模型到数据库 (仅供开发环境参考)
func migrateModels() {
	// 根据模型自动迁移数据库表结构
	err := DB.AutoMigrate(
		&User{},
		&University{},
		&Student{},
	)

	if err != nil {
		log.Fatalf("自动迁移数据库失败: %v", err)
	}

	log.Println("数据库迁移完成")
}
