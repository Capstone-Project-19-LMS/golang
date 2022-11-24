package model

import (
	"time"

	"gorm.io/gorm"
)

type MediaModule struct {
	ID        string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Url       string         `json:"url" gorm:"notNull;size:255"`
	ModuleID  string         `json:"module_id" gorm:"notNull;size:255"`
}
