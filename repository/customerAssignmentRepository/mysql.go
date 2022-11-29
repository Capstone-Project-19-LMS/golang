package customerassignmentrepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type customerAssignmentRepository struct {
	db *gorm.DB
}

// CreatecustomerAssignment implements customerAssignmentRepository
func (ctr *customerAssignmentRepository) CreateCustomerAssignment(customerAssignment dto.CustomerAssignmentTransaction) error {
	var customerAssignmentModel model.CustomerAssignment
	err := copier.Copy(&customerAssignmentModel, &customerAssignment)
	if err != nil {
		return err
	}

	err = ctr.db.Model(&model.CustomerAssignment{}).Create(&customerAssignmentModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeletecustomerAssignment implements customerAssignmentRepository
func (ctr *customerAssignmentRepository) DeleteCustomerAssignment(id string) error {
	// delete data customerAssignment from database by id
	err := ctr.db.Where("id = ?", id).Unscoped().Delete(&model.CustomerAssignment{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAllcustomerAssignment implements customerAssignmentRepository
func (ctr *customerAssignmentRepository) GetAllCustomerAssignment() ([]dto.CustomerAssignment, error) {
	var customerAssignmentModels []model.CustomerAssignment
	// get data sub category from database by user
	err := ctr.db.Model(&model.CustomerAssignment{}).Find(&customerAssignmentModels).Error
	if err != nil {
		return nil, err
	}
	// copy data from model to dto
	var customerAssignments []dto.CustomerAssignment
	err = copier.Copy(&customerAssignments, &customerAssignmentModels)
	if err != nil {
		return nil, err
	}
	return customerAssignments, nil
}

// GetcustomerAssignmentByID implements customerAssignmentRepository
func (ctr *customerAssignmentRepository) GetCustomerAssignmentByID(id string) (dto.CustomerAssignment, error) {
	var customerAssignmentModel model.CustomerAssignment
	err := ctr.db.Model(&model.CustomerAssignment{}).Where("id = ?", id).Find(&customerAssignmentModel)
	if err.Error != nil {
		return dto.CustomerAssignment{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.CustomerAssignment{}, gorm.ErrRecordNotFound
	}

	// copy data from model to dto
	var customerAssignment dto.CustomerAssignment
	errCopy := copier.Copy(&customerAssignment, &customerAssignmentModel)
	if errCopy != nil {
		return dto.CustomerAssignment{}, errCopy
	}
	return customerAssignment, nil
}

// UpdatecustomerAssignment implements customerAssignmentRepository
func (ctr *customerAssignmentRepository) UpdateCustomerAssignment(customerAssignment dto.CustomerAssignmentTransaction) error {
	var customerAssignmentModel model.CustomerAssignment
	errCopy := copier.Copy(&customerAssignmentModel, &customerAssignment)
	if errCopy != nil {
		return errCopy
	}
	// update account with new data
	err := ctr.db.Model(&model.CustomerAssignment{}).Where("id = ?", customerAssignment.ID).Updates(&customerAssignmentModel)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewcustomerAssignmentRepository(db *gorm.DB) CustomerAssignmentRepository {
	return &customerAssignmentRepository{
		db: db,
	}
}
