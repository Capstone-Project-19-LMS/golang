package model

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	ID string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Rating int `json:"rating" gorm:"notNull"`
	Testimonial string `json:"testimonial"`
	CustomerID string `json:"customer_id" gorm:"notNull;size:255"`
	CourseID string `json:"course_id" gorm:"notNull;size:255"`
}
