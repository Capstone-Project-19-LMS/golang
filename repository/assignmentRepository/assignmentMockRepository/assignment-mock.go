package assignmentmockrepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type AssignmentMock struct {
	mock.Mock
}

func (c *AssignmentMock) CreateAssignment(assignment dto.AssignmentTransaction) error {
	args := c.Called(assignment)

	return args.Error(0)
}

func (c *AssignmentMock) DeleteAssignment(id string) error {
	args := c.Called(id)

	return args.Error(0)
}

func (c *AssignmentMock) GetAssignmentByID(id string) (dto.Assignment, error) {
	args := c.Called(id)

	return args.Get(0).(dto.Assignment), args.Error(1)
}

func (c *AssignmentMock) GetAllAssignment() ([]dto.Assignment, error) {
	args := c.Called()

	return args.Get(0).([]dto.Assignment), args.Error(1)
}

func (c *AssignmentMock) UpdateAssignment(assignment dto.AssignmentTransaction) error {
	args := c.Called(assignment)

	return args.Error(0)
}
