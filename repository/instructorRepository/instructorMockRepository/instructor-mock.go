package instructormockrepository

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type InstructorMock struct {
	mock.Mock
}

func (c *InstructorMock) CreateInstructor(instructor dto.InstructorRegister) error {
	args := c.Called(instructor)

	return args.Error(0)
}

func (c *InstructorMock) LoginInstructor(instructor dto.InstructorLogin) (dto.InstructorResponseGet, error) {
	args := c.Called(instructor)

	return args.Get(0).(dto.InstructorResponseGet), args.Error(0)
}
