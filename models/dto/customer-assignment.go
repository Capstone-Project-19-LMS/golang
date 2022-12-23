package dto

import (
	"time"

	"gorm.io/gorm"
)

type CustomerAssignment struct {
	ID           string         `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	File         string         `json:"file"`
	Grade        int            `json:"grade" gorm:"notNull"`
	AssignmentID string         `json:"assignment_id"`
	CustomerID   string         `json:"customer_id"`
}
type CustomerAssignmentAcc struct {
	ID           string         `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	File         string         `json:"file"`
	Grade        int            `json:"grade" gorm:"notNull"`
	AssignmentID string         `json:"assignment_id"`
	CustomerID   string         `json:"customer_id"`
	Customer     struct {
		Name string `json:"name" gorm:"notNull;size:255"`
	} `json:"customer"`
}

type CustomerAssignmentTransaction struct {
	ID           string `json:"id" `
	File         string `json:"file" validate:"required"`
	Grade        int    `json:"grade" gorm:"notNull" validate:"required"`
	AssignmentID string `json:"assignment_id" validate:"required"`
	CustomerID   string `json:"customer_id" validate:"required"`
}
