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
func (cr *categoryRepository) CreateCategory(category dto.CategoryTransaction) error {
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
func (cr *categoryRepository) GetAllCategory() ([]dto.CategoryTransaction, error) {
	var categories []dto.CategoryTransaction
	// get data sub category from database by user
	err := cr.db.Model(&model.Category{}).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryByID implements CategoryRepository
func (cr *categoryRepository) GetCategoryByID(id string) (dto.Category, error) {
	var category dto.Category
	err := cr.db.Model(&model.Category{}).Preload("Courses").Where("id = ?", id).Find(&category)
	if err.Error != nil {
		return dto.Category{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Category{}, gorm.ErrRecordNotFound
	}
	return category, nil
}

// UpdateCategory implements CategoryRepository
func (cr *categoryRepository) UpdateCategory(category dto.CategoryTransaction) error {
	// update account with new data
	err := cr.db.Model(&model.Category{}).Where("id = ?", category.ID).Updates(&model.Category{
		Name: category.Name,
		Description: category.Description,
	})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
