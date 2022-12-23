package customerassignmentrepository

import "golang/models/dto"

type CustomerAssignmentRepository interface {
	CreateCustomerAssignment(dto.CustomerAssignmentTransaction) error
	DeleteCustomerAssignment(id string) error
	GetAllCustomerAssignment() ([]dto.CustomerAssignmentAcc, error)
	GetCustomerAssignmentByID(id string) (dto.CustomerAssignmentAcc, error)
	UpdateCustomerAssignment(dto.CustomerAssignmentTransaction) error
}
