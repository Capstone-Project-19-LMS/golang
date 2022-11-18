package customerRepository

import (
	"golang/models/dto"
)

type CustomerRepository interface {
	CreateCustomer(customer dto.UserRegister) error
	LoginCustomer(customer dto.UserLogin) (dto.UserResponseGet, error)
}