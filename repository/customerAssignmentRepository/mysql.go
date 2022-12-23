package customerassignmentrepository

import (
	"errors"
	"fmt"
	"golang/constant/constantError"
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

	var getAllCustomerAssignment []model.CustomerAssignment

	ctr.db.Find(&getAllCustomerAssignment)

	for _, gaca := range getAllCustomerAssignment {
		if customerAssignment.CustomerID == gaca.CustomerID {
			if customerAssignment.AssignmentID == gaca.AssignmentID {
				return errors.New(constantError.ErrorDuplicateAssignmentCustomer)
			}
		}
	}

	result := ctr.db.Model(&model.CustomerAssignment{}).Create(&customerAssignmentModel)
	var storage dto.CustomerAssignment = dto.CustomerAssignment{}
	result.Last(&storage)
	fmt.Println(storage)

	var getAssignment model.Assignment
	ctr.db.Where("id=?", storage.AssignmentID).Find(&getAssignment)

	var getModule model.Module
	ctr.db.Where("id=?", getAssignment.ModuleID).Find(&getModule)

	var updateCustomerCourse model.CustomerCourse

	ctr.db.Where("customer_id =?", storage.CustomerID).Where("course_id=?", getModule.CourseID).Find(&updateCustomerCourse)

	updateCustomerCourse.NoModule = updateCustomerCourse.NoModule + 1

	var getAllModule []model.Module
	ctr.db.Where("course_id=?", getModule.CourseID).Find(&getAllModule)

	// check if customer course is finished
	if updateCustomerCourse.NoModule > len(getAllModule) {
		updateCustomerCourse.IsFinish = true
	}
	for _, gam := range getAllModule {
		if updateCustomerCourse.NoModule >= gam.NoModule {
			ctr.db.Save(&updateCustomerCourse)
		}
	}

	// fmt.Println(updateCustomerCourse)

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
func (ctr *customerAssignmentRepository) GetAllCustomerAssignment() ([]dto.CustomerAssignmentAcc, error) {
	var customerAssignmentModels []model.CustomerAssignment
	// get data sub category from database by user
	err := ctr.db.Model(&model.CustomerAssignment{}).Find(&customerAssignmentModels).Error
	if err != nil {
		return nil, err
	}

	// copy data from model to dto
	var customerAssignments []dto.CustomerAssignmentAcc
	var customer []model.Customer
	err = copier.Copy(&customerAssignments, &customerAssignmentModels)
	if err != nil {
		return nil, err
	}
	var result []dto.CustomerAssignmentAcc
	for _, cam := range customerAssignments {
		ctr.db.Where("id", cam.CustomerID).Find(&customer)

		for _, cs := range customer {
			cam.Customer.Name = cs.Name
			fmt.Println(cam)

		}
		result = append(result, cam)
	}

	return result, nil
}

// GetcustomerAssignmentByID implements customerAssignmentRepository
func (ctr *customerAssignmentRepository) GetCustomerAssignmentByID(id string) (dto.CustomerAssignmentAcc, error) {
	var customerAssignmentModel model.CustomerAssignment
	err := ctr.db.Model(&model.CustomerAssignment{}).Where("id = ?", id).Find(&customerAssignmentModel)
	if err.Error != nil {
		return dto.CustomerAssignmentAcc{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.CustomerAssignmentAcc{}, gorm.ErrRecordNotFound
	}

	// copy data from model to dto
	var customerAssignment dto.CustomerAssignmentAcc

	var customer model.Customer
	ctr.db.Where("id=?", customerAssignmentModel.CustomerID).Find(&customer)
	customerAssignment.Customer.Name = customer.Name
	errCopy := copier.Copy(&customerAssignment, &customerAssignmentModel)
	if errCopy != nil {
		return dto.CustomerAssignmentAcc{}, errCopy
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
