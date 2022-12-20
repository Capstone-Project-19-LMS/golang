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
		ID:          category.ID,
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
func (cr *categoryRepository) GetCategoryByID(id string, user dto.User) (dto.Category, error) {
	var category dto.Category
	var err *gorm.DB
	if user.Role == "instructor" {
		err = cr.db.Model(&model.Category{}).Preload("Courses", "instructor_id = ?", user.ID).Preload("Courses.CustomerCourses").Preload("Courses.Favorites").Preload("Courses.Ratings").Preload("Courses.Modules").Where("id = ?", id).Find(&category)
	} else if user.Role == "customer" {
		err = cr.db.Model(&model.Category{}).Preload("Courses.CustomerCourses", "customer_id = ?", user.ID).Preload("Courses.Favorites").Preload("Courses.Ratings").Preload("Courses.Modules").Where("id = ?", id).Find(&category)
	}
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
		Name:        category.Name,
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
