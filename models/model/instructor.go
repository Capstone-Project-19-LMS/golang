package model

import (
	"time"

	"gorm.io/gorm"
)

type Instructor struct {
	ID           string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `json:"name" gorm:"notNull;size:255"`
	Email        string         `json:"email" gorm:"notNull;unique;size:255"`
	Password     string         `json:"password" gorm:"notNull"`
	ProfileImage string         `json:"profile_image" gorm:"size:255;default:null"`
	Courses      []Course
}
