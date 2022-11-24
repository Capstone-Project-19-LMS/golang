package dto

import (
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name		   string `json:"name"`
	Content		string `json:"content"`
	CourseID string `json:"course_id"`
	MediaModules []MediaModule
	Assignment Assignment
}
