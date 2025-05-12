package config

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey     string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	Issuer        string
}

// JWT 声明结构体
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   int8   `json:"role"`
	jwt.RegisteredClaims
}

// GetJWTConfig 获取JWT配置
func GetJWTConfig() *JWTConfig {
	return &JWTConfig{
		// 从环境变量中读取，如果不存在则使用默认值
		SecretKey:     getEnv("JWT_SECRET_KEY", "your-secret-key-change-in-production"),
		TokenExpiry:   time.Duration(getEnvInt("JWT_TOKEN_EXPIRY", 24)) * time.Hour,    // 默认24小时
		RefreshExpiry: time.Duration(getEnvInt("JWT_REFRESH_EXPIRY", 168)) * time.Hour, // 默认7天
		Issuer:        getEnv("JWT_ISSUER", "student-management-system"),
	}
}

// getEnvInt 获取整型环境变量，如果不存在则返回默认值
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intValue := 0
	_, err := fmt.Sscanf(value, "%d", &intValue)
	if err != nil {
		return defaultValue
	}

	return intValue
}
