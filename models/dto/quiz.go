package dto

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
type QuizTransaction struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id" validate:"required"`
	Link     string `json:"link" validate:"required"`
}
type TakeQuizTransaction struct {
	CourseID   string `json:"course_id" validate:"required"`
	CustomerID string `json:"customer_id" validate:"required"`
}
