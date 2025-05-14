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

// UpdateUniversity 更新大学信息
func (s *UniversityService) UpdateUniversity(university *model.University) error {
	return s.universityDAO.Update(university)
}

// DeleteUniversity 删除大学
func (s *UniversityService) DeleteUniversity(id int64) error {
	return s.universityDAO.Delete(id)
}

// GetUniversityList 获取大学列表（支持分页）
func (s *UniversityService) GetUniversityList(page, pageSize int) ([]*model.University, int64, error) {
	return s.universityDAO.GetList(page, pageSize)
}

// GetAllUniversities 获取所有大学（不分页）
func (s *UniversityService) GetAllUniversities() ([]*model.University, error) {
	return s.universityDAO.GetAll()
}
