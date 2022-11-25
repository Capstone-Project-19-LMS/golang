package mediamoduleservice

import (
	"golang/helper"
	"golang/models/dto"
	mediamodulerepository "golang/repository/mediaModuleRepository"
)

type MediaModuleService interface {
	CreateMediaModule(dto.MediaModuleTransaction) error
	DeleteMediaModule(id string) error
	GetAllMediaModule() ([]dto.MediaModule, error)
	GetMediaModuleByID(id string) (dto.MediaModule, error)
	UpdateMediaModule(dto.MediaModuleTransaction) error
}

type mediaModuleService struct {
	mediamoduleRepo mediamodulerepository.MediaModuleRepository
}

// CreateModule implements ModuleService
func (mms *mediaModuleService) CreateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	id := helper.GenerateUUID()
	mediaModule.ID = id
	err := mms.mediamoduleRepo.CreateMediaModule(mediaModule)
	if err != nil {
		return err
	}
	return nil
}

// DeleteModule implements ModuleService
func (mms *mediaModuleService) DeleteMediaModule(id string) error {
	// call repository to delete account
	err := mms.mediamoduleRepo.DeleteMediaModule(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllModule implements ModuleService
func (mms *mediaModuleService) GetAllMediaModule() ([]dto.MediaModule, error) {
	mediaModules, err := mms.mediamoduleRepo.GetAllMediaModule()
	if err != nil {
		return nil, err
	}
	return mediaModules, nil
}

// GetModuleByID implements ModuleService
func (mms *mediaModuleService) GetMediaModuleByID(id string) (dto.MediaModule, error) {
	mediaModule, err := mms.mediamoduleRepo.GetMediaModuleByID(id)
	if err != nil {
		return dto.MediaModule{}, err
	}
	return mediaModule, nil
}

// UpdateModule implements ModuleService
func (mms *mediaModuleService) UpdateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	// call repository to update Module
	err := mms.mediamoduleRepo.UpdateMediaModule(mediaModule)
	if err != nil {
		return err
	}
	return nil
}

func NewMediaModuleService(mediamoduleRepo mediamodulerepository.MediaModuleRepository) MediaModuleService {
	return &mediaModuleService{
		mediamoduleRepo: mediamoduleRepo,
	}
}
