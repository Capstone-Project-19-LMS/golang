package categoryRepository

import "golang/models/dto"

type CategoryRepository interface {
	CreateCategory(dto.CategoryTransaction) error
	DeleteCategory(id string) error
	GetCategoryByID(id string, user dto.User) (dto.Category, error)
	GetAllCategory() ([]dto.CategoryTransaction, error)
	UpdateCategory(dto.CategoryTransaction) error
}
