package controllers

import (
	"errors"
	"mvc-demo/dao/model"
	"mvc-demo/models/dto"
	"mvc-demo/service"
	"mvc-demo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UniversityController 大学控制器
type UniversityController struct {
	universityService *service.UniversityService
}

// NewUniversityController 创建大学控制器
func NewUniversityController(universityService *service.UniversityService) *UniversityController {
	return &UniversityController{
		universityService: universityService,
	}
}

// Create 创建大学
func (u *UniversityController) Create(c *gin.Context) {
	var req dto.UniversityRequest

	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 创建大学对象
	university := &model.University{
		Name: req.Name,
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if exists {
		uid := int64(userID.(uint))
		university.CreatedBy = &uid
		university.UpdatedBy = &uid
	}

	// 创建大学
	if err := u.universityService.CreateUniversity(university); err != nil {
		utils.InternalError(c, "创建大学失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	response := &dto.UniversityResponse{
		ID:        university.ID,
		Name:      university.Name,
		CreatedAt: university.CreatedAt,
		UpdatedAt: university.UpdatedAt,
		CreatedBy: university.CreatedBy,
		UpdatedBy: university.UpdatedBy,
	}

	utils.SuccessWithMsg(c, "创建成功", response)
}

// Get 获取大学详情
func (u *UniversityController) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的大学ID")
		return
	}

	university, err := u.universityService.GetUniversityByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "大学不存在")
		} else {
			utils.InternalError(c, "获取大学失败: "+err.Error())
		}
		return
	}

	// 转换为响应DTO
	response := &dto.UniversityResponse{
		ID:        university.ID,
		Name:      university.Name,
		CreatedAt: university.CreatedAt,
		UpdatedAt: university.UpdatedAt,
		CreatedBy: university.CreatedBy,
		UpdatedBy: university.UpdatedBy,
	}

	utils.Success(c, response)
}

// Update 更新大学
func (u *UniversityController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的大学ID")
		return
	}

	// 检查大学是否存在
	existingUniversity, err := u.universityService.GetUniversityByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "大学不存在")
		} else {
			utils.InternalError(c, "获取大学失败: "+err.Error())
		}
		return
	}

	// 绑定请求参数
	var req dto.UniversityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 更新大学属性
	existingUniversity.Name = req.Name

	// 获取当前用户ID，设置更新者
	userID, exists := c.Get("user_id")
	if exists {
		uid := int64(userID.(uint))
		existingUniversity.UpdatedBy = &uid
	}

	// 更新大学
	if err := u.universityService.UpdateUniversity(existingUniversity); err != nil {
		utils.InternalError(c, "更新大学失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	response := &dto.UniversityResponse{
		ID:        existingUniversity.ID,
		Name:      existingUniversity.Name,
		CreatedAt: existingUniversity.CreatedAt,
		UpdatedAt: existingUniversity.UpdatedAt,
		CreatedBy: existingUniversity.CreatedBy,
		UpdatedBy: existingUniversity.UpdatedBy,
	}

	utils.SuccessWithMsg(c, "更新成功", response)
}

// Delete 删除大学
func (u *UniversityController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的大学ID")
		return
	}

	// 检查大学是否存在
	_, err = u.universityService.GetUniversityByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "大学不存在")
		} else {
			utils.InternalError(c, "获取大学失败: "+err.Error())
		}
		return
	}

	// 删除大学
	if err := u.universityService.DeleteUniversity(id); err != nil {
		utils.InternalError(c, "删除大学失败: "+err.Error())
		return
	}

	utils.SuccessWithMsg(c, "删除成功", nil)
}

// List 获取大学列表
func (u *UniversityController) List(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取大学列表
	universities, total, err := u.universityService.GetUniversityList(page, pageSize)
	if err != nil {
		utils.InternalError(c, "获取大学列表失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	var responseList []*dto.UniversityResponse
	for _, university := range universities {
		responseList = append(responseList, &dto.UniversityResponse{
			ID:        university.ID,
			Name:      university.Name,
			CreatedAt: university.CreatedAt,
			UpdatedAt: university.UpdatedAt,
			CreatedBy: university.CreatedBy,
			UpdatedBy: university.UpdatedBy,
		})
	}

	response := &dto.UniversityListResponse{
		List:  responseList,
		Total: total,
		Page:  page,
		Size:  pageSize,
	}

	utils.Success(c, response)
}

// All 获取所有大学（下拉列表使用）
func (u *UniversityController) All(c *gin.Context) {
	// 获取所有大学
	universities, err := u.universityService.GetAllUniversities()
	if err != nil {
		utils.InternalError(c, "获取大学列表失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	var responseList []*dto.UniversityResponse
	for _, university := range universities {
		responseList = append(responseList, &dto.UniversityResponse{
			ID:        university.ID,
			Name:      university.Name,
			CreatedAt: university.CreatedAt,
			UpdatedAt: university.UpdatedAt,
			CreatedBy: university.CreatedBy,
			UpdatedBy: university.UpdatedBy,
		})
	}

	utils.Success(c, responseList)
}
