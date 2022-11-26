package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `json:"name" gorm:"notNull;size:255"`
	Description string         `json:"description"`
	Courses     []Course
}
