package dto

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `json:"deleted_at"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Objective       string           `json:"objective"`
	Price           float64          `json:"price"`
	Discount        float64          `json:"discount"`
	Thumbnail       string           `json:"thumbnail"`
	Capacity        int              `json:"capacity"`
	InstructorID    string           `json:"instructor_id"`
	CategoryID      string           `json:"category_id"`
	Rating          float64          `json:"rating"`
	Favorite        bool             `json:"favorite"`
	NumberOfModules int              `json:"number_of_modules"`
	Enroll          bool             `json:"enroll"`
	CustomerCourses []CustomerCourse `json:"customer_courses"`
	Favorites       []Favorite       `json:"favorites"`
	Ratings         []Rating         `json:"ratings"`
	Modules         []Module         `json:"modules"`
}

type GetCourseCategory struct {
	ID              string           `json:"id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `json:"deleted_at"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Objective       string           `json:"objective"`
	Price           float64          `json:"price"`
	Discount        float64          `json:"discount"`
	Thumbnail       string           `json:"thumbnail"`
	Capacity        int              `json:"capacity"`
	InstructorID    string           `json:"instructor_id"`
	CategoryID      string           `json:"category_id"`
	Category        Category        `json:"category"`
	CustomerCourses []CustomerCourse `json:"customer_courses" gorm:"foreignKey:CourseID"` // foreignKey:CourseID is not needed
	Favorites       []Favorite       `json:"favorites" gorm:"foreignKey:CourseID"`         // foreignKey:CourseID is not needed
	Ratings         []Rating         `json:"ratings" gorm:"foreignKey:CourseID"`           // foreignKey:CourseID is not needed
	Modules         []Module         `json:"modules" gorm:"foreignKey:CourseID"`           // foreignKey:CourseID is not needed
}

type CourseTransaction struct {
	ID           string  `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Objective    string  `json:"objective"`
	Price        float64 `json:"price"`
	Discount     float64 `json:"discount"`
	Thumbnail    string  `json:"thumbnail"`
	Capacity     int     `json:"capacity" validate:"required,numeric"`
	InstructorID string  `json:"instructor_id"`
	CategoryID   string  `json:"category_id" validate:"required,alphanum"`
}

type GetCourse struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Objective       string   `json:"objective"`
	Price           float64  `json:"price"`
	Discount        float64  `json:"discount"`
	Thumbnail       string   `json:"thumbnail"`
	Capacity        int      `json:"capacity"`
	InstructorID    string   `json:"instructor_id"`
	Category        Category `json:"category"`
	Rating          float64  `json:"rating"`
	Favorite        bool     `json:"favorite"`
	NumberOfModules int      `json:"number_of_modules"`
	Enroll          bool     `json:"enroll"`
}

type GetCourseInstructor struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Objective       string   `json:"objective"`
	Price           float64  `json:"price"`
	Discount        float64  `json:"discount"`
	Thumbnail       string   `json:"thumbnail"`
	Capacity        int      `json:"capacity"`
	InstructorID    string   `json:"instructor_id"`
	Category        Category `json:"category" gorm:"references:CategoryID"` 
	Rating          float64  `json:"rating"`
	NumberOfModules int      `json:"number_of_modules"`
}

type CourseCustomerEnroll struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Objective       string           `json:"objective"`
	Price           float64          `json:"price"`
	Discount        float64          `json:"discount"`
	Thumbnail       string           `json:"thumbnail"`
	Capacity        int              `json:"capacity"`
	InstructorID    string           `json:"instructor_id"`
	CategoryID      string           `json:"category_id"`
	NumberOfModules int              `json:"number_of_modules"`
	CustomerCourses []CustomerCourse `json:"customer_courses"`
}
