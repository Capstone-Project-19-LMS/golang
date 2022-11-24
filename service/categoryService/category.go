package categoryService

import (
	"golang/helper"
	"golang/models/dto"
	"golang/repository/categoryRepository"
)

type CategoryService interface {
	CreateCategory(dto.Category) error
	DeleteCategory(id string) error
	GetCategoryByID(id string) (dto.Category, error)
	GetAllCategory() ([]dto.Category, error)
	UpdateCategory(dto.Category) error
}

type categoryService struct {
	categoryRepo categoryRepository.CategoryRepository
}

// CreateCategory implements CategoryService
func (cs *categoryService) CreateCategory(category dto.Category) error {
	id := helper.GenerateUUID()
	category.ID = id
	err := cs.categoryRepo.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory implements CategoryService
func (cs *categoryService) DeleteCategory(id string) error {
	// call repository to delete account
	err := cs.categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCategory implements CategoryService
func (*categoryService) GetAllCategory() ([]dto.Category, error) {
	panic("unimplemented")
}

// GetCategoryByID implements CategoryService
func (*categoryService) GetCategoryByID(id string) (dto.Category, error) {
	panic("unimplemented")
}

// UpdateCategory implements CategoryService
func (*categoryService) UpdateCategory(dto.Category) error {
	panic("unimplemented")
}

func NewCategoryService(categoryRepo categoryRepository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}
