package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCode struct {
	ID        string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `json:"email"`
	Code      string         `json:"code" gorm:"notNull;size:255"`
}
