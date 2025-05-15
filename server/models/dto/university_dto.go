package dto

import (
	"time"

	"gorm.io/gorm"
)

// 大学请求
type UniversityRequest struct {
	Name string `json:"name" binding:"required"`
}

// 大学响应
type UniversityResponse struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	CreatedBy *int64         `json:"created_by"`
	UpdatedBy *int64         `json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// 大学列表响应
type UniversityListResponse struct {
	List  []*UniversityResponse `json:"list"`
	Total int64                 `json:"total"`
	Page  int                   `json:"page"`
	Size  int                   `json:"size"`
}
