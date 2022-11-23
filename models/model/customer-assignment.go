package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerAssignment struct {
	ID string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	File string `json:"file" gorm:"notNull;size:255"`
	Grade int `json:"grade" gorm:"notNull"`
	AssignmentID string `json:"assignment_id" gorm:"notNull;size:255"`
	CustomerID string `json:"customer_id" gorm:"notNull;size:255"`
}
