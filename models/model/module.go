package model

import (
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name		   string `json:"name" gorm:"notNull;size:255"`
	Content		string `json:"content"`
	CourseID string `json:"course_id" gorm:"notNull;size:255"`
	MediaModules []MediaModule
	Assignment Assignment
}
