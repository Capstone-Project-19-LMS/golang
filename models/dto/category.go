package dto

import "golang/models/model"

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required,unique"`
	Description string `json:"description" `
}

type CategoryGet struct {
	ID          string `json:"id" `
	Name        string `json:"name"`
	Description string `json:"description"`
	Courses     []model.Course
}

