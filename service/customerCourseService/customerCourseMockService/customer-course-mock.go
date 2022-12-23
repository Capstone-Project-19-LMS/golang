package customerCourseMockService

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CustomerCourseMock struct {
	mock.Mock
}

func (c *CustomerCourseMock) DeleteCustomerCourse(courseID, customerID string) error {
	args := c.Called(courseID, customerID)

	return args.Error(0)
}

func (c *CustomerCourseMock) GetHistoryCourseByCustomerID(customerID string) ([]dto.GetCourse, error) {
	args := c.Called(customerID)

	return args.Get(0).([]dto.GetCourse), args.Error(1)
}

func (c *CustomerCourseMock) GetCustomerCourseEnrollByID(id string) (dto.CustomerCourseEnroll, error) {
	args := c.Called(id)

	return args.Get(0).(dto.CustomerCourseEnroll), args.Error(1)
}

func (c *CustomerCourseMock) TakeCourse(customerCourse dto.CustomerCourseTransaction) error {
	args := c.Called(customerCourse)

	return args.Error(0)
}

func (c *CustomerCourseMock) UpdateEnrollmentStatus(customerCourse dto.CustomerCourseTransaction, instructorID string) error {
	args := c.Called(customerCourse, instructorID)

	return args.Error(0)
}
