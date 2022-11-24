package model

import (
	"time"

	"gorm.io/gorm"
)

type CustomerCode struct {
	ID         string `json:"id" gorm:"primaryKey;notNull;size:255"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	CustomerID string         `json:"customer_id" gorm:"notNull;size:255"`
	Code       string         `json:"code" gorm:"notNull;size:255"`
}
