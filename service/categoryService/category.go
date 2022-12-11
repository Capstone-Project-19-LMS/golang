package categoryService

import (
	"golang/helper"
	"golang/models/dto"
	"golang/repository/categoryRepository"

	"github.com/jinzhu/copier"
)

type CategoryService interface {
	CreateCategory(dto.CategoryTransaction) error
	DeleteCategory(id string) error
	GetCategoryByID(id string, user dto.User) (dto.GetCategory, error)
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
func (cs *categoryService) GetCategoryByID(id string, user dto.User) (dto.GetCategory, error) {
	category, err := cs.categoryRepo.GetCategoryByID(id, user)
	if err != nil {
		return dto.GetCategory{}, err
	}

	// get rating of all courses
	for i, course := range category.Courses {
		rating := helper.GetRatingCourse(course)
		category.Courses[i].Rating = rating

		// get number of module
		numberOfModule := len(course.Modules)
		category.Courses[i].NumberOfModules = numberOfModule

		if user.Role == "customer" {
			// get favorite of all courses
			favorite := helper.GetFavoriteCourse(course, user.ID)
			category.Courses[i].Favorite = favorite

			// get enrolled of all courses
			helper.GetEnrolledCourse(&course, user.ID)
			category.Courses[i].StatusEnroll = course.StatusEnroll
			category.Courses[i].ProgressModule = course.ProgressModule

			// get progress of all courses
			ProgressPercentage := helper.GetProgressCourse(&category.Courses[i])
			category.Courses[i].ProgressPercentage = ProgressPercentage

		}
	}

	var getCategory dto.GetCategory
	err = copier.Copy(&getCategory, &category)
	if err != nil {
		return dto.GetCategory{}, err
	}

	return getCategory, nil
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
