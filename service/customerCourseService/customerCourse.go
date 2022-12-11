package customerCourseService

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/repository/courseRepository"
	"golang/repository/customerCourseRepository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CustomerCourseService interface {
	DeleteCustomerCourse(courseID, customerID string) error
	GetHistoryCourseByCustomerID(string) ([]dto.GetCourse, error)
	TakeCourse(dto.CustomerCourseTransaction) error
	UpdateEnrollmentStatus(customerCourse dto.CustomerCourseTransaction, instructorID string) error
}

type customerCourseService struct {
	courseRepo         courseRepository.CourseRepository
	customerCourseRepo customerCourseRepository.CustomerCourseRepository
}

// DeleteCustomerCourse implements CustomerCourseService
func (ccs *customerCourseService) DeleteCustomerCourse(courseID, customerID string) error {
	// get customer course by id
	customerCourse, err := ccs.customerCourseRepo.GetCustomerCourse(courseID, customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constantError.ErrorCustomerNotEnrolled)
		}
		return err
	}

	// check if customer course is not belong to customer
	if customerCourse.CustomerID != customerID {
		return errors.New(constantError.ErrorNotAuthorized)
	}

	// delete customer course
	err = ccs.customerCourseRepo.DeleteCustomerCourse(customerCourse.ID)
	if err != nil {
		return err
	}

	// get course by id
	course, err := ccs.courseRepo.GetCourseByID(customerCourse.CourseID)
	if err != nil {
		return err
	}

	// update capacity course
	courseUpdate := dto.CourseTransaction{
		ID:       customerCourse.CourseID,
		Capacity: course.Capacity + 1,
	}
	err = ccs.courseRepo.UpdateCourse(courseUpdate)
	if err != nil {
		return err
	}

	return nil
}

// GetHistoryCourseByCustomerID implements CustomerCourseService
func (ccs *customerCourseService) GetHistoryCourseByCustomerID(customerID string) ([]dto.GetCourse, error) {
	courses, err := ccs.customerCourseRepo.GetHistoryCourseByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	for i, course := range courses {
		// get rating of all courses
		rating := helper.GetRatingCourse(course)
		courses[i].Rating = rating

		// get favorite of all courses
		favorite := helper.GetFavoriteCourse(course, customerID)
		courses[i].Favorite = favorite

		// get number of module
		numberOfModule := len(course.Modules)
		courses[i].NumberOfModules = numberOfModule

		// get status enroll of all courses
		courses[i].StatusEnroll = course.CustomerCourses[0].Status
	}

	// copy courses from dto.course to dto.GetCustomerCourse
	var customerCourses []dto.GetCourse
	err = copier.Copy(&customerCourses, &courses)
	if err != nil {
		return nil, err
	}
	return customerCourses, nil
}

// TakeCourse implements CustomerCourseService
func (ccs *customerCourseService) TakeCourse(customerCourse dto.CustomerCourseTransaction) error {
	// check if course is not found
	course, err := ccs.courseRepo.GetCourseByID(customerCourse.CourseID)
	if err != nil {
		return err
	}
	// check capacity course
	if course.Capacity == 0 {
		return errors.New(constantError.ErrorCourseCapacity)
	}

	// check if customer already take the course
	_, err = ccs.customerCourseRepo.GetCustomerCourse(customerCourse.CourseID, customerCourse.CustomerID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(constantError.ErrorCustomerAlreadyTakeCourse)
	}

	// create uuid for customer course
	id := helper.GenerateUUID()
	customerCourse.ID = id

	// set no module
	customerCourse.NoModule = 1

	// call repository to take the course
	err = ccs.customerCourseRepo.TakeCourse(customerCourse)
	if err != nil {
		return err
	}

	// update capacity course
	var courseUpdate dto.CourseTransaction
	err = copier.Copy(&courseUpdate, &course)
	if err != nil {
		return err
	}
	courseUpdate.Capacity -= 1
	err = ccs.courseRepo.UpdateCourse(courseUpdate)
	if err != nil {
		return err
	}

	return nil
}

// UpdateEnrollmentStatus implements CustomerCourseService
func (ccs *customerCourseService) UpdateEnrollmentStatus(customerCourse dto.CustomerCourseTransaction, instructorID string) error {
	// check if course is not belong to instructor
	course, err := ccs.courseRepo.GetCourseByID(customerCourse.CourseID)
	if err != nil {
		return err
	}

	if course.InstructorID != instructorID {
		return errors.New(constantError.ErrorNotAuthorized)
	}
	
	// get data enrollment course
	_, err = ccs.customerCourseRepo.GetCustomerCourse(customerCourse.CourseID, customerCourse.CustomerID)
	// check if enrollment course is not found
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constantError.ErrorCustomerNotEnrolled)
		}
		return err
	}
	
	// update enrollment status
	err = ccs.customerCourseRepo.UpdateEnrollmentStatus(customerCourse)
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerCourseService(customerCourseRepo customerCourseRepository.CustomerCourseRepository,
	courseRepo courseRepository.CourseRepository) CustomerCourseService {
	return &customerCourseService{
		customerCourseRepo: customerCourseRepo,
		courseRepo:         courseRepo,
	}
}
