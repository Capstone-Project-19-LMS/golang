package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name		string `json:"name" gorm:"notNull;size:255;uniqueIndex"`
	Description string `json:"description"`
	Objective string `json:"objective"`
	Price		float64 `json:"price" gorm:"notNull;default:0"`
	Discount	float64 `json:"discount" gorm:"notNull;default:0"`
	Thumbnail	string `json:"thumbnail" gorm:"size:255"`
	Capacity	int `json:"capacity" gorm:"notNull;default:0"`
	InstructorID string `json:"instructor_id" gorm:"notNull;size:255"`
	CategoryID string `json:"category_id" gorm:"notNull;size:255"`
	CustomerCourses []CustomerCourse
	Favorites []Favorite
	Ratings []Rating
	Modules []Module
}