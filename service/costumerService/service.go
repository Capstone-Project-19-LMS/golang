package costumerService

import (
	"golang/helper"
	"golang/models/dto"
	"golang/repository/customerRepository"
)

type CostumerService interface {
	CreateCustomer(user dto.CostumerRegister) error
	LoginCostumer(user dto.CostumerLogin) (dto.CostumerResponseGet, error)
}

type costumerService struct {
	customerRepo customerRepository.CustomerRepository
}

// CreateCustomer implements costumerService
func (u *costumerService) CreateCustomer(user dto.CostumerRegister) error {
	// hash password
	password, errPassword := helper.HashPassword(user.Password)
	user.Password = password
	if errPassword != nil {
		return errPassword
	}

	// call repository to save user
	err := u.customerRepo.CreateCustomer(user)
	if err != nil {
		return err
	}
	return nil
}

// LoginCostumer implements costumerService
func (u *costumerService) LoginCostumer(user dto.CostumerLogin) (dto.CostumerResponseGet, error) {
	// call repository to get user
	CostumerLogin, err := u.customerRepo.LoginCustomer(user)
	if err != nil {
		return dto.CostumerResponseGet{}, err
	}
	return CostumerLogin, nil
}

func NewcostumerService(customerRepo customerRepository.CustomerRepository) CostumerService {
	return &costumerService{
		customerRepo: customerRepo,
	}
}
