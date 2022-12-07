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
	DeleteRating(id string) error
	GetRatingByCourseID(courseID string) ([]dto.Rating, error)
	// GetRating(courseID, customerID string) (dto.Rating, error)
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
	_, err = rs.ratingRepo.GetRating(rating.CourseID, rating.CustomerID)
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
func (*ratingService) DeleteRating(id string) error {
	panic("unimplemented")
}

// GetRatingByCourseID implements RatingService
func (*ratingService) GetRatingByCourseID(courseID string) ([]dto.Rating, error) {
	panic("unimplemented")
}

func NewRatingService(ratingRepo ratingRepository.RatingRepository,
	courseRepo courseRepository.CourseRepository) RatingService {
	return &ratingService{
		ratingRepo: ratingRepo,
		courseRepo: courseRepo,
	}
}
