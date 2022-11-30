package customerCourseRepository

import "golang/models/dto"

type CustomerCourseRepository interface {
	GetCustomerCourse(courseID, customerID string) (dto.CustomerCourse, error)
	GetHistoryCourseByCustomerID(string) ([]dto.Course, error)
	TakeCourse(dto.CustomerCourseTransaction) error
}