package customerRepository

import (
	"golang/models/dto"
)

type CustomerRepository interface {
	CreateCustomer(customer dto.CostumerRegister) error
	VerifikasiCustomer(input dto.CustomerVerif) error
	LoginCustomer(customer dto.CostumerLogin) (dto.CostumerResponseGet, error)
}
