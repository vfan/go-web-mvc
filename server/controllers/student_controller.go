package controllers

import (
	"errors"
	"mvc-demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// StudentController 学生控制器
type StudentController struct{}

// 创建学生
func (s *StudentController) Create(c *gin.Context) {
	var student models.Student

	// 绑定请求参数
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 创建学生
	if err := models.CreateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建学生失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "创建学生成功",
		"data":    student,
	})
}

// 获取学生详情
func (s *StudentController) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的学生ID",
		})
		return
	}

	student, err := models.GetStudentByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "学生不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取学生失败: " + err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

// 更新学生
func (s *StudentController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的学生ID",
		})
		return
	}

	// 检查学生是否存在
	existingStudent, err := models.GetStudentByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "学生不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取学生失败: " + err.Error(),
			})
		}
		return
	}

	// 绑定请求参数
	var updateData models.Student
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误: " + err.Error(),
		})
		return
	}

	// 更新学生属性
	updateData.ID = existingStudent.ID

	// 更新学生
	if err := models.UpdateStudent(&updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新学生失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新学生成功",
		"data":    updateData,
	})
}

// 删除学生
func (s *StudentController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的学生ID",
		})
		return
	}

	// 检查学生是否存在
	_, err = models.GetStudentByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "学生不存在",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "获取学生失败: " + err.Error(),
			})
		}
		return
	}

	// 删除学生
	if err := models.DeleteStudent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "删除学生失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "删除学生成功",
	})
}

// 获取学生列表
func (s *StudentController) List(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取筛选条件
	filters := make(map[string]interface{})

	// 添加可筛选的字段
	if name := c.Query("name"); name != "" {
		filters["name"] = name
	}

	if universityID := c.Query("university_id"); universityID != "" {
		uID, err := strconv.ParseUint(universityID, 10, 32)
		if err == nil {
			filters["university_id"] = uint(uID)
		}
	}

	if status := c.Query("status"); status != "" {
		studentStatus := models.StudentStatus(status)
		filters["status"] = studentStatus
	}

	// 获取学生列表
	students, total, err := models.GetStudents(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取学生列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":  students,
			"total": total,
			"page":  page,
			"size":  pageSize,
		},
	})
}
