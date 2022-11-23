package categoryRepository

import "golang/models/dto"

type CategoryRepository interface {
	CreateCategory(dto.Category) error
	DeleteCategory(id string) error
	GetCategoryByID(id string) (dto.Category, error)
	GetAllCategory() ([]dto.Category, error)
	UpdateCategory(dto.Category) error
}