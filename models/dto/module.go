package dto

import (
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `json:"name"`
	Content   string `json:"content"`
	CourseID  string `json:"course_id"`
	Course    struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Objective   string `json:"objective"`
	} `json:"course"`
	NoModule int `json:"no_module"`
}

type ModuleCourse struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `json:"name"`
	Content   string `json:"content"`
	CourseID  string `json:"course_id"`
	Course    struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Objective   string `json:"objective"`
	} `json:"course"`
	MediaModules []MediaModule
	Assignment   Assignment
	NoModule     int `json:"no_module"`
}
type ModuleAcc struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `json:"name"`
	Content   string `json:"content"`
	CourseID  string `json:"course_id"`
	Course    struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Objective   string `json:"objective"`
	} `json:"course"`
	MediaModules []MediaModule
	Assignment   Assignment
	NoModule     int `json:"no_module"`
}
type ModuleCourseAcc struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string `json:"name"`
	Content   string `json:"content"`
	CourseID  string `json:"course_id"`
	Course    struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Objective   string `json:"objective"`
	} `json:"course"`
	MediaModules []MediaModule
	Assignment   Assignment
	NoModule     int `json:"no_module"`
}

type ModuleTransaction struct {
	ID            string `json:"id"`
	Name          string `json:"name" validate:"required"`
	Content       string `json:"content" validate:"required"`
	CourseID      string `json:"course_id" validate:"required"`
	NoModule      int    `json:"no_module" validate:"required"`
	MediaModuleID string `json:"media_module_id"`
	Url           string `json:"url" validate:"required"`
}
