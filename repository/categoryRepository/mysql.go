package categoryRepository

import (
	"golang/models/dto"
	"golang/models/model"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements CategoryRepository
func (cr *categoryRepository) CreateCategory(category dto.Category) error {
	err := cr.db.Model(&model.Category{}).Create(&model.Category{
		ID: 		category.ID,
		Name:        category.Name,
		Description: category.Description,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory implements CategoryRepository
func (*categoryRepository) DeleteCategory(id string) error {
	panic("unimplemented")
}

// GetAllCategory implements CategoryRepository
func (*categoryRepository) GetAllCategory() ([]dto.Category, error) {
	panic("unimplemented")
}

// GetCategoryByID implements CategoryRepository
func (*categoryRepository) GetCategoryByID(id string) (dto.Category, error) {
	panic("unimplemented")
}

// UpdateCategory implements CategoryRepository
func (*categoryRepository) UpdateCategory(dto.Category) error {
	panic("unimplemented")
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
