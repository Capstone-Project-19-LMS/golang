package customermockrepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type CustomerMock struct {
	mock.Mock
}

func (c *CustomerMock) CreateCustomer(customer dto.CostumerRegister) error {
	args := c.Called(customer)

	return args.Error(0)
}
func (c *CustomerMock) VerifikasiCustomer(input dto.CustomerVerif) error {
	args := c.Called(input)

	return args.Error(0)
}
func (c *CustomerMock) LoginCustomer(customer dto.CostumerLogin) (dto.CostumerResponseGet, error) {
	args := c.Called(customer)

	return args.Get(0).(dto.CostumerResponseGet), args.Error(0)
}
