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


// GetCustomerEnrollByID implements CustomerCourseRepository
func (ccr *customerCourseRepository) GetCustomerCourseEnrollByID(id string) (dto.CustomerCourseEnroll, error) {
	var customer dto.CustomerCourseEnroll
	err := ccr.db.Model(&model.Customer{}).Select("*", "customer_courses.id AS id", "customers.id AS customer_id" , "customer_courses.status AS status_enroll").Joins("JOIN customer_courses ON customer_courses.customer_id = customers.id").Where("customer_courses.id = ?", id).Find(&customer)
	if err.Error != nil {
		return dto.CustomerCourseEnroll{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.CustomerCourseEnroll{}, nil
	}
	return customer, nil
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

// UpdateEnrollmentStatus implements CustomerCourseRepository
func (ccr *customerCourseRepository) UpdateEnrollmentStatus(customerCourse dto.CustomerCourseTransaction) error {
	err := ccr.db.Model(&model.CustomerCourse{}).Where("course_id = ? AND customer_id = ?", customerCourse.CourseID, customerCourse.CustomerID).Update("status", customerCourse.Status).Error
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
