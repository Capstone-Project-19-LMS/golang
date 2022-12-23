package model

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	ID        string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CourseID  string         `json:"course_id"`
	Link      string         `json:"link"`
}
