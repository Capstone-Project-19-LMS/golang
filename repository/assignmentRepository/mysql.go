package assignmentrepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type assignmentRepository struct {
	db *gorm.DB
}

// CreateAssignment implements AssignmentRepository
func (ar *assignmentRepository) CreateAssignment(assignment dto.AssignmentTransaction) error {
	var assignmentModel model.Assignment
	copier.Copy(&assignmentModel, &assignment)

	err := ar.db.Model(&model.Assignment{}).Create(&assignmentModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteAssignment implements AssignmentRepository
func (ar *assignmentRepository) DeleteAssignment(id string) error {
	// delete data Assignment from database by id
	err := ar.db.Where("id = ?", id).Delete(&model.Assignment{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAllAssignment implements AssignmentRepository
func (ar *assignmentRepository) GetAllAssignment() ([]dto.Assignment, error) {
	var assignmentModels []model.Assignment
	// get data sub category from database by user
	err := ar.db.Model(&model.Assignment{}).Find(&assignmentModels).Error
	if err != nil {
		return nil, err
	}
	// copy data from model to dto
	var assignments []dto.Assignment
	copier.Copy(&assignments, &assignmentModels)

	return assignments, nil
}

// GetAssignmentByID implements AssignmentRepository
func (ar *assignmentRepository) GetAssignmentByID(id string) (dto.Assignment, error) {
	var assignmentModel model.Assignment
	err := ar.db.Model(&model.Assignment{}).Where("id = ?", id).Find(&assignmentModel)
	if err.Error != nil {
		return dto.Assignment{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.Assignment{}, gorm.ErrRecordNotFound
	}

	// copy data from model to dto
	var assignment dto.Assignment
	copier.Copy(&assignment, &assignmentModel)

	return assignment, nil
}

// UpdateAssignment implements AssignmentRepository
func (ar *assignmentRepository) UpdateAssignment(assignment dto.AssignmentTransaction) error {
	var assignmentModel model.Assignment
	copier.Copy(&assignmentModel, &assignment)

	// update account with new data
	err := ar.db.Model(&model.Assignment{}).Where("id = ?", assignment.ID).Updates(&assignmentModel)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{
		db: db,
	}
}
