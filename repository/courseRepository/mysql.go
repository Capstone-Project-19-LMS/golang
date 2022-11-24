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
func (*courseRepository) DeleteCourse(id string) error {
	panic("unimplemented")
}

// GetAllCourse implements CourseRepository
func (*courseRepository) GetAllCourse(instructorId string) ([]dto.CourseTransaction, error) {
	panic("unimplemented")
}

// GetCourseByID implements CourseRepository
func (*courseRepository) GetCourseByID(id string) (dto.Course, error) {
	panic("unimplemented")
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
