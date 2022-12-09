package dto

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	ID                  string         `json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
	Title               string         `json:"title"`
	Description         string         `json:"description"`
	ModuleID            string         `json:"module_id"`
	CustomerAssignments []CustomerAssignment
}
type GetAssignment struct {
	ID                  string         `json:"id"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
	Title               string         `json:"title"`
	Description         string         `json:"description"`
	ModuleID            string         `json:"module_id"`
	CustomerAssignments []CustomerAssignment
}

type AssignmentTransaction struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ModuleID    string `json:"module_id"`
}
