package controllers

import (
	"errors"
	"mvc-demo/dao/model"
	"mvc-demo/models/dto"
	"mvc-demo/service"
	"mvc-demo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建用户控制器
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Create 创建用户
func (u *UserController) Create(c *gin.Context) {
	var req dto.RegisterRequest

	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 创建用户对象
	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
		Username: req.Username,
	}

	// 创建用户
	if err := u.userService.CreateUser(user); err != nil {
		utils.InternalError(c, "创建用户失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	response := &dto.UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		Role:          user.Role,
		Status:        user.Status,
		LastLoginTime: user.LastLoginTime,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	utils.SuccessWithMsg(c, "创建用户成功", response)
}

// Get 获取用户详情
func (u *UserController) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的用户ID")
		return
	}

	user, err := u.userService.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "用户不存在")
		} else {
			utils.InternalError(c, "获取用户失败: "+err.Error())
		}
		return
	}

	// 转换为响应DTO
	response := &dto.UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		Role:          user.Role,
		Status:        user.Status,
		LastLoginTime: user.LastLoginTime,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	utils.Success(c, response)
}

// Update 更新用户
func (u *UserController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的用户ID")
		return
	}

	// 检查用户是否存在
	existingUser, err := u.userService.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "用户不存在")
		} else {
			utils.InternalError(c, "获取用户失败: "+err.Error())
		}
		return
	}

	// 绑定请求参数
	var updateData dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 更新用户属性
	existingUser.Email = updateData.Email
	existingUser.Username = updateData.Username
	if updateData.Role > 0 {
		existingUser.Role = updateData.Role
	}
	if updateData.Status > 0 {
		existingUser.Status = updateData.Status
	}

	// 如果有提供密码，则更新密码
	if updateData.Password != "" {
		// 加密新密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.InternalError(c, "密码加密失败: "+err.Error())
			return
		}
		existingUser.Password = string(hashedPassword)
	}

	// 更新用户
	if err := u.userService.UpdateUser(existingUser); err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	// 转换为响应DTO
	response := &dto.UserResponse{
		ID:            existingUser.ID,
		Email:         existingUser.Email,
		Username:      existingUser.Username,
		Role:          existingUser.Role,
		Status:        existingUser.Status,
		LastLoginTime: existingUser.LastLoginTime,
		CreatedAt:     existingUser.CreatedAt,
		UpdatedAt:     existingUser.UpdatedAt,
	}

	utils.SuccessWithMsg(c, "更新用户成功", response)
}

// Delete 删除用户
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的用户ID")
		return
	}

	// 检查用户是否存在
	_, err = u.userService.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "用户不存在")
		} else {
			utils.InternalError(c, "获取用户失败: "+err.Error())
		}
		return
	}

	// 删除用户
	if err := u.userService.DeleteUser(id); err != nil {
		utils.InternalError(c, err.Error())
		return
	}

	utils.SuccessWithMsg(c, "删除用户成功", nil)
}

// List 获取用户列表
func (u *UserController) List(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取用户列表
	users, total, err := u.userService.GetUserList(page, pageSize)
	if err != nil {
		utils.InternalError(c, "获取用户列表失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	var responseList []*dto.UserResponse
	for _, user := range users {
		responseList = append(responseList, &dto.UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			Username:      user.Username,
			Role:          user.Role,
			Status:        user.Status,
			LastLoginTime: user.LastLoginTime,
			CreatedAt:     user.CreatedAt,
			UpdatedAt:     user.UpdatedAt,
		})
	}

	response := &dto.UserListResponse{
		List:  responseList,
		Total: total,
		Page:  page,
		Size:  pageSize,
	}

	utils.Success(c, response)
}
