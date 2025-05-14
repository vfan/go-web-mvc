package service

import (
	"mvc-demo/dao"
	"mvc-demo/dao/model"
	"time"
)

// StudentService 学生服务
type StudentService struct {
	studentDAO *dao.StudentDAO
}

// NewStudentService 创建学生服务实例
func NewStudentService(studentDAO *dao.StudentDAO) *StudentService {
	return &StudentService{
		studentDAO: studentDAO,
	}
}

// CreateStudent 创建学生
func (s *StudentService) CreateStudent(student *model.Student) error {
	return s.studentDAO.Create(student)
}

// GetStudentByID 根据ID获取学生
func (s *StudentService) GetStudentByID(id int64) (*model.Student, error) {
	return s.studentDAO.GetByID(id)
}

// GetStudentByEmail 根据邮箱获取学生
func (s *StudentService) GetStudentByEmail(email string) (*model.Student, error) {
	return s.studentDAO.GetByEmail(email)
}

// UpdateStudent 更新学生信息
func (s *StudentService) UpdateStudent(student *model.Student) error {
	return s.studentDAO.Update(student)
}

// DeleteStudent 删除学生
func (s *StudentService) DeleteStudent(id int64) error {
	return s.studentDAO.Delete(id)
}

// GetStudentList 获取学生列表（支持分页和筛选）
func (s *StudentService) GetStudentList(page, pageSize int, filters map[string]interface{}) ([]*model.Student, int64, error) {
	return s.studentDAO.GetList(page, pageSize, filters)
}

// Login 学生登录
func (s *StudentService) Login(email, password string) (*model.Student, error) {
	// 验证登录信息
	student, err := s.studentDAO.ValidateLogin(email, password)
	if err != nil {
		return nil, err
	}

	// 更新最后登录时间
	now := time.Now()
	student.LastLoginTime = &now
	err = s.studentDAO.UpdateLastLoginTime(student.ID)
	if err != nil {
		// 只记录错误，不影响登录
		// log.Printf("更新登录时间失败: %v", err)
	}

	return student, nil
}

// ResetPassword 重置密码
func (s *StudentService) ResetPassword(email, newPassword string) error {
	return s.studentDAO.ResetPassword(email, newPassword)
}
