package utils

import (
	"errors"
	"mvc-demo/config"
	"mvc-demo/dao/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成JWT令牌
func GenerateToken(user *model.User) (string, error) {
	jwtConfig := config.GetJWTConfig()

	// 创建Claims
	claims := &config.JWTClaims{
		UserID: uint(user.ID),
		Email:  user.Email,
		Role:   int8(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jwtConfig.Issuer,
		},
	}

	// 生成令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名
	tokenString, err := token.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证JWT令牌
func ValidateToken(tokenString string) (*config.JWTClaims, error) {
	jwtConfig := config.GetJWTConfig()

	// 解析令牌
	token, err := jwt.ParseWithClaims(
		tokenString,
		&config.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 确保使用正确的签名方法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("无效的签名方法")
			}
			return []byte(jwtConfig.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	// 验证令牌并提取声明
	if claims, ok := token.Claims.(*config.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌")
}
