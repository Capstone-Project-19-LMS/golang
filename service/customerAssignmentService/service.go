package customerAssignmentService

import (
	"golang/helper"
	"golang/models/dto"
	customerAssignmentrepository "golang/repository/customerAssignmentRepository"
)

type CustomerAssignmentService interface {
	CreateCustomerAssignment(dto.CustomerAssignmentTransaction) error
	DeleteCustomerAssignment(id string) error
	GetAllCustomerAssignment() ([]dto.CustomerAssignment, error)
	GetCustomerAssignmentByID(id string) (dto.CustomerAssignment, error)
	UpdateCustomerAssignment(dto.CustomerAssignmentTransaction) error
}

type customerAssignmentService struct {
	customerAssignmentRepo customerAssignmentrepository.CustomerAssignmentRepository
}

// CreateCustomerAssignment implements CustomerAssignmentService
func (cas *customerAssignmentService) CreateCustomerAssignment(customerAssignment dto.CustomerAssignmentTransaction) error {
	id := helper.GenerateUUID()
	customerAssignment.ID = id
	err := cas.customerAssignmentRepo.CreateCustomerAssignment(customerAssignment)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCustomerAssignment implements CustomerAssignmentService
func (cas *customerAssignmentService) DeleteCustomerAssignment(id string) error {
	// call repository to delete account
	err := cas.customerAssignmentRepo.DeleteCustomerAssignment(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCustomerAssignment implements CustomerAssignmentService
func (cas *customerAssignmentService) GetAllCustomerAssignment() ([]dto.CustomerAssignment, error) {
	customerAssignments, err := cas.customerAssignmentRepo.GetAllCustomerAssignment()
	if err != nil {
		return nil, err
	}
	return customerAssignments, nil
}

// GetCustomerAssignmentByID implements CustomerAssignmentService
func (cas *customerAssignmentService) GetCustomerAssignmentByID(id string) (dto.CustomerAssignment, error) {
	customerAssignment, err := cas.customerAssignmentRepo.GetCustomerAssignmentByID(id)
	if err != nil {
		return dto.CustomerAssignment{}, err
	}
	return customerAssignment, nil
}

// UpdateCustomerAssignment implements CustomerAssignmentService
func (cas *customerAssignmentService) UpdateCustomerAssignment(customerAssignment dto.CustomerAssignmentTransaction) error {
	// call repository to update CustomerAssignment
	err := cas.customerAssignmentRepo.UpdateCustomerAssignment(customerAssignment)
	if err != nil {
		return err
	}
	return nil
}

func NewcustomerAssignmentService(customerAssignmentRepo customerAssignmentrepository.CustomerAssignmentRepository) CustomerAssignmentService {
	return &customerAssignmentService{
		customerAssignmentRepo: customerAssignmentRepo,
	}
}
