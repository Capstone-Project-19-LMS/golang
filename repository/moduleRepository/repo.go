package modulerepository

import "golang/models/dto"

type ModuleRepository interface {
	CreateModule(dto.ModuleTransaction) error
	DeleteModule(id string) error
	GetAllModule() ([]dto.Module, error)
	GetModuleByID(id string) (dto.Module, error)
	GetModuleByCourseID(courseID string) ([]dto.Module, error)
	UpdateModule(dto.ModuleTransaction) error
}
