package dto

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `json:"deleted_at"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Objective       string           `json:"objective"`
	Price           float64          `json:"price"`
	Discount        float64          `json:"discount"`
	Thumbnail       string           `json:"thumbnail"`
	Capacity        int              `json:"capacity"`
	InstructorID    string           `json:"instructor_id"`
	CategoryID      string           `json:"category_id"`
	Rating float64 `json:"rating"`
	CustomerCourses []CustomerCourse `json:"customer_courses"`
	Favorites       []Favorite       `json:"favorites"`
	Ratings         []Rating         `json:"ratings"`
	Modules         []Module         `json:"modules"`
}

type CourseTransaction struct {
	ID           string  `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Objective    string  `json:"objective"`
	Price        float64 `json:"price"`
	Discount     float64 `json:"discount"`
	Thumbnail    string  `json:"thumbnail"`
	Capacity     int     `json:"capacity" validate:"required,numeric"`
	InstructorID string  `json:"instructor_id"`
	CategoryID   string  `json:"category_id" validate:"required,alphanum"`
}
