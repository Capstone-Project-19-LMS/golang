package dto

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          string         `json:"id" `
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Courses     []Course       `json:"courses" gorm:"foreignKey:CategoryID"`
}

type CategoryTransaction struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" `
}

type GetCategory struct {
	ID          string         `json:"id" `
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Courses     []GetCourseWithoutCategory       `json:"courses" gorm:"foreignKey:CategoryID"`
}

type GetCategoryInstructor struct {
	ID          string         `json:"id" `
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Courses     []GetCourseInstructorWithoutCategory       `json:"courses" gorm:"foreignKey:CategoryID"`
}
