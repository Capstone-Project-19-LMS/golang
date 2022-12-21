package mediamodulemockservice

import (
	"golang/models/dto"

	"github.com/stretchr/testify/mock"
)

type MediaModuleMock struct {
	mock.Mock
}

func (c *MediaModuleMock) CreateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	args := c.Called(mediaModule)

	return args.Error(0)
}

func (c *MediaModuleMock) DeleteMediaModule(id string) error {
	args := c.Called(id)

	return args.Error(0)
}

func (c *MediaModuleMock) GetAllMediaModule() ([]dto.MediaModule, error) {
	args := c.Called()

	return args.Get(0).([]dto.MediaModule), args.Error(1)
}

func (c *MediaModuleMock) GetMediaModuleByID(id string) (dto.MediaModule, error) {
	args := c.Called(id)

	return args.Get(0).(dto.MediaModule), args.Error(1)
}

func (c *MediaModuleMock) UpdateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	args := c.Called(mediaModule)

	return args.Error(0)
}
