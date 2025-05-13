package controllers

import (
	"errors"
	"mvc-demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController 用户控制器
type UserController struct{}

// 创建用户
func (u *UserController) Create(c *gin.Context) {
	var req models.RegisterRequest

	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 创建用户对象
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	// 创建用户
	if err := models.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建用户失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建用户成功",
		"data":    user,
	})
}

// 获取用户详情
func (u *UserController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的用户ID",
		})
		return
	}

	user, err := models.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用户不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取用户失败: " + err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// 更新用户
func (u *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的用户ID",
		})
		return
	}

	// 检查用户是否存在
	existingUser, err := models.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用户不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取用户失败: " + err.Error(),
			})
		}
		return
	}

	// 绑定请求参数
	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 更新用户属性
	updateData.ID = existingUser.ID

	// 更新用户
	if err := models.UpdateUser(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新用户失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新用户成功",
		"data":    updateData,
	})
}

// 删除用户
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的用户ID",
		})
		return
	}

	// 检查用户是否存在
	_, err = models.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "用户不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取用户失败: " + err.Error(),
			})
		}
		return
	}

	// 删除用户
	if err := models.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除用户失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除用户成功",
	})
}

// 获取用户列表
func (u *UserController) List(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取用户列表
	users, total, err := models.GetUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取用户列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":  users,
			"total": total,
			"page":  page,
			"size":  pageSize,
		},
	})
}
