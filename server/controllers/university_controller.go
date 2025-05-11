package controllers

import (
	"errors"
	"mvc-demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UniversityController 大学控制器
type UniversityController struct{}

// 创建大学
func (u *UniversityController) Create(c *gin.Context) {
	var university models.University

	// 绑定请求参数
	if err := c.ShouldBindJSON(&university); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 创建大学
	if err := models.CreateUniversity(&university); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建大学失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建大学成功",
		"data":    university,
	})
}

// 获取大学详情
func (u *UniversityController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的大学ID",
		})
		return
	}

	university, err := models.GetUniversityByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "大学不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取大学失败: " + err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": university,
	})
}

// 更新大学
func (u *UniversityController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的大学ID",
		})
		return
	}

	// 检查大学是否存在
	existingUniversity, err := models.GetUniversityByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "大学不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取大学失败: " + err.Error(),
			})
		}
		return
	}

	// 绑定请求参数
	var updateData models.University
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 更新大学属性
	updateData.ID = existingUniversity.ID

	// 更新大学
	if err := models.UpdateUniversity(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新大学失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新大学成功",
		"data":    updateData,
	})
}

// 删除大学
func (u *UniversityController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的大学ID",
		})
		return
	}

	// 检查大学是否存在
	_, err = models.GetUniversityByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "大学不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取大学失败: " + err.Error(),
			})
		}
		return
	}

	// 删除大学
	if err := models.DeleteUniversity(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除大学失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除大学成功",
	})
}

// 获取大学列表
func (u *UniversityController) List(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取大学列表
	universities, total, err := models.GetUniversities(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取大学列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":  universities,
			"total": total,
			"page":  page,
			"size":  pageSize,
		},
	})
}

// 获取所有大学（下拉列表使用）
func (u *UniversityController) All(c *gin.Context) {
	// 获取所有大学
	universities, err := models.GetAllUniversities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取大学列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": universities,
	})
}
