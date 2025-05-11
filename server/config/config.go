package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig 应用配置
type AppConfig struct {
	ServerPort string
	Mode       string
}

// 初始化环境变量
func init() {
	// 加载.env文件
	err := godotenv.Load()
	if err != nil {
		log.Println("警告: 未找到.env文件或无法加载, 将使用默认配置")
	}
}

// GetConfig 获取应用配置
func GetConfig() *AppConfig {
	return &AppConfig{
		// 从环境变量中读取，如果不存在则使用默认值
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		Mode:       getEnv("GIN_MODE", "debug"),
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
