package courseRepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

// CreateCourse implements CourseRepository
func (cr *courseRepository) CreateCourse(course dto.CourseTransaction) error {
	var courseModel model.Course
	err := copier.Copy(&courseModel, &course)
	if err != nil {
		return err
	}

	err = cr.db.Model(&model.Course{}).Create(&courseModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCourse implements CourseRepository
func (cr *courseRepository) DeleteCourse(id string) error {
	// delete data course from database by id
	err := cr.db.Select("modules", "Favorites", "Ratings").Where("id = ?", id).Delete(&model.Course{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAllCourse implements CourseRepository
func (cr *courseRepository) GetAllCourse(instructorId string) ([]dto.Course, error) {
	var courseModels []model.Course
	// get data sub category from database by user
	err := cr.db.Model(&model.Course{}).Preload("CustomerCourses").Preload("Favorites").Preload("Ratings").Preload("Modules").Where("instructor_id = ?", instructorId).Find(&courseModels).Error
	if err != nil {
		return nil, err
	}
	// copy data from model to dto
	var courses []dto.Course
	err = copier.Copy(&courses, &courseModels)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// GetCourseByID implements CourseRepository
func (cr *courseRepository) GetCourseByID(id, instructorId string) (dto.Course, error) {
	var courseModel model.Course
	err := cr.db.Model(&model.Course{}).Preload("CustomerCourses").Preload("Favorites").Preload("Ratings").Preload("Modules").Where("id = ? AND instructor_id = ?", id, instructorId).Find(&courseModel)
	if err.Error != nil {
		return dto.Course{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Course{}, gorm.ErrRecordNotFound
	}
	
	// copy data from model to dto
	var course dto.Course
	errCopy := copier.Copy(&course, &courseModel)
	if errCopy != nil {
		return dto.Course{}, errCopy
	}
	return course, nil
}

// UpdateCourse implements CourseRepository
func (cr *courseRepository) UpdateCourse(course dto.CourseTransaction) error {
	var courseModel model.Course
	errCopy := copier.Copy(&courseModel, &course)
	if errCopy != nil {
		return errCopy
	}
	// update account with new data
	err := cr.db.Model(&model.Course{}).Where("id = ?", course.ID).Updates(&courseModel)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}
