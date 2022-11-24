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
func (cr *categoryRepository) DeleteCategory(id string) error {
	// delete data category from database by id
	err := cr.db.Select("courses").Where("id = ?", id).Delete(&model.Category{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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
