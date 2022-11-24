package instructorrepository

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/models/model"

	"gorm.io/gorm"
)

type instructorrepository struct {
	db *gorm.DB
}

// Createinstructor implements instructorrepository
func (u *instructorrepository) CreateInstructor(instructor dto.InstructorRegister) error {
	instructorModel := model.Instructor{
		ID:           instructor.ID,
		Name:         instructor.Name,
		Email:        instructor.Email,
		Password:     instructor.Password,
		ProfileImage: "https://t3.ftcdn.net/jpg/03/46/83/96/360_F_346839683_6nAPzbhpSkIpb8pmAwufkC7c5eD7wYws.jpg",
	}
	err := u.db.Create(&instructorModel).Error
	if err != nil {
		return err
	}
	return nil
}

// Logininstructor implements instructorrepository
func (u *instructorrepository) LoginInstructor(instructor dto.InstructorLogin) (dto.InstructorResponseGet, error) {
	var instructorLogin dto.Instructor
	err := u.db.Model(&model.Instructor{}).First(&instructorLogin, "email = ?", instructor.Email).Error
	if err != nil {
		return dto.InstructorResponseGet{}, err
	}
	match := helper.CheckPasswordHash(instructor.Password, instructorLogin.Password)
	if !match {
		return dto.InstructorResponseGet{}, errors.New(constantError.ErrorEmailOrPasswordNotMatch)
	}
	var instructorLoginResponse dto.InstructorResponseGet = dto.InstructorResponseGet{
		ID:           instructorLogin.ID,
		Name:         instructorLogin.Name,
		Email:        instructorLogin.Email,
		Password:     instructorLogin.Password,
		ProfileImage: instructorLogin.ProfileImage,
	}
	return instructorLoginResponse, nil
}

func Newinstructorrepository(db *gorm.DB) InstructorRepository {
	return &instructorrepository{
		db: db,
	}
}
