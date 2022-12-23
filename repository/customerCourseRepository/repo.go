package customerCourseRepository

import "golang/models/dto"

type CustomerCourseRepository interface {
	DeleteCustomerCourse(string) error
	GetCustomerCourse(courseID, customerID string) (dto.CustomerCourse, error)
	GetCustomerCourseByID(id string) (dto.CustomerCourse, error)
	GetHistoryCourseByCustomerID(string) ([]dto.Course, error)
	GetCustomerCourseEnrollByID(id string) (dto.CustomerCourseEnroll, error)
	TakeCourse(dto.CustomerCourseTransaction) error
	UpdateEnrollmentStatus(dto.CustomerCourseTransaction) error
}