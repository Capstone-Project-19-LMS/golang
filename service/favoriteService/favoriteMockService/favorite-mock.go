package favoriteMockService

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type FavoriteMock struct {
	mock.Mock
}

func (f *FavoriteMock) AddFavorite(favorite dto.FavoriteTransaction) error {
	args := f.Called(favorite)

	return args.Error(0)
}

func (f *FavoriteMock) DeleteFavorite(courseID, customerID string) error {
	args := f.Called(courseID, customerID)

	return args.Error(0)
}

func (f *FavoriteMock) GetFavoriteByCustomerID(customerID string) ([]dto.GetCourse, error) {
	args := f.Called(customerID)

	return args.Get(0).([]dto.GetCourse), args.Error(1)
}
