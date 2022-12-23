package ratingService

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/repository/courseRepository"
	"golang/repository/ratingRepository"

	"gorm.io/gorm"
)

type RatingService interface {
	AddRating(rating dto.RatingTransaction) error
	DeleteRating(courseID, customerID string) error
	GetRatingByCourseID(courseID, InstructorID string) ([]dto.Rating, error)
	GetRatingByCourseIDCustomerID(courseID, customerID string) (dto.Rating, error)
	UpdateRating(rating dto.RatingTransaction) error
}

type ratingService struct {
	courseRepo courseRepository.CourseRepository
	ratingRepo ratingRepository.RatingRepository
}

// AddRating implements RatingService
func (rs *ratingService) AddRating(rating dto.RatingTransaction) error {
	// check if course is not found
	course, err := rs.courseRepo.GetCourseByID(rating.CourseID)
	if err != nil {
		return err
	}

	// check if customer already finished the course
	helper.GetEnrolledCourse(&course, rating.CustomerID)
	if !course.IsFinish {
		return errors.New(constantError.ErrorCustomerNotFinishedCourse)
	}


	// check if customer already review the course
	_, err = rs.ratingRepo.GetRatingByCourseIDCustomerID(rating.CourseID, rating.CustomerID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(constantError.ErrorCustomerAlreadyRatingCourse)
	}

	// create uuid for rating
	id := helper.GenerateUUID()
	rating.ID = id

	// call repository to create rating
	err = rs.ratingRepo.AddRating(rating)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRating implements RatingService
func (rs *ratingService) DeleteRating(courseID, customerID string) error {
	// get customer course by id
	rating, err := rs.ratingRepo.GetRatingByCourseIDCustomerID(courseID, customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constantError.ErrorCustomerNotRatingCourse)
		}
		return err
	}

	// check if rating is not belong to customer
	if rating.CustomerID != customerID {
		return errors.New(constantError.ErrorNotAuthorized)
	}

	// delete rating course
	err = rs.ratingRepo.DeleteRating(rating.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetRatingByCourseID implements RatingService
func (rs *ratingService) GetRatingByCourseID(courseID, instructorID string) ([]dto.Rating, error) {
	// get course and check if it exists
	course, err := rs.courseRepo.GetCourseByID(courseID)
	if err != nil {
		return nil, err
	}

	// check if instructor id in the course is the same as the instructor id in the token
	if course.InstructorID != instructorID {
		return nil, errors.New(constantError.ErrorNotAuthorized)
	}

	var ratings []dto.Rating
	ratings, errRating := rs.ratingRepo.GetRatingByCourseID(courseID)
	if errRating != nil {
		return ratings, errRating
	}

	return ratings, nil
}

// GetRating implements RatingService
func (rs *ratingService) GetRatingByCourseIDCustomerID(courseID string, customerID string) (dto.Rating, error) {
	var rating dto.Rating
	rating, err := rs.ratingRepo.GetRatingByCourseIDCustomerID(courseID, customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rating, errors.New(constantError.ErrorCustomerNotRatingCourse)
		}
		return rating, err
	}

	return rating, nil
}

// UpdateRating implements RatingService
func (rs *ratingService) UpdateRating(rating dto.RatingTransaction) error {
	// call repository to update rating
	err := rs.ratingRepo.UpdateRating(rating)
	if err != nil {
		return err
	}

	return nil
}

func NewRatingService(ratingRepo ratingRepository.RatingRepository,
	courseRepo courseRepository.CourseRepository) RatingService {
	return &ratingService{
		ratingRepo: ratingRepo,
		courseRepo: courseRepo,
	}
}
