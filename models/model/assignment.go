package model

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	ID                  string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	Title               string         `json:"title" gorm:"notNull;size:255"`
	Description         string         `json:"description" gorm:"notNull;size:255"`
	ModuleID            string         `json:"module_id" gorm:"notNull;size:255"`
	CustomerAssignments []CustomerAssignment
}
