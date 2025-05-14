package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {

	fmt.Println("test env", os.Getenv("DB_HOST"))
	// 获取数据库配置信息
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "mvc_demo")
	dbCharset := getEnv("DB_CHARSET", "utf8mb4")

	// 调试信息：打印环境变量
	log.Printf("数据库配置信息: HOST=%s, PORT=%s, USER=%s, DB=%s",
		dbHost, dbPort, dbUser, dbName)
	log.Printf("原始环境变量: DB_HOST=%s", os.Getenv("DB_HOST"))

	// 构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset)

	// 打印构建的DSN (隐藏密码)
	dsnLog := fmt.Sprintf("%s:***@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbUser, dbHost, dbPort, dbName, dbCharset)
	log.Printf("数据库连接DSN: %s", dsnLog)

	// 配置日志记录器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// 连接数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("获取DB实例失败: %v", err)
	}

	// 设置空闲连接池中的最大连接数
	sqlDB.SetMaxIdleConns(10)

	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)

	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("数据库连接成功")
}

// 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
