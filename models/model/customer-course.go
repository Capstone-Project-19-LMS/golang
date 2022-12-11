package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCourse struct {
	ID         string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CustomerID string         `json:"customer_id" gorm:"notNull;size:255"`
	CourseID   string         `json:"course_id" gorm:"notNull;size:255"`
	Status     bool           `json:"status" gorm:"notNull;default:true"`
	NoModule   int            `json:"no_module" gorm:"notNull;default:1"`
	IsFinish   bool           `json:"is_finish" gorm:"notNull;default:false"`
}
