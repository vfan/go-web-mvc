package controllers

import (
	"errors"
	"mvc-demo/models"
	"mvc-demo/utils"
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
		utils.ParamError(c, "参数错误: "+err.Error())
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
		utils.InternalError(c, "创建用户失败: "+err.Error())
		return
	}

	utils.SuccessWithMsg(c, "创建用户成功", user)
}

// 获取用户详情
func (u *UserController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ParamError(c, "无效的用户ID")
		return
	}

	user, err := models.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "用户不存在")
		} else {
			utils.InternalError(c, "获取用户失败: "+err.Error())
		}
		return
	}

	utils.Success(c, user)
}

// 更新用户
func (u *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ParamError(c, "无效的用户ID")
		return
	}

	// 检查用户是否存在
	existingUser, err := models.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "用户不存在")
		} else {
			utils.InternalError(c, "获取用户失败: "+err.Error())
		}
		return
	}

	// 绑定请求参数
	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 更新用户属性
	updateData.ID = existingUser.ID

	// 更新用户
	if err := models.UpdateUser(&updateData); err != nil {
		utils.InternalError(c, "更新用户失败: "+err.Error())
		return
	}

	utils.SuccessWithMsg(c, "更新用户成功", updateData)
}

// 删除用户
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ParamError(c, "无效的用户ID")
		return
	}

	// 检查用户是否存在
	_, err = models.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "用户不存在")
		} else {
			utils.InternalError(c, "获取用户失败: "+err.Error())
		}
		return
	}

	// 删除用户
	if err := models.DeleteUser(uint(id)); err != nil {
		utils.InternalError(c, "删除用户失败: "+err.Error())
		return
	}

	utils.SuccessWithMsg(c, "删除用户成功", nil)
}

// 获取用户列表
func (u *UserController) List(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取用户列表
	users, total, err := models.GetUsers(page, pageSize)
	if err != nil {
		utils.InternalError(c, "获取用户列表失败: "+err.Error())
		return
	}

	utils.Success(c, gin.H{
		"list":  users,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}
