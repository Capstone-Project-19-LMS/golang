package courseService

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/repository/categoryRepository"
	"golang/repository/courseRepository"
)

type CourseService interface {
	CreateCourse(dto.CourseTransaction) error
	DeleteCourse(id, instructorId string) error
	GetAllCourse(instructorId string) ([]dto.CourseTransaction, error)
	GetCourseByID(id, instructorId string) (dto.Course, error)
	UpdateCourse(dto.CourseTransaction) error
}

type courseService struct {
	courseRepo courseRepository.CourseRepository
	categoryRepo categoryRepository.CategoryRepository
}

// CreateCourse implements CourseService
func (cs *courseService) CreateCourse(course dto.CourseTransaction) error {
	// check if category is not found
	_, err := cs.categoryRepo.GetCategoryByID(course.CategoryID)
	if err != nil {
		return errors.New(constantError.ErrorCategoryNotFound)
	}

	// create uuid for course
	id := helper.GenerateUUID()
	course.ID = id

	// default thumbnail course
	if course.Thumbnail == "" {
		course.Thumbnail = "https://via.placeholder.com/150x100"
	}

	// call repository to create course
	err = cs.courseRepo.CreateCourse(course)
	if err != nil {
		return err
	}
	return nil
}

func (cs *courseService) DeleteCourse(id string, instructorId string) error {
	// check if instructor id is not same
	course, err := cs.courseRepo.GetCourseByID(id, instructorId)
	if err != nil {
		return err
	}

	// check if instructor id in the course is the same as the instructor id in the token
	if course.InstructorID != instructorId {
		return errors.New(constantError.ErrorNotAuthorized)
	}

	// call repository to delete account
	err = cs.courseRepo.DeleteCourse(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCourse implements CourseService
func (*courseService) GetAllCourse(instructorId string) ([]dto.CourseTransaction, error) {
	panic("unimplemented")
}

// GetCourseByID implements CourseService
func (cs *courseService) GetCourseByID(id, instructorId string) (dto.Course, error) {
	course, err := cs.courseRepo.GetCourseByID(id, instructorId)
	if err != nil {
		return dto.Course{}, err
	}
	return course, nil
}

// UpdateCourse implements CourseService
func (*courseService) UpdateCourse(dto.CourseTransaction) error {
	panic("unimplemented")
}

func NewCourseService(courseRepo courseRepository.CourseRepository, categoryRepo categoryRepository.CategoryRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
		categoryRepo: categoryRepo,
	}
}
