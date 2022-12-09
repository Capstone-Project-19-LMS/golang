package categoryMockService

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CategoryMock struct {
	mock.Mock
}

func (c *CategoryMock) CreateCategory(category dto.CategoryTransaction) error {
	args := c.Called(category)

	return args.Error(0)
}

func (c *CategoryMock) DeleteCategory(id string) error {
	args := c.Called(id)

	return args.Error(0)
}
func (c *CategoryMock) GetCategoryByID(id string, user dto.User) (dto.GetCategory, error) {
	args := c.Called(id, user)

	return args.Get(0).(dto.GetCategory), args.Error(1)
}

func (c *CategoryMock) GetAllCategory() ([]dto.CategoryTransaction, error) {
	args := c.Called()

	return args.Get(0).([]dto.CategoryTransaction), args.Error(1)
}

func (c *CategoryMock) UpdateCategory(dto.CategoryTransaction) error {
	args := c.Called()

	return args.Error(0)
}