package mediamodulerepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type mediaModuleRepository struct {
	db *gorm.DB
}

// CreateModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) CreateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	var mediaModuleModel model.MediaModule
	copier.Copy(&mediaModuleModel, &mediaModule)

	err := mmr.db.Model(&model.MediaModule{}).Create(&mediaModuleModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) DeleteMediaModule(id string) error {
	// delete data Module from database by id
	err := mmr.db.Where("id = ?", id).Unscoped().Delete(&model.MediaModule{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAllModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) GetAllMediaModule() ([]dto.MediaModule, error) {
	var mediaModuleModels []model.MediaModule
	// get data sub category from database by user
	err := mmr.db.Model(&model.MediaModule{}).Find(&mediaModuleModels).Error
	if err != nil {
		return nil, err
	}
	// copy data from model to dto
	var mediaModules []dto.MediaModule
	copier.Copy(&mediaModules, &mediaModuleModels)

	return mediaModules, nil
}

// GetModuleByID implements MediaModuleRepository
func (mmr *mediaModuleRepository) GetMediaModuleByID(id string) (dto.MediaModule, error) {
	var mediaModuleModel model.MediaModule
	err := mmr.db.Model(&model.MediaModule{}).Where("id = ?", id).Find(&mediaModuleModel)
	if err.Error != nil {
		return dto.MediaModule{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.MediaModule{}, gorm.ErrRecordNotFound
	}

	// copy data from model to dto
	var Module dto.MediaModule
	copier.Copy(&Module, &mediaModuleModel)

	return Module, nil
}

// UpdateModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) UpdateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	var mediaModuleModel model.MediaModule
	copier.Copy(&mediaModuleModel, &mediaModule)

	// update account with new data
	err := mmr.db.Model(&model.MediaModule{}).Where("id = ?", mediaModule.ID).Updates(&mediaModuleModel)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewMediaModuleRepository(db *gorm.DB) MediaModuleRepository {
	return &mediaModuleRepository{
		db: db,
	}
}
