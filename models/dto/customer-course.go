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
}

type CustomerCourseTransaction struct {
	ID string `json:"id"`
	CustomerID string `json:"customer_id"`
	CourseID string `json:"course_id" validate:"required,alphanum"` 
	Status     bool           `json:"status"`
}
