package categoryService

import (
	"golang/helper"
	"golang/models/dto"
	"golang/repository/categoryRepository"
)

type CategoryService interface {
	CreateCategory(dto.CategoryTransaction) error
	DeleteCategory(id string) error
	GetCategoryByID(id string, user dto.User) (dto.Category, error)
	GetAllCategory() ([]dto.CategoryTransaction, error)
	UpdateCategory(dto.CategoryTransaction) error
}

type categoryService struct {
	categoryRepo categoryRepository.CategoryRepository
}

// CreateCategory implements CategoryService
func (cs *categoryService) CreateCategory(category dto.CategoryTransaction) error {
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
func (cs *categoryService) GetAllCategory() ([]dto.CategoryTransaction, error) {
	categories, err := cs.categoryRepo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryByID implements CategoryService
func (cs *categoryService) GetCategoryByID(id string, user dto.User) (dto.Category, error) {
	category, err := cs.categoryRepo.GetCategoryByID(id, user)
	if err != nil {
		return dto.Category{}, err
	}
	return category, nil
}

// UpdateCategory implements CategoryService
func (cs *categoryService) UpdateCategory(category dto.CategoryTransaction) error {
	// call repository to update category
	err := cs.categoryRepo.UpdateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryService(categoryRepo categoryRepository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}
