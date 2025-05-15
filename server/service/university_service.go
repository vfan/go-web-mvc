package service

import (
	"mvc-demo/dao"
	"mvc-demo/dao/model"
)

// UniversityService 大学服务
type UniversityService struct {
	universityDAO *dao.UniversityDAO
}

// NewUniversityService 创建大学服务实例
func NewUniversityService(universityDAO *dao.UniversityDAO) *UniversityService {
	return &UniversityService{
		universityDAO: universityDAO,
	}
}

// CreateUniversity 创建大学
func (s *UniversityService) CreateUniversity(university *model.University) error {
	return s.universityDAO.Create(university)
}

// GetUniversityByID 根据ID获取大学
func (s *UniversityService) GetUniversityByID(id int64) (*model.University, error) {
	return s.universityDAO.GetByID(id)
}

// GetUniversityByName 根据名称获取大学
func (s *UniversityService) GetUniversityByName(name string) (*model.University, error) {
	return s.universityDAO.GetByName(name)
}

// CheckUniversityNameExistsWithDeleted 检查大学名称是否存在（包括已软删除的记录）
func (s *UniversityService) CheckUniversityNameExistsWithDeleted(name string) (bool, error) {
	return s.universityDAO.CheckNameExistsWithDeleted(name)
}

// CheckUniversityNameExistsExcludeID 检查排除某ID外是否存在同名大学（包括已软删除的记录）
func (s *UniversityService) CheckUniversityNameExistsExcludeID(name string, excludeID int64) (bool, error) {
	return s.universityDAO.CheckNameExistsExcludeID(name, excludeID)
}

// UpdateUniversity 更新大学信息
func (s *UniversityService) UpdateUniversity(university *model.University) error {
	return s.universityDAO.Update(university)
}

// DeleteUniversity 删除大学
func (s *UniversityService) DeleteUniversity(id int64) error {
	return s.universityDAO.Delete(id)
}

// GetUniversityList 获取大学列表（支持分页和显示已删除）
func (s *UniversityService) GetUniversityList(page, pageSize int, showDeleted bool) ([]*model.University, int64, error) {
	return s.universityDAO.GetList(page, pageSize, showDeleted)
}

// GetAllUniversities 获取所有大学（不分页）
func (s *UniversityService) GetAllUniversities() ([]*model.University, error) {
	return s.universityDAO.GetAll()
}

// RestoreUniversity 恢复已删除的大学
func (s *UniversityService) RestoreUniversity(id int64) error {
	return s.universityDAO.Restore(id)
}

// GetUniversityByIDWithDeleted 根据ID获取大学（包括已删除的）
func (s *UniversityService) GetUniversityByIDWithDeleted(id int64) (*model.University, error) {
	return s.universityDAO.GetByIDWithDeleted(id)
}
