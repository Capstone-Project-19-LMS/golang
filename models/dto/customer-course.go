package dto

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCourse struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	CustomerID string `json:"customer_id"`
	CourseID string `json:"course_id"`
	Status     bool           `json:"status"`
	NoModule   int            `json:"no_module"`
	IsFinish   bool           `json:"is_finish"`
}

type CustomerCourseTransaction struct {
	ID string `json:"id"`
	CustomerID string `json:"customer_id" validate:"required,alphanum"`
	CourseID string `json:"course_id" validate:"required,alphanum"` 
	Status     bool           `json:"status" validate:"required,boolean"`
	NoModule   int            `json:"no_module"`
	IsFinish   bool           `json:"is_finish"`
}
