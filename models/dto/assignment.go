package dto

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	ID                  string               `json:"id"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	DeletedAt           gorm.DeletedAt       `json:"deleted_at"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	ModuleID            string               `json:"module_id"`
	CustomerAssignments []CustomerAssignment `json:"customer_assignments"`
}
type AssignmentCourse struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	ModuleID    string         `json:"module_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
type GetAssignment struct {
	ID                  string               `json:"id"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	DeletedAt           gorm.DeletedAt       `json:"deleted_at"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	ModuleID            string               `json:"module_id"`
	CustomerAssignments []CustomerAssignment `json:"customer_assignments" `
}

type AssignmentTransaction struct {
	ID          string `json:"id" `
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ModuleID    string `json:"module_id"  validate:"required"`
}
