package ratingRepository

import "golang/models/dto"

type RatingRepository interface {
	AddRating(rating dto.RatingTransaction) error
	DeleteRating(id string) error
	GetRatingByCourseID(courseID string) ([]dto.Rating, error)
	GetRatingByCourseIDCustomerID(courseID, customerID string) (dto.Rating, error)
}