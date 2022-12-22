package modulerepository

import "golang/models/dto"

type ModuleRepository interface {
	CreateModule(dto.ModuleTransaction) error
	DeleteModule(id string) error
	GetAllModule() ([]dto.Module, error)
	GetModuleByID(id, customerID string) (dto.ModuleCourseAcc, error)
	GetModuleByIDifInstructor(id string) (dto.ModuleCourseAcc, error)
	GetModuleByCourseID(courseID, customerID string) ([]dto.ModuleCourse, error)
	// GetModuleByCourseIDifInstructror(courseID string) ([]dto.ModuleCourse, error)
	UpdateModule(dto.ModuleTransaction) error
}
