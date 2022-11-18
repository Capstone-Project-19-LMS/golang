package userService

import (
	"golang/helper"
	"golang/models/dto"
	"golang/repository/customerRepository"
)

type UserService interface {
	CreateUser(user dto.UserRegister) error
	LoginUser(user dto.UserLogin) (dto.UserResponseGet, error)
}

type userService struct {
	customerRepo customerRepository.CustomerRepository
}

// CreateUser implements UserService
func (u *userService) CreateUser(user dto.UserRegister) error {
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

// LoginUser implements UserService
func (u *userService) LoginUser(user dto.UserLogin) (dto.UserResponseGet, error) {
	// call repository to get user
	userLogin, err := u.customerRepo.LoginCustomer(user)
	if err != nil {
		return dto.UserResponseGet{}, err
	}
	return userLogin, nil
}

func NewUserService(customerRepo customerRepository.CustomerRepository) UserService {
	return &userService{
		customerRepo: customerRepo,
	}
}
