package favoriteRepository

import "golang/models/dto"

type FavoriteRepository interface {
	AddFavorite(favorite dto.FavoriteTransaction) error
	DeleteFavorite(id string) error
	GetFavoriteByCustomerID(customerID string) ([]dto.Course, error)
	GetFavorite(courseID, customerID string) (dto.Favorite, error)
}