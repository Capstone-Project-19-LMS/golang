package courseRepository

import "golang/models/dto"

type CourseRepository interface {
	CreateCourse(dto.CourseTransaction) error
	DeleteCourse(string) error
	GetCourseByID(string) (dto.Course, error)
	GetAllCourse(dto.User) ([]dto.Course, error)
	UpdateCourse(dto.CourseTransaction) error
}