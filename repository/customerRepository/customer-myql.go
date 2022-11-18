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
func (u *customerRepository) CreateCustomer(customer dto.UserRegister) error {
	customerModel := model.Customer{
		Name:     customer.Name,
		Email:    customer.Email,
		Password: customer.Password,
	}
	err := u.db.Create(&customerModel).Error
	if err != nil {
		return err
	}
	return nil
}

// LoginCustomer implements CustomerRepository
func (u *customerRepository) LoginCustomer(customer dto.UserLogin) (dto.UserResponseGet, error) {
	var customerLogin dto.User
	err := u.db.Model(&model.Customer{}).First(&customerLogin, "email = ?", customer.Email).Error
	if err != nil {
		return dto.UserResponseGet{}, err
	}
	match := helper.CheckPasswordHash(customer.Password, customerLogin.Password)
	if !match {
		return dto.UserResponseGet{}, errors.New(constantError.ErrorEmailOrPasswordNotMatch)
	}
	var customerLoginResponse dto.UserResponseGet = dto.UserResponseGet{
		ID:             customerLogin.ID,
		Name:           customerLogin.Name,
		Email:          customerLogin.Email,
		Password:       customerLogin.Password,
		ProfilePicture: customerLogin.ProfilePicture,
		Role:           "customer",
	}
	return customerLoginResponse, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}
