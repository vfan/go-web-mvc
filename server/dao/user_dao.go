package dao

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"mvc-demo/dao/model"
)

// UserDAO 用户数据访问对象
type UserDAO struct {
	DB *gorm.DB
}

// NewUserDAO 创建用户DAO实例
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{DB: db}
}

// Create 创建用户
func (dao *UserDAO) Create(user *model.User) error {
	// 密码加密
	log.Println("创建用户密码：", user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return dao.DB.Create(user).Error
}

// GetByID 根据ID获取用户
func (dao *UserDAO) GetByID(id int64) (*model.User, error) {
	var user model.User
	err := dao.DB.First(&user, id).Error
	return &user, err
}

// GetByEmail 根据邮箱获取用户
func (dao *UserDAO) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := dao.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// Update 更新用户
func (dao *UserDAO) Update(user *model.User) error {
	// 使用Updates而不是Save，只更新非零值字段
	return dao.DB.Model(user).Updates(map[string]interface{}{
		"email":           user.Email,
		"role":            user.Role,
		"status":          user.Status,
		"last_login_time": user.LastLoginTime,
		"username":        user.Username,
		// 不更新password和created_at字段
	}).Error
}

// Delete 删除用户
func (dao *UserDAO) Delete(id int64) error {
	return dao.DB.Delete(&model.User{}, id).Error
}

// GetList 获取用户列表（支持分页）
func (dao *UserDAO) GetList(page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	// 查询总数
	err := dao.DB.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取数据列表
	offset := (page - 1) * pageSize
	err = dao.DB.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}

// ValidateLogin 验证用户登录
func (dao *UserDAO) ValidateLogin(email, password string) (*model.User, error) {
	// 查找用户
	user, err := dao.GetByEmail(email)
	if err != nil {
		log.Println("查找用户错误：", err)
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("密码错误：", err)
		return nil, err
	}

	// 验证用户状态
	if user.Status != 1 {
		log.Println("用户状态不可用")
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}

// UpdateLastLoginTime 更新用户最后登录时间
func (dao *UserDAO) UpdateLastLoginTime(userID int64) error {
	now := time.Now()
	return dao.DB.Model(&model.User{}).Where("id = ?", userID).Update("last_login_time", &now).Error
}

// ResetPassword 重置用户密码
func (dao *UserDAO) ResetPassword(email, newPassword string) error {
	user, err := dao.GetByEmail(email)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return dao.DB.Save(user).Error
}
