package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) string
}

type Repository interface {
	Register(userDomain *Domain) Domain
	Login(userDomain *Domain) Domain
}
