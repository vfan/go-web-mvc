package dto

import "time"

// 学生创建请求
type StudentCreateRequest struct {
	Name           string     `json:"name" binding:"required"`
	Password       string     `json:"password" binding:"required,min=6"`
	Email          string     `json:"email" binding:"required,email"`
	Gender         *int32     `json:"gender,omitempty"`
	Birthday       *time.Time `json:"birthday,omitempty"`
	Phone          *string    `json:"phone,omitempty"`
	UniversityID   *int64     `json:"university_id,omitempty"`
	Major          *string    `json:"major,omitempty"`
	Education      *string    `json:"education,omitempty"`
	GraduationYear *int64     `json:"graduation_year,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Remarks        *string    `json:"remarks,omitempty"`
}

// 学生更新请求
type StudentUpdateRequest struct {
	Name           string     `json:"name,omitempty"`
	Gender         *int32     `json:"gender,omitempty"`
	Birthday       *time.Time `json:"birthday,omitempty"`
	Phone          *string    `json:"phone,omitempty"`
	UniversityID   *int64     `json:"university_id,omitempty"`
	Major          *string    `json:"major,omitempty"`
	Education      *string    `json:"education,omitempty"`
	GraduationYear *int64     `json:"graduation_year,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Remarks        *string    `json:"remarks,omitempty"`
	Avatar         *string    `json:"avatar,omitempty"`
}

// 学生重置密码请求
type StudentResetPasswordRequest struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// 学生响应
type StudentResponse struct {
	ID             int64               `json:"id"`
	Name           string              `json:"name"`
	Email          string              `json:"email"`
	Gender         *int32              `json:"gender"`
	Birthday       *time.Time          `json:"birthday"`
	Phone          *string             `json:"phone"`
	UniversityID   *int64              `json:"university_id"`
	University     *UniversityResponse `json:"university,omitempty"`
	Major          *string             `json:"major"`
	Education      *string             `json:"education"`
	GraduationYear *int64              `json:"graduation_year"`
	Status         *string             `json:"status"`
	Remarks        *string             `json:"remarks"`
	Avatar         *string             `json:"avatar"`
	LastLoginTime  *time.Time          `json:"last_login_time"`
	CreatedAt      *time.Time          `json:"created_at"`
	UpdatedAt      *time.Time          `json:"updated_at"`
	CreatedBy      *int64              `json:"created_by"`
	UpdatedBy      *int64              `json:"updated_by"`
}

// 学生列表响应
type StudentListResponse struct {
	List  []*StudentResponse `json:"list"`
	Total int64              `json:"total"`
	Page  int                `json:"page"`
	Size  int                `json:"size"`
}
