package dto

import (
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID         string         `json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	CustomerID string         `json:"customer_id"`
	CourseID   string         `json:"course_id"`
}
