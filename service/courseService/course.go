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
	CreateCourse(dto.CourseTransaction, dto.User) error
	DeleteCourse(id, instructorId string) error
	GetAllCourse(dto.User) ([]dto.Course, error)
	GetCourseByID(id string) (dto.Course, error)
	UpdateCourse(dto.CourseTransaction) error
	GetRatingCourse(dto.Course) float64
}

type courseService struct {
	courseRepo   courseRepository.CourseRepository
	categoryRepo categoryRepository.CategoryRepository
}

// CreateCourse implements CourseService
func (cs *courseService) CreateCourse(course dto.CourseTransaction, user dto.User) error {
	// check if category is not found
	_, err := cs.categoryRepo.GetCategoryByID(course.CategoryID, user)
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
	course, err := cs.courseRepo.GetCourseByID(id)
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
func (cs *courseService) GetAllCourse(user dto.User) ([]dto.Course, error) {
	courses, err := cs.courseRepo.GetAllCourse(user)
	if err != nil {
		return nil, err
	}
	// get rating of all courses
	for i, course := range courses {
		rating := cs.GetRatingCourse(course)
		courses[i].Rating = rating
	}
	return courses, nil
}

// GetCourseByID implements CourseService
func (cs *courseService) GetCourseByID(id string) (dto.Course, error) {
	course, err := cs.courseRepo.GetCourseByID(id)
	if err != nil {
		return dto.Course{}, err
	}

	// get rating of course
	rating := cs.GetRatingCourse(course)
	course.Rating = rating

	return course, nil
}

// UpdateCourse implements CourseService
func (cs *courseService) UpdateCourse(course dto.CourseTransaction) error {
	// check if instructor id is not same
	oldCourse, err := cs.courseRepo.GetCourseByID(course.ID)
	if err != nil {
		return err
	}

	// check if instructor id in the course is the same as the instructor id in the token
	if oldCourse.InstructorID != course.InstructorID {
		return errors.New(constantError.ErrorNotAuthorized)
	}

	// call repository to update course
	err = cs.courseRepo.UpdateCourse(course)
	if err != nil {
		return err
	}
	return nil
}

// RatingCourse implements CourseService
func (cs *courseService) GetRatingCourse(course dto.Course) float64 {
	if len(course.Ratings) == 0 {
		return 0
	}
	// get rating of course
	for _, rating := range course.Ratings {
		course.Rating += float64(rating.Rating)
	}
	// average rating
	average := course.Rating / float64(len(course.Ratings))
	return average
}

func NewCourseService(courseRepo courseRepository.CourseRepository, categoryRepo categoryRepository.CategoryRepository) CourseService {
	return &courseService{
		courseRepo:   courseRepo,
		categoryRepo: categoryRepo,
	}
}
