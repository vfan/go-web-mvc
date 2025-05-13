package models

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
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
	// 密码加密
	log.Println("创建用户密码：", user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

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
	// 使用Updates而不是Save，只更新非零值字段
	return DB.Model(user).Updates(map[string]interface{}{
		"email":           user.Email,
		"role":            user.Role,
		"status":          user.Status,
		"last_login_time": user.LastLoginTime,
		// 不更新password和created_at字段
	}).Error
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

// 验证用户登录
func ValidateUserLogin(email, password string) (*User, error) {
	// 查找用户
	user, err := GetUserByEmail(email)

	tempPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	log.Println("tempPwd", string(tempPwd))
	log.Println("用户信息：", user)
	log.Println("输入的密码：", password)
	log.Println("数据库密码：", user.Password)

	if err != nil {
		log.Println("查找用户错误：", err)
		return nil, err
	}

	// 验证密码
	log.Println("开始验证密码")
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	log.Println("验证密码结果：", err)
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

// 更新用户最后登录时间
func UpdateUserLastLoginTime(userID uint) error {
	now := time.Now()
	return DB.Model(&User{}).Where("id = ?", userID).Update("last_login_time", &now).Error
}

// 创建重置密码函数
func ResetUserPassword(email, newPassword string) error {
	user, err := GetUserByEmail(email)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return DB.Save(user).Error
}
