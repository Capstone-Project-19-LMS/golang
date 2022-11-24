package dto

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID              string `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time	`json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Price           float64        `json:"price"`
	Discount        float64        `json:"discount"`
	Thumbnail       string         `json:"thumbnail"`
	Capacity        int            `json:"capacity"`
	InstructorID    string         `json:"instructor_id"`
	CategoryID      string         `json:"category_id"`
	// CustomerCourses []CustomerCourse
	// Favorites       []Favorite
	// Ratings         []Rating
	// Modules         []Module
}