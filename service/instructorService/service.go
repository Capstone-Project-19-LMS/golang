package instructorservice

import (
	"golang/helper"
	"golang/models/dto"
	instructorrepository "golang/repository/instructorRepository"
)

type InstructorService interface {
	CreateInstructor(user dto.InstructorRegister) error
	LoginInstructor(user dto.InstructorLogin) (dto.InstructorResponseGet, error)
}

type instructorService struct {
	instructorRepo instructorrepository.InstructorRepository
}

// CreateInstructor implements instructorService
func (u *instructorService) CreateInstructor(user dto.InstructorRegister) error {
	// hash password
	id := helper.GenerateUUID()

	user.ID = id
	password, errPassword := helper.HashPassword(user.Password)
	user.Password = password
	if errPassword != nil {
		return errPassword
	}

	// call repository to save user
	err := u.instructorRepo.CreateInstructor(user)
	if err != nil {
		return err
	}
	return nil
}

// LoginInstructor implements instructorService
func (u *instructorService) LoginInstructor(user dto.InstructorLogin) (dto.InstructorResponseGet, error) {
	// call repository to get user
	InstructorLogin, err := u.instructorRepo.LoginInstructor(user)
	if err != nil {
		return dto.InstructorResponseGet{}, err
	}
	return InstructorLogin, nil
}

func NewinstructorService(instructorRepo instructorrepository.InstructorRepository) InstructorService {
	return &instructorService{
		instructorRepo: instructorRepo,
	}
}
