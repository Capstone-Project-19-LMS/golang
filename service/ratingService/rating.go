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
	GetRatingByCourseID(courseID string) ([]dto.Rating, error)
	GetRatingByCourseIDCustomerID(courseID, customerID string) (dto.Rating, error)
}

type ratingService struct {
	courseRepo courseRepository.CourseRepository
	ratingRepo ratingRepository.RatingRepository
}

// AddRating implements RatingService
func (rs *ratingService) AddRating(rating dto.RatingTransaction) error {
	// check if course is not found
	_, err := rs.courseRepo.GetCourseByID(rating.CourseID)
	if err != nil {
		return err
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
func (*ratingService) GetRatingByCourseID(courseID string) ([]dto.Rating, error) {
	panic("unimplemented")
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

func NewRatingService(ratingRepo ratingRepository.RatingRepository,
	courseRepo courseRepository.CourseRepository) RatingService {
	return &ratingService{
		ratingRepo: ratingRepo,
		courseRepo: courseRepo,
	}
}
