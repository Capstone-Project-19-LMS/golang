package instructorrepository

import (
	"golang/models/dto"
)

type InstructorRepository interface {
	CreateInstructor(instructor dto.InstructorRegister) error
	LoginInstructor(instructor dto.InstructorLogin) (dto.InstructorResponseGet, error)
}
