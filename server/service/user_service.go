package service

import (
	"mvc-demo/dao"
	"mvc-demo/dao/model"
	"time"
)

// UserService 用户服务
type UserService struct {
	userDAO *dao.UserDAO
}

// NewUserService 创建用户服务实例
func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *model.User) error {
	return s.userDAO.Create(user)
}

// GetUserByID 根据ID获取用户
func (s *UserService) GetUserByID(id int64) (*model.User, error) {
	return s.userDAO.GetByID(id)
}

// GetUserByEmail 根据邮箱获取用户
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userDAO.GetByEmail(email)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(user *model.User) error {
	return s.userDAO.Update(user)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id int64) error {
	return s.userDAO.Delete(id)
}

// GetUserList 获取用户列表（支持分页）
func (s *UserService) GetUserList(page, pageSize int) ([]*model.User, int64, error) {
	return s.userDAO.GetList(page, pageSize)
}

// Login 用户登录
func (s *UserService) Login(email, password string) (*model.User, error) {
	// 验证登录信息
	user, err := s.userDAO.ValidateLogin(email, password)
	if err != nil {
		return nil, err
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLoginTime = &now
	err = s.userDAO.UpdateLastLoginTime(user.ID)
	if err != nil {
		// 只记录错误，不影响登录
		// log.Printf("更新登录时间失败: %v", err)
	}

	return user, nil
}

// ResetPassword 重置密码
func (s *UserService) ResetPassword(email, newPassword string) error {
	return s.userDAO.ResetPassword(email, newPassword)
}
