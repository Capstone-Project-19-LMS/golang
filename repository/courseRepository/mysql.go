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
	copier.Copy(&courseModel, &course)
	
	err := cr.db.Model(&model.Course{}).Create(&courseModel).Error
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
	var courses []dto.Course
	// get data sub category from database by user
	err := cr.db.Model(&model.Course{}).Preload("CustomerCourses").Preload("Favorites").Preload("Ratings").Preload("Modules").Where("instructor_id = ?", instructorId).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}

// GetCourseByID implements CourseRepository
func (cr *courseRepository) GetCourseByID(id, instructorId string) (dto.Course, error) {
	var course dto.Course
	err := cr.db.Model(&model.Course{}).Preload("CustomerCourses").Preload("Favorites").Preload("Ratings").Preload("Modules").Where("id = ? AND instructor_id = ?", id, instructorId).Find(&course)
	if err.Error != nil {
		return dto.Course{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Course{}, gorm.ErrRecordNotFound
	}
	return course, nil
}

// UpdateCourse implements CourseRepository
func (*courseRepository) UpdateCourse(dto.CourseTransaction) error {
	panic("unimplemented")
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{
		db: db,
	}
}
