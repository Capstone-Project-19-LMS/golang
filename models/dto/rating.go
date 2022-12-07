package dto

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	ID string `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Rating int `json:"rating"`
	Testimonial string `json:"testimonial"`
	IsPublish bool `json:"is_publish"`
	CustomerID string `json:"customer_id"`
	CourseID string `json:"course_id"`
}

type RatingTransaction struct {
	ID string `json:"id"`
	Rating int `json:"rating" validate:"required,min=1,max=5,numeric"`
	Testimonial string `json:"testimonial" validate:"required"`
	IsPublish bool `json:"is_publish"`
	CustomerID string `json:"customer_id"`
	CourseID string `json:"course_id" validate:"required,alphanum"`
}
