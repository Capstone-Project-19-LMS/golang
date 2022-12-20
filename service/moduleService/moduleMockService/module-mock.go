package modulemockservice

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type ModuleMock struct {
	mock.Mock
}

func (c *ModuleMock) CreateModule(module dto.ModuleTransaction) error {
	args := c.Called(module)

	return args.Error(0)
}

func (c *ModuleMock) DeleteModule(id string) error {
	args := c.Called(id)

	return args.Error(0)
}

func (c *ModuleMock) GetModuleByID(id, customerID string) (dto.ModuleAcc, error) {
	args := c.Called(id)

	return args.Get(0).(dto.ModuleAcc), args.Error(1)
}

func (c *ModuleMock) GetAllModule() ([]dto.Module, error) {
	args := c.Called()

	return args.Get(0).([]dto.Module), args.Error(1)
}
func (c *ModuleMock) GetModuleByCourseID(courseID, customerID string) ([]dto.Module, error) {
	args := c.Called(courseID, customerID)

	return args.Get(0).([]dto.Module), args.Error(1)
}

func (c *ModuleMock) UpdateModule(module dto.ModuleTransaction) error {
	args := c.Called(module)

	return args.Error(0)
}
