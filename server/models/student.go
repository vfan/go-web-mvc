package models

import (
	"time"

	"gorm.io/gorm"
)

// 教育程度枚举
type Education string

const (
	EducationJuniorCollege Education = "专科"
	EducationBachelor      Education = "本科"
	EducationMaster        Education = "硕士"
	EducationDoctor        Education = "博士"
)

// 学生状态枚举
type StudentStatus string

const (
	StatusStudying  StudentStatus = "在读"
	StatusSuspended StudentStatus = "休学"
	StatusDropped   StudentStatus = "退学"
	StatusGraduated StudentStatus = "毕业"
)

// Student 学生模型（对应students表）
type Student struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `gorm:"size:50;not null;index" json:"name"`
	Password       string         `gorm:"size:255;not null" json:"-"` // 不返回密码
	Email          string         `gorm:"uniqueIndex;size:100;not null" json:"email"`
	LastLoginTime  *time.Time     `gorm:"default:null" json:"last_login_time"`
	Gender         *int8          `gorm:"default:1" json:"gender"` // 1:男,2:女,3:其他
	Birthday       *time.Time     `json:"birthday"`
	Phone          *string        `gorm:"size:20" json:"phone"`
	ResumePath     *string        `gorm:"size:255" json:"resume_path"`
	UniversityID   *uint          `gorm:"index" json:"university_id"`
	University     *University    `gorm:"foreignKey:UniversityID" json:"university,omitempty"`
	Major          *string        `gorm:"size:100" json:"major"`
	Education      *Education     `gorm:"type:enum('专科','本科','硕士','博士');default:'本科'" json:"education"`
	GraduationYear *int           `json:"graduation_year"`
	Status         *StudentStatus `gorm:"type:enum('在读','休学','退学','毕业');default:'在读'" json:"status"`
	Remarks        *string        `gorm:"type:text" json:"remarks"`
	Avatar         *string        `gorm:"size:255" json:"avatar"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	CreatedBy      *uint          `json:"created_by"`
	UpdatedBy      *uint          `json:"updated_by"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (Student) TableName() string {
	return "students"
}

// 创建学生
func CreateStudent(student *Student) error {
	return DB.Create(student).Error
}

// 根据ID获取学生
func GetStudentByID(id uint) (*Student, error) {
	var student Student
	err := DB.Preload("University").First(&student, id).Error
	return &student, err
}

// 根据邮箱获取学生
func GetStudentByEmail(email string) (*Student, error) {
	var student Student
	err := DB.Where("email = ?", email).First(&student).Error
	return &student, err
}

// 更新学生
func UpdateStudent(student *Student) error {
	return DB.Save(student).Error
}

// 删除学生
func DeleteStudent(id uint) error {
	return DB.Delete(&Student{}, id).Error
}

// 获取学生列表（支持分页和筛选）
func GetStudents(page, pageSize int, filters map[string]interface{}) ([]Student, int64, error) {
	var students []Student
	var total int64

	// 构建查询
	query := DB.Model(&Student{})

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
