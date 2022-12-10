package mediamodulerepository

import "golang/models/dto"

type MediaModuleRepository interface {
	CreateMediaModule(dto.MediaModuleTransaction) error
	DeleteMediaModule(id string) error
	GetAllMediaModule() ([]dto.MediaModule, error)
	GetMediaModuleByID(id string) (dto.MediaModule, error)
	UpdateMediaModule(dto.MediaModuleTransaction) error
}
