package modulerepository

import "golang/models/dto"

type ModuleRepository interface {
	CreateModule(dto.ModuleTransaction) error
	DeleteModule(id string) error
	GetAllModule() ([]dto.Module, error)
	GetModuleByID(id, customerID string) (dto.ModuleAcc, error)
	GetModuleByCourseID(courseID, customerID string) ([]dto.Module, error)
	UpdateModule(dto.ModuleTransaction) error
}
