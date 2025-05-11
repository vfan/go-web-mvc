package controllers

import (
	"mvc-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController 处理首页相关的请求
type HomeController struct{}

// NewHomeController 创建一个新的HomeController实例
func NewHomeController() *HomeController {
	return &HomeController{}
}

// Index 处理首页请求
func (h *HomeController) Index(c *gin.Context) {
	response := models.NewResponse(200, "Hello World!", nil)
	c.JSON(http.StatusOK, response)
}

// Hello 处理问候请求
func (h *HomeController) Hello(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	response := models.NewResponse(200, "Hello "+name+"!", nil)
	c.JSON(http.StatusOK, response)
}
