package customerCourseMockRepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CustomerCourseMock struct {
	mock.Mock
}

func (c *CustomerCourseMock) DeleteCustomerCourse(id string) error {
	args := c.Called(id)

	return args.Error(0)
}

func (c *CustomerCourseMock) GetCustomerCourse(courseID, customerID string) (dto.CustomerCourse, error) {
	args := c.Called(courseID, customerID)

	return args.Get(0).(dto.CustomerCourse), args.Error(1)
}

func (c *CustomerCourseMock) GetHistoryCourseByCustomerID(customerID string) ([]dto.Course, error) {
	args := c.Called(customerID)

	return args.Get(0).([]dto.Course), args.Error(1)
}

func (c *CustomerCourseMock) TakeCourse(customerCourse dto.CustomerCourseTransaction) error {
	args := c.Called(customerCourse)

	return args.Error(0)
}

func (c *CustomerCourseMock) UpdateEnrollmentStatus(customerCourse dto.CustomerCourseTransaction) error {
	args := c.Called(customerCourse)

	return args.Error(0)
}
