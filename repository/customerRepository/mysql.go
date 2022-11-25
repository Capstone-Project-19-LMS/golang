package customerRepository

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/models/model"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

// CreateCustomer implements CustomerRepository
func (u *customerRepository) CreateCustomer(customer dto.CostumerRegister) error {
	customerModel := model.Customer{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		Password:     customer.Password,
		ProfileImage: "https://t3.ftcdn.net/jpg/03/46/83/96/360_F_346839683_6nAPzbhpSkIpb8pmAwufkC7c5eD7wYws.jpg",
	}
	err := u.db.Create(&customerModel).Error
	if err != nil {
		return err
	}
	return nil
}

// LoginCustomer implements CustomerRepository
func (u *customerRepository) LoginCustomer(customer dto.CostumerLogin) (dto.CostumerResponseGet, error) {
	var customerLogin dto.Costumer
	err := u.db.Model(&model.Customer{}).First(&customerLogin, "email = ?", customer.Email).Error
	if err != nil {
		return dto.CostumerResponseGet{}, err
	}
	match := helper.CheckPasswordHash(customer.Password, customerLogin.Password)
	if !match {
		return dto.CostumerResponseGet{}, errors.New(constantError.ErrorEmailOrPasswordNotMatch)
	}
	var customerLoginResponse dto.CostumerResponseGet = dto.CostumerResponseGet{
		ID:           customerLogin.ID,
		Name:         customerLogin.Name,
		Email:        customerLogin.Email,
		Password:     customerLogin.Password,
		ProfileImage: customerLogin.ProfileImage,
	}
	return customerLoginResponse, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
