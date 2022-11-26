package courseRepository

import "golang/models/dto"

type CourseRepository interface {
	CreateCourse(dto.CourseTransaction) error
	DeleteCourse(id string) error
	GetCourseByID(id string) (dto.Course, error)
	GetAllCourse(instructorId string) ([]dto.Course, error)
	UpdateCourse(dto.CourseTransaction) error
}