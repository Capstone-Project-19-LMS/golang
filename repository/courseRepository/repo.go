package courseRepository

import "golang/models/dto"

type CourseRepository interface {
	CreateCourse(dto.CourseTransaction) error
	DeleteCourse(id string) error
	GetCourseByID(id, instructorId string) (dto.Course, error)
	GetAllCourse(instructorId string) ([]dto.CourseTransaction, error)
	UpdateCourse(dto.CourseTransaction) error
}