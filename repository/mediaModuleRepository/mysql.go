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

// CreateMediaModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) CreateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	var mediaModuleModel model.MediaModule
	err := copier.Copy(&mediaModuleModel, &mediaModule)
	if err != nil {
		return err
	}

	err = mmr.db.Model(&model.MediaModule{}).Create(&mediaModuleModel).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteMediaModule implements MediaModuleRepository
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

// GetAllMediaModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) GetAllMediaModule() ([]dto.MediaModule, error) {
	var mediaModuleModels []model.MediaModule
	// get data sub category from database by user
	err := mmr.db.Model(&model.MediaModule{}).Find(&mediaModuleModels).Error
	if err != nil {
		return nil, err
	}
	// copy data from model to dto
	var mediaModules []dto.MediaModule
	err = copier.Copy(&mediaModules, &mediaModuleModels)
	if err != nil {
		return nil, err
	}
	return mediaModules, nil
}

// GetMediaModuleByID implements MediaModuleRepository
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
	errCopy := copier.Copy(&Module, &mediaModuleModel)
	if errCopy != nil {
		return dto.MediaModule{}, errCopy
	}
	return Module, nil
}

// UpdateMediaModule implements MediaModuleRepository
func (mmr *mediaModuleRepository) UpdateMediaModule(mediaModule dto.MediaModuleTransaction) error {
	var mediaModuleModel model.MediaModule
	errCopy := copier.Copy(&mediaModuleModel, &mediaModule)
	if errCopy != nil {
		return errCopy
	}
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
