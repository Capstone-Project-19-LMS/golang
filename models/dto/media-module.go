package dto

import (
	"time"

	"gorm.io/gorm"
)

type MediaModule struct {
	ID string `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`json:"deleted_at"`
	Url string `json:"url"`
	ModuleID string `json:"module_id"`
}
