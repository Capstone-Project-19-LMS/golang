package customerassignmentrepository

import "golang/models/dto"

type CustomerAssignmentRepository interface {
	CreateCustomerAssignment(dto.CustomerAssignmentTransaction) error
	DeleteCustomerAssignment(id string) error
	GetAllCustomerAssignment() ([]dto.CustomerAssignment, error)
	GetCustomerAssignmentByID(id string) (dto.CustomerAssignment, error)
	UpdateCustomerAssignment(dto.CustomerAssignmentTransaction) error
}
