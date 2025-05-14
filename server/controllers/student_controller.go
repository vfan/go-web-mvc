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

// StudentController 学生控制器
type StudentController struct {
	studentService    *service.StudentService
	universityService *service.UniversityService
}

// NewStudentController 创建学生控制器
func NewStudentController(studentService *service.StudentService, universityService *service.UniversityService) *StudentController {
	return &StudentController{
		studentService:    studentService,
		universityService: universityService,
	}
}

// Create 创建学生
func (s *StudentController) Create(c *gin.Context) {
	var req dto.StudentCreateRequest

	// 绑定请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 创建学生对象
	student := &model.Student{
		Name:           req.Name,
		Password:       req.Password,
		Email:          req.Email,
		Gender:         req.Gender,
		Birthday:       req.Birthday,
		Phone:          req.Phone,
		UniversityID:   req.UniversityID,
		Major:          req.Major,
		Education:      req.Education,
		GraduationYear: req.GraduationYear,
		Status:         req.Status,
		Remarks:        req.Remarks,
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if exists {
		uid := int64(userID.(uint))
		student.CreatedBy = &uid
		student.UpdatedBy = &uid
	}

	// 创建学生
	if err := s.studentService.CreateStudent(student); err != nil {
		utils.InternalError(c, "创建学生失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	response := s.convertToStudentResponse(student)
	utils.SuccessWithMsg(c, "创建学生成功", response)
}

// Get 获取学生详情
func (s *StudentController) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的学生ID")
		return
	}

	student, err := s.studentService.GetStudentByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "学生不存在")
		} else {
			utils.InternalError(c, "获取学生失败: "+err.Error())
		}
		return
	}

	// 转换为响应DTO
	response := s.convertToStudentResponse(student)
	utils.Success(c, response)
}

// Update 更新学生
func (s *StudentController) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的学生ID")
		return
	}

	// 检查学生是否存在
	existingStudent, err := s.studentService.GetStudentByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "学生不存在")
		} else {
			utils.InternalError(c, "获取学生失败: "+err.Error())
		}
		return
	}

	// 绑定请求参数
	var req dto.StudentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 更新学生属性
	if req.Name != "" {
		existingStudent.Name = req.Name
	}
	if req.Gender != nil {
		existingStudent.Gender = req.Gender
	}
	if req.Birthday != nil {
		existingStudent.Birthday = req.Birthday
	}
	if req.Phone != nil {
		existingStudent.Phone = req.Phone
	}
	if req.UniversityID != nil {
		existingStudent.UniversityID = req.UniversityID
	}
	if req.Major != nil {
		existingStudent.Major = req.Major
	}
	if req.Education != nil {
		existingStudent.Education = req.Education
	}
	if req.GraduationYear != nil {
		existingStudent.GraduationYear = req.GraduationYear
	}
	if req.Status != nil {
		existingStudent.Status = req.Status
	}
	if req.Remarks != nil {
		existingStudent.Remarks = req.Remarks
	}
	if req.Avatar != nil {
		existingStudent.Avatar = req.Avatar
	}

	// 获取当前用户ID
	userID, exists := c.Get("user_id")
	if exists {
		uid := int64(userID.(uint))
		existingStudent.UpdatedBy = &uid
	}

	// 更新学生
	if err := s.studentService.UpdateStudent(existingStudent); err != nil {
		utils.InternalError(c, "更新学生失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	response := s.convertToStudentResponse(existingStudent)
	utils.SuccessWithMsg(c, "更新学生成功", response)
}

// Delete 删除学生
func (s *StudentController) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParamError(c, "无效的学生ID")
		return
	}

	// 检查学生是否存在
	_, err = s.studentService.GetStudentByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.NotFound(c, "学生不存在")
		} else {
			utils.InternalError(c, "获取学生失败: "+err.Error())
		}
		return
	}

	// 删除学生
	if err := s.studentService.DeleteStudent(id); err != nil {
		utils.InternalError(c, "删除学生失败: "+err.Error())
		return
	}

	utils.SuccessWithMsg(c, "删除学生成功", nil)
}

// List 获取学生列表
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
		uID, err := strconv.ParseInt(universityID, 10, 64)
		if err == nil {
			filters["university_id"] = uID
		}
	}

	if status := c.Query("status"); status != "" {
		filters["status"] = status
	}

	// 获取学生列表
	students, total, err := s.studentService.GetStudentList(page, pageSize, filters)
	if err != nil {
		utils.InternalError(c, "获取学生列表失败: "+err.Error())
		return
	}

	// 转换为响应DTO
	var responseList []*dto.StudentResponse
	for _, student := range students {
		responseList = append(responseList, s.convertToStudentResponse(student))
	}

	response := &dto.StudentListResponse{
		List:  responseList,
		Total: total,
		Page:  page,
		Size:  pageSize,
	}

	utils.Success(c, response)
}

// 转换为学生响应DTO
func (s *StudentController) convertToStudentResponse(student *model.Student) *dto.StudentResponse {
	response := &dto.StudentResponse{
		ID:             student.ID,
		Name:           student.Name,
		Email:          student.Email,
		Gender:         student.Gender,
		Birthday:       student.Birthday,
		Phone:          student.Phone,
		UniversityID:   student.UniversityID,
		Major:          student.Major,
		Education:      student.Education,
		GraduationYear: student.GraduationYear,
		Status:         student.Status,
		Remarks:        student.Remarks,
		Avatar:         student.Avatar,
		LastLoginTime:  student.LastLoginTime,
		CreatedAt:      student.CreatedAt,
		UpdatedAt:      student.UpdatedAt,
		CreatedBy:      student.CreatedBy,
		UpdatedBy:      student.UpdatedBy,
	}

	// 如果有大学信息，添加大学响应
	if student.UniversityID != nil && *student.UniversityID > 0 {
		university, err := s.universityService.GetUniversityByID(*student.UniversityID)
		if err == nil {
			response.University = &dto.UniversityResponse{
				ID:        university.ID,
				Name:      university.Name,
				CreatedAt: university.CreatedAt,
				UpdatedAt: university.UpdatedAt,
				CreatedBy: university.CreatedBy,
				UpdatedBy: university.UpdatedBy,
			}
		}
	}

	return response
}
