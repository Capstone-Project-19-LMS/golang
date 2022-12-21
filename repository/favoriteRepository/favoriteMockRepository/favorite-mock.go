package favoriteMockRepository

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
func (f *FavoriteMock) DeleteFavorite(id string) error {
	args := f.Called(id)

	return args.Error(0)
}
func (f *FavoriteMock) GetFavoriteByCustomerID(customerID string) ([]dto.Course, error) {
	args := f.Called(customerID)

	return args.Get(0).([]dto.Course), args.Error(1)
}
func (f *FavoriteMock) GetFavorite(courseID, customerID string) (dto.Favorite, error) {
	args := f.Called(courseID, customerID)

	return args.Get(0).(dto.Favorite), args.Error(1)
}