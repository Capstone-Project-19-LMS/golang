package courseMockService

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CourseMock struct {
	mock.Mock
}

func (c *CourseMock)CreateCourse(course dto.CourseTransaction, user dto.User) error {
	args := c.Called(course, user)

	return args.Error(0)
}
func (c *CourseMock)DeleteCourse(id, instructorId string) error {
	args := c.Called(id, instructorId)

	return args.Error(0)
}
func (c *CourseMock)GetAllCourse(user dto.User) ([]dto.GetCourse, error) {
	args := c.Called(user)

	return args.Get(0).([]dto.GetCourse), args.Error(1)
}
func (c *CourseMock)GetCourseByID(id string, user dto.User) (dto.GetCourseByID, error) {
	args := c.Called(id, user)

	return args.Get(0).(dto.GetCourseByID), args.Error(1)
}
func (c *CourseMock)GetCourseEnrollByID(id string, user dto.User) ([]dto.CustomerCourseEnroll, error) {
	args := c.Called(id, user)

	return args.Get(0).([]dto.CustomerCourseEnroll), args.Error(1)
}
func (c *CourseMock)UpdateCourse(course dto.CourseTransaction) error {
	args := c.Called(course)

	return args.Error(0)
}