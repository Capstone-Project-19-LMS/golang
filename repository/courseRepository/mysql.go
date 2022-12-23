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
func (cr *courseRepository) GetAllCourse(user dto.User) ([]dto.Course, error) {
	var courseModels []dto.GetCourseCategory
	// get data sub category from database by user
	var err error
	if user.Role == "instructor" {
		err = cr.db.Model(&model.Course{}).Preload("Category").Preload("CustomerCourses").Preload("Favorites").Preload("Ratings").Preload("Modules").Where("instructor_id = ?", user.ID).Find(&courseModels).Error
	} else if user.Role == "customer" {
		err = cr.db.Model(&model.Course{}).Preload("Category").Preload("CustomerCourses", "customer_id = ?", user.ID).Preload("Favorites", "customer_id = ?", user.ID).Preload("Ratings").Preload("Modules").Find(&courseModels).Error
	}
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
func (cr *courseRepository) GetCourseByID(id string) (dto.Course, error) {
	var courseModel dto.GetCourseCategory
	err := cr.db.Model(&model.Course{}).Preload("Category").Preload("CustomerCourses").Preload("Favorites").Preload("Ratings").Preload("Modules").Where("id = ? ", id).Find(&courseModel)
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

// GetCourseEnrollByID implements CourseRepository
func (cr *courseRepository) GetCourseEnrollByID(id string) ([]dto.CustomerCourseEnroll, error) {
	var customers []dto.CustomerCourseEnroll
	err := cr.db.Model(&model.Customer{}).Select("*", "customer_courses.id AS id", "customers.id AS customer_id" , "customer_courses.status AS status_enroll").Joins("JOIN customer_courses ON customer_courses.customer_id = customers.id").Where("customer_courses.course_id = ? ", id).Find(&customers)
	if err.Error != nil {
		return nil, err.Error
	}
	if err.RowsAffected <= 0 {
		return []dto.CustomerCourseEnroll{}, nil
	}
	return customers, nil
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
