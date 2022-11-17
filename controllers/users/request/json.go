package request

import (
	"golang/businesses/users"

	"github.com/go-playground/validator/v10"
)

type Register struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *Register) ToDomainRegister() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *Login) ToDomainLogin() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *Register) ValidateRegister() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
func (req *Login) ValidateLogin() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
