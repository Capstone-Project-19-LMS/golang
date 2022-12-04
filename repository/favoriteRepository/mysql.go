package favoriteRepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type favoriteRepository struct {
	db *gorm.DB
}


// AddFavorite implements FavoriteRepository
func (fr *favoriteRepository) AddFavorite(favorite dto.FavoriteTransaction) error {
	var favoriteModel model.Favorite
	err := copier.Copy(&favoriteModel, &favorite)
	if err != nil {
		return err
	}
	// save customer course to database
	err = fr.db.Model(&model.Favorite{}).Create(&favoriteModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteFavorite implements FavoriteRepository
func (fr *favoriteRepository) DeleteFavorite(id string) error {
	err := fr.db.Unscoped().Delete(&model.Favorite{}, "id = ?", id)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

// GetFavoriteByCustomerID implements FavoriteRepository
func (*favoriteRepository) GetFavoriteByCustomerID(customerID string) ([]dto.Favorite, error) {
	panic("unimplemented")
}

// GetFavorite implements FavoriteRepository
func (fr *favoriteRepository) GetFavorite(courseID string, customerID string) (dto.Favorite, error) {
	var favorite dto.Favorite
	err := fr.db.Model(&model.Favorite{}).Where("course_id = ? AND customer_id = ?", courseID, customerID).First(&favorite)
	if err.Error != nil {
		return dto.Favorite{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Favorite{}, gorm.ErrRecordNotFound
	}
	return favorite, nil
}

func NewFavoriteRepository(db *gorm.DB) FavoriteRepository {
	return &favoriteRepository{
		db: db,
	}
}
