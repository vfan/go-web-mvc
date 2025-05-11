package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型（对应users表）
type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Email         string         `gorm:"uniqueIndex;size:100;not null" json:"email"`
	Password      string         `gorm:"size:255;not null" json:"-"`     // 不返回密码
	Role          int8           `gorm:"default:2;not null" json:"role"` // 1:管理员,2:普通用户
	LastLoginTime *time.Time     `gorm:"default:null" json:"last_login_time"`
	Status        int8           `gorm:"default:1;not null" json:"status"` // 0:禁用,1:启用
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// 创建用户
func CreateUser(user *User) error {
	return DB.Create(user).Error
}

// 根据ID获取用户
func GetUserByID(id uint) (*User, error) {
	var user User
	err := DB.First(&user, id).Error
	return &user, err
}

// 根据邮箱获取用户
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// 更新用户
func UpdateUser(user *User) error {
	return DB.Save(user).Error
}

// 删除用户
func DeleteUser(id uint) error {
	return DB.Delete(&User{}, id).Error
}

// 获取用户列表（支持分页）
func GetUsers(page, pageSize int) ([]User, int64, error) {
	var users []User
	var total int64

	// 查询总数
	err := DB.Model(&User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取数据列表
	offset := (page - 1) * pageSize
	err = DB.Offset(offset).Limit(pageSize).Find(&users).Error
	return users, total, err
}
