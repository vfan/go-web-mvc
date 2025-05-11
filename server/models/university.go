package models

import (
	"time"

	"gorm.io/gorm"
)

// University 大学模型（对应universities表）
type University struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"uniqueIndex;size:100;not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedBy *uint          `json:"created_by"`
	UpdatedBy *uint          `json:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (University) TableName() string {
	return "universities"
}

// 创建大学
func CreateUniversity(university *University) error {
	return DB.Create(university).Error
}

// 根据ID获取大学
func GetUniversityByID(id uint) (*University, error) {
	var university University
	err := DB.First(&university, id).Error
	return &university, err
}

// 根据名称获取大学
func GetUniversityByName(name string) (*University, error) {
	var university University
	err := DB.Where("name = ?", name).First(&university).Error
	return &university, err
}

// 更新大学
func UpdateUniversity(university *University) error {
	return DB.Save(university).Error
}

// 删除大学
func DeleteUniversity(id uint) error {
	return DB.Delete(&University{}, id).Error
}

// 获取大学列表（支持分页）
func GetUniversities(page, pageSize int) ([]University, int64, error) {
	var universities []University
	var total int64

	// 查询总数
	err := DB.Model(&University{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取数据列表
	offset := (page - 1) * pageSize
	err = DB.Offset(offset).Limit(pageSize).Find(&universities).Error
	return universities, total, err
}

// 获取所有大学（不分页）
func GetAllUniversities() ([]University, error) {
	var universities []University
	err := DB.Find(&universities).Error
	return universities, err
}
