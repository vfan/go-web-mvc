package dao

import (
	"mvc-demo/dao/model"

	"gorm.io/gorm"
)

// UniversityDAO 大学数据访问对象
type UniversityDAO struct {
	DB *gorm.DB
}

// NewUniversityDAO 创建大学DAO实例
func NewUniversityDAO(db *gorm.DB) *UniversityDAO {
	return &UniversityDAO{DB: db}
}

// Create 创建大学
func (dao *UniversityDAO) Create(university *model.University) error {
	return dao.DB.Create(university).Error
}

// GetByID 根据ID获取大学
func (dao *UniversityDAO) GetByID(id int64) (*model.University, error) {
	var university model.University
	err := dao.DB.First(&university, id).Error
	return &university, err
}

// GetByName 根据名称获取大学
func (dao *UniversityDAO) GetByName(name string) (*model.University, error) {
	var university model.University
	err := dao.DB.Where("name = ?", name).First(&university).Error
	return &university, err
}

// CheckNameExistsWithDeleted 检查大学名称是否存在（包括已软删除的记录）
func (dao *UniversityDAO) CheckNameExistsWithDeleted(name string) (bool, error) {
	var count int64
	err := dao.DB.Unscoped().Model(&model.University{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

// CheckNameExistsExcludeID 检查排除某ID外是否存在同名大学（包括已软删除的记录）
func (dao *UniversityDAO) CheckNameExistsExcludeID(name string, excludeID int64) (bool, error) {
	var count int64
	err := dao.DB.Unscoped().Model(&model.University{}).
		Where("name = ? AND id != ?", name, excludeID).Count(&count).Error
	return count > 0, err
}

// Update 更新大学
func (dao *UniversityDAO) Update(university *model.University) error {
	return dao.DB.Save(university).Error
}

// Delete 删除大学
func (dao *UniversityDAO) Delete(id int64) error {
	return dao.DB.Delete(&model.University{}, id).Error
}

// GetList 获取大学列表（支持分页）
func (dao *UniversityDAO) GetList(page, pageSize int) ([]*model.University, int64, error) {
	var universities []*model.University
	var total int64

	// 查询总数
	err := dao.DB.Model(&model.University{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取数据列表
	offset := (page - 1) * pageSize
	err = dao.DB.Offset(offset).Limit(pageSize).Find(&universities).Error
	return universities, total, err
}

// GetAll 获取所有大学（不分页）
func (dao *UniversityDAO) GetAll() ([]*model.University, error) {
	var universities []*model.University
	err := dao.DB.Find(&universities).Error
	return universities, err
}
