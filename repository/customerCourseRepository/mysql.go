package customerCourseRepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type customerCourseRepository struct {
	db *gorm.DB
}



// DeleteCustomerCourse implements CustomerCourseRepository
func (ccr *customerCourseRepository) DeleteCustomerCourse(id string) error {
	err := ccr.db.Unscoped().Delete(&model.CustomerCourse{}, "id = ?", id)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

// GetCustomerCourse implements CustomerCourseRepository
func (ccr *customerCourseRepository) GetCustomerCourse(courseID string, customerID string) (dto.CustomerCourse, error) {
	var customerCourse dto.CustomerCourse
	err := ccr.db.Model(&model.CustomerCourse{}).Where("course_id = ? AND customer_id = ?", courseID, customerID).First(&customerCourse)
	if err.Error != nil {
		return dto.CustomerCourse{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.CustomerCourse{}, gorm.ErrRecordNotFound
	}
	return customerCourse, nil
}


// HistoryCourse implements CustomerCourseRepository
func (ccr *customerCourseRepository) GetHistoryCourseByCustomerID(customerId string) ([]dto.Course, error) {
	var courseModels []dto.CourseCustomerEnroll

	// get data course from database by customer id
	err := ccr.db.Model(&model.Course{}).Joins("JOIN customer_courses ON customer_courses.course_id = courses.id").Preload("Category").Preload("CustomerCourses", "customer_id = ?", customerId).Preload("Favorites", "customer_id = ?", customerId).Preload("Ratings").Preload("Modules").Unscoped().Where("customer_courses.customer_id = ?", customerId).Find(&courseModels).Error
	if err != nil {
		return nil, err
	}

	var courses []dto.Course
	err = copier.Copy(&courses, &courseModels)
	if err != nil {
		return nil, err
	}
	return courses, err
}

// TakeCourse implements CustomerCourseRepository
func (ccr *customerCourseRepository) TakeCourse(customerCourse dto.CustomerCourseTransaction) error {
	var customerCourseModel model.CustomerCourse
	err := copier.Copy(&customerCourseModel, &customerCourse)
	if err != nil {
		return err
	}
	// save customer course to database
	err = ccr.db.Model(&model.CustomerCourse{}).Create(&customerCourseModel).Error
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerCourseRepository(db *gorm.DB) CustomerCourseRepository {
	return &customerCourseRepository{
		db: db,
	}
}
