package middleware

import (
	"mvc-demo/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从Cookie中获取token
		tokenString, err := c.Cookie("token")

		// 如果Cookie中没有token，则尝试从Authorization头获取
		if err != nil {
			// 获取Authorization头
			auth := c.GetHeader("Authorization")
			if auth == "" {
				utils.Unauthorized(c, "未提供授权令牌")
				c.Abort()
				return
			}

			// 检查Bearer前缀
			parts := strings.SplitN(auth, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				utils.Unauthorized(c, "无效的授权格式")
				c.Abort()
				return
			}

			tokenString = parts[1]
		}

		// 验证令牌
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.Unauthorized(c, "无效的令牌: "+err.Error())
			c.Abort()
			return
		}

		// 将用户信息存储在上下文中
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// AdminAuth 管理员权限中间件
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 需要先经过JWTAuth中间件
		role, exists := c.Get("role")

		if !exists {
			utils.Unauthorized(c, "未授权")
			c.Abort()
			return
		}

		// 检查用户角色是否为管理员(role=1)
		if role.(int8) != 1 {
			utils.Forbidden(c, "需要管理员权限")
			c.Abort()
			return
		}

		c.Next()
	}
}
