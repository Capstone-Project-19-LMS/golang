package courseService

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/repository/categoryRepository"
	"golang/repository/courseRepository"

	"github.com/jinzhu/copier"
)

type CourseService interface {
	CreateCourse(dto.CourseTransaction, dto.User) error
	DeleteCourse(id, instructorId string) error
	GetAllCourse(dto.User) ([]dto.GetCourse, error)
	GetCourseByID(id string, user dto.User) (dto.Course, error)
	GetCourseEnrollByID(string) ([]dto.CustomerEnroll, error)
	UpdateCourse(dto.CourseTransaction) error
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

	course.InstructorID = user.ID

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
func (cs *courseService) GetAllCourse(user dto.User) ([]dto.GetCourse, error) {
	courses, err := cs.courseRepo.GetAllCourse(user)
	if err != nil {
		return nil, err
	}
	// check if courses is empty
	if len(courses) == 0 {
		return []dto.GetCourse{}, nil
	}

	// get rating of all courses
	for i, course := range courses {
		rating := helper.GetRatingCourse(course)
		courses[i].Rating = rating
	}

	// get number of module
	for i, course := range courses {
		numberOfModule := len(course.Modules)
		courses[i].NumberOfModules = numberOfModule
	}

	if user.Role == "customer" {
		// get favorite of all courses
		for i, course := range courses {
			favorite := helper.GetFavoriteCourse(course, user.ID)
			courses[i].Favorite = favorite
		}

		// get enrolled of all courses
		for i, course := range courses {
			helper.GetEnrolledCourse(&course, user.ID)
			courses[i].Enroll = course.Enroll
		}
	}
	var getCourses []dto.GetCourse
	err = copier.Copy(&getCourses, &courses)
	if err != nil {
		return nil, err
	}

	return getCourses, nil
}

// GetCourseByID implements CourseService
func (cs *courseService) GetCourseByID(id string, user dto.User) (dto.Course, error) {
	course, err := cs.courseRepo.GetCourseByID(id)
	if err != nil {
		return dto.Course{}, err
	}

	// get rating of course
	rating := helper.GetRatingCourse(course)
	course.Rating = rating

	// get number of module
	numberOfModule := len(course.Modules)
	course.NumberOfModules = numberOfModule

	if user.Role == "customer" {
		// get favorites of course
		favorite := helper.GetFavoriteCourse(course, user.ID)
		course.Favorite = favorite

		// get enrolled of course
		helper.GetEnrolledCourse(&course, user.ID)
	}

	return course, nil
}

// GetCourseEnrollByID implements CourseService
func (cs *courseService) GetCourseEnrollByID(id string) ([]dto.CustomerEnroll, error) {
	course, err := cs.courseRepo.GetCourseEnrollByID(id)
	if err != nil {
		return nil, err
	}
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

func NewCourseService(courseRepo courseRepository.CourseRepository, categoryRepo categoryRepository.CategoryRepository) CourseService {
	return &courseService{
		courseRepo:   courseRepo,
		categoryRepo: categoryRepo,
	}
}
