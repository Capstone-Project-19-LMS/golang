package moduleservice

import (
	"golang/helper"
	"golang/models/dto"
	modulerepository "golang/repository/moduleRepository"
)

type ModuleService interface {
	CreateModule(dto.ModuleTransaction) error
	DeleteModule(id string) error
	GetAllModule() ([]dto.Module, error)
	GetModuleByID(id, customerID string) (dto.ModuleCourseAcc, error)
	GetModuleByCourseID(courseID, customerID string) ([]dto.ModuleCourse, error)
	UpdateModule(dto.ModuleTransaction) error
}

type moduleService struct {
	moduleRepo modulerepository.ModuleRepository
}

// CreateModule implements ModuleService
func (ms *moduleService) CreateModule(module dto.ModuleTransaction) error {
	id := helper.GenerateUUID()
	mediamoduleID := helper.GenerateUUID()
	module.ID = id
	module.MediaModuleID = mediamoduleID
	err := ms.moduleRepo.CreateModule(module)
	if err != nil {
		return err
	}
	return nil
}

// DeleteModule implements ModuleService
func (ms *moduleService) DeleteModule(id string) error {
	// call repository to delete account
	err := ms.moduleRepo.DeleteModule(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllModule implements ModuleService
func (ms *moduleService) GetAllModule() ([]dto.Module, error) {
	modules, err := ms.moduleRepo.GetAllModule()
	if err != nil {
		return nil, err
	}
	return modules, nil
}

// GetModuleByID implements ModuleService
func (ms *moduleService) GetModuleByID(id, customerID string) (dto.ModuleCourseAcc, error) {
	module, err := ms.moduleRepo.GetModuleByID(id, customerID)
	if err != nil {
		return dto.ModuleCourseAcc{}, err
	}
	return module, nil
}

func (ms *moduleService) GetModuleByCourseID(courseID, customerID string) ([]dto.ModuleCourse, error) {
	modules, err := ms.moduleRepo.GetModuleByCourseID(courseID, customerID)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

// UpdateModule implements ModuleService
func (ms *moduleService) UpdateModule(module dto.ModuleTransaction) error {
	// call repository to update Module
	err := ms.moduleRepo.UpdateModule(module)
	if err != nil {
		return err
	}
	return nil
}

func NewModuleService(moduleRepo modulerepository.ModuleRepository) ModuleService {
	return &moduleService{
		moduleRepo: moduleRepo,
	}
}
