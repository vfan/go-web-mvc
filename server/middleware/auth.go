package middleware

import (
	"mvc-demo/models"
	"mvc-demo/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:    http.StatusUnauthorized,
				Message: "未提供授权令牌",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(auth, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:    http.StatusUnauthorized,
				Message: "无效的授权格式",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 验证令牌
		tokenString := parts[1]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:    http.StatusUnauthorized,
				Message: "无效的令牌: " + err.Error(),
				Data:    nil,
			})
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
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:    http.StatusUnauthorized,
				Message: "未授权",
				Data:    nil,
			})
			c.Abort()
			return
		}

		// 检查用户角色是否为管理员(role=1)
		if role.(int8) != 1 {
			c.JSON(http.StatusForbidden, models.Response{
				Code:    http.StatusForbidden,
				Message: "需要管理员权限",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
