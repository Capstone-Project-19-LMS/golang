package ratingRepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ratingRepository struct {
	db *gorm.DB
}

// AddRating implements RatingRepository
func (rr *ratingRepository) AddRating(rating dto.RatingTransaction) error {
	var ratingModel model.Rating
	err := copier.Copy(&ratingModel, &rating)
	if err != nil {
		return err
	}
	// save rating to database
	err = rr.db.Model(&model.Rating{}).Create(&ratingModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteRating implements RatingRepository
func (rr *ratingRepository) DeleteRating(id string) error {
	err := rr.db.Unscoped().Delete(&model.Rating{}, "id = ?", id)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

// GetRatingByCourseID implements RatingRepository
func (rr *ratingRepository) GetRatingByCourseID(courseID string) ([]dto.Rating, error) {
	var ratingModels []model.Rating

	// get data rating from database by course id
	err := rr.db.Model(&model.Rating{}).Where("course_id = ?", courseID).Find(&ratingModels).Error
	if err != nil {
		return nil, err
	}

	var ratings []dto.Rating
	err = copier.Copy(&ratings, &ratingModels)
	if err != nil {
		return nil, err
	}

	return ratings, err
}

// GetRating implements RatingRepository
func (rr *ratingRepository) GetRatingByCourseIDCustomerID(courseID string, customerID string) (dto.Rating, error) {
	var rating dto.Rating
	err := rr.db.Model(&model.Rating{}).Where("course_id = ? AND customer_id = ?", courseID, customerID).First(&rating)
	if err.Error != nil {
		return dto.Rating{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Rating{}, gorm.ErrRecordNotFound
	}
	return rating, nil
}

// GetRatingByID implements RatingRepository
func (rr *ratingRepository) GetRatingByID(id string) (dto.Rating, error) {
	var rating dto.Rating
	err := rr.db.Model(&model.Rating{}).Where("id = ?", id).First(&rating)
	if err.Error != nil {
		return dto.Rating{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Rating{}, gorm.ErrRecordNotFound
	}
	return rating, nil
}

// UpdateRating implements RatingRepository
func (rr *ratingRepository) UpdateRating(rating dto.RatingTransaction) error {
	// update rating to database
	err := rr.db.Model(&model.Rating{}).Where("id = ?", rating.ID).Update("is_publish", rating.IsPublish).Error
	if err != nil {
		return err
	}

	return nil

}

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{
		db: db,
	}
}
