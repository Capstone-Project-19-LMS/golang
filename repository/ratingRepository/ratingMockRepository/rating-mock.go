package ratingMockRepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type RatingMock struct {
	mock.Mock
}

func (r *RatingMock) AddRating(rating dto.RatingTransaction) error {
	args := r.Called(rating)

	return args.Error(0)
}
func (r *RatingMock) DeleteRating(id string) error {
	args := r.Called(id)

	return args.Error(0)
}
func (r *RatingMock) GetRatingByCourseID(courseID string) ([]dto.Rating, error) {
	args := r.Called(courseID)

	return args.Get(0).([]dto.Rating), args.Error(1)
}
func (r *RatingMock) GetRatingByCourseIDCustomerID(courseID, customerID string) (dto.Rating, error) {
	args := r.Called(courseID, customerID)

	return args.Get(0).(dto.Rating), args.Error(1)
}
func (r *RatingMock) GetRatingByID(id string) (dto.Rating, error) {
	args := r.Called(id)

	return args.Get(0).(dto.Rating), args.Error(1)
}
func (r *RatingMock) UpdateRating(rating dto.RatingTransaction) error {
	args := r.Called(rating)

	return args.Error(0)
}