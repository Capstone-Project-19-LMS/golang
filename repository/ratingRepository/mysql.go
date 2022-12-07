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
func (*ratingRepository) DeleteRating(id string) error {
	panic("unimplemented")
}

// GetRatingByCourseID implements RatingRepository
func (*ratingRepository) GetRatingByCourseID(courseID string) ([]dto.Rating, error) {
	panic("unimplemented")
}

// GetRating implements RatingRepository
func (rr *ratingRepository) GetRating(courseID string, customerID string) (dto.Rating, error) {
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

func NewRatingRepository(db *gorm.DB) RatingRepository {
	return &ratingRepository{
		db: db,
	}
}
