package courseMockRepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CourseMock struct {
	mock.Mock
}

func (c *CourseMock) CreateCourse(course dto.CourseTransaction) error {
	args := c.Called(course)

	return args.Error(0)
}
func (c *CourseMock) DeleteCourse(id string) error {
	args := c.Called(id)

	return args.Error(0)
}
func (c *CourseMock) GetCourseByID(id string) (dto.Course, error) {
	args := c.Called(id)

	return args.Get(0).(dto.Course), args.Error(1)
}
func (c *CourseMock) GetCourseEnrollByID(id string) ([]dto.CustomerCourseEnroll, error) {
	args := c.Called(id)

	return args.Get(0).([]dto.CustomerCourseEnroll), args.Error(1)
}
func (c *CourseMock) GetAllCourse(user dto.User) ([]dto.Course, error) {
	args := c.Called(user)

	return args.Get(0).([]dto.Course), args.Error(1)
}
func (c *CourseMock) UpdateCourse(course dto.CourseTransaction) error {
	args := c.Called(course)

	return args.Error(0)
}