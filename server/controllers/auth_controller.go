package controllers

import (
	"mvc-demo/config"
	"mvc-demo/models"
	"mvc-demo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController 认证控制器
type AuthController struct{}

// Login 用户登录
func (a *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest

	// 请求参数绑定
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "无效的请求参数",
			Data:    nil,
		})
		return
	}

	// 验证用户是否存在且密码正确
	user, err := models.ValidateUserLogin(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			Code:    http.StatusUnauthorized,
			Message: "用户名或密码错误",
			Data:    nil,
		})
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: "生成令牌失败",
			Data:    nil,
		})
		return
	}

	// 更新用户最后登录时间
	models.UpdateUserLastLoginTime(user.ID)

	// 构建响应
	jwtConfig := config.GetJWTConfig()
	expiresIn := int(jwtConfig.TokenExpiry.Seconds())

	// 返回响应
	c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data: models.LoginResponse{
			Token:     token,
			TokenType: "Bearer",
			ExpiresIn: expiresIn,
		},
	})
}
