package customerAssignmentMockrepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CustomerAssignmentMock struct {
	mock.Mock
}

func (c *CustomerAssignmentMock) CreateCustomerAssignment(customerAssignment dto.CustomerAssignmentTransaction) error {
	args := c.Called(customerAssignment)

	return args.Error(0)
}

func (c *CustomerAssignmentMock) DeleteCustomerAssignment(id string) error {
	args := c.Called(id)

	return args.Error(0)
}

func (c *CustomerAssignmentMock) GetAllCustomerAssignment() ([]dto.CustomerAssignment, error) {
	args := c.Called()

	return args.Get(0).([]dto.CustomerAssignment), args.Error(1)
}

func (c *CustomerAssignmentMock) GetCustomerAssignmentByID(id string) (dto.CustomerAssignment, error) {
	args := c.Called(id)

	return args.Get(0).(dto.CustomerAssignment), args.Error(1)
}

func (c *CustomerAssignmentMock) UpdateCustomerAssignment(customerAssignment dto.CustomerAssignmentTransaction) error {
	args := c.Called(customerAssignment)

	return args.Error(0)
}
