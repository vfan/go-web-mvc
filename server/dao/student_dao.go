package dao

import (
	"mvc-demo/dao/model"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// StudentDAO 学生数据访问对象
type StudentDAO struct {
	DB *gorm.DB
}

// NewStudentDAO 创建学生DAO实例
func NewStudentDAO(db *gorm.DB) *StudentDAO {
	return &StudentDAO{DB: db}
}

// Create 创建学生
func (dao *StudentDAO) Create(student *model.Student) error {
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student.Password = string(hashedPassword)

	return dao.DB.Create(student).Error
}

// GetByID 根据ID获取学生
func (dao *StudentDAO) GetByID(id int64) (*model.Student, error) {
	var student model.Student
	err := dao.DB.Preload("University").First(&student, id).Error
	return &student, err
}

// GetByEmail 根据邮箱获取学生
func (dao *StudentDAO) GetByEmail(email string) (*model.Student, error) {
	var student model.Student
	err := dao.DB.Where("email = ?", email).First(&student).Error
	return &student, err
}

// Update 更新学生
func (dao *StudentDAO) Update(student *model.Student) error {
	return dao.DB.Save(student).Error
}

// Delete 删除学生
func (dao *StudentDAO) Delete(id int64) error {
	return dao.DB.Delete(&model.Student{}, id).Error
}

// GetList 获取学生列表（支持分页和筛选）
func (dao *StudentDAO) GetList(page, pageSize int, filters map[string]interface{}) ([]*model.Student, int64, error) {
	var students []*model.Student
	var total int64

	// 构建查询
	query := dao.DB.Model(&model.Student{})

	// 添加筛选条件
	for key, value := range filters {
		if value != nil && value != "" {
			// 支持模糊查询的字段
			if key == "name" {
				query = query.Where(key+" LIKE ?", "%"+value.(string)+"%")
			} else {
				query = query.Where(key+" = ?", value)
			}
		}
	}

	// 查询总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取数据列表
	offset := (page - 1) * pageSize
	err = query.Preload("University").Offset(offset).Limit(pageSize).Find(&students).Error
	return students, total, err
}

// ValidateLogin 验证学生登录
func (dao *StudentDAO) ValidateLogin(email, password string) (*model.Student, error) {
	// 查找学生
	student, err := dao.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return student, nil
}

// UpdateLastLoginTime 更新学生最后登录时间
func (dao *StudentDAO) UpdateLastLoginTime(studentID int64) error {
	now := time.Now()
	return dao.DB.Model(&model.Student{}).Where("id = ?", studentID).Update("last_login_time", &now).Error
}

// ResetPassword 重置学生密码
func (dao *StudentDAO) ResetPassword(email, newPassword string) error {
	student, err := dao.GetByEmail(email)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	student.Password = string(hashedPassword)
	return dao.DB.Save(student).Error
}
