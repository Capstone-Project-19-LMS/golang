package modulerepository

import (
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type moduleRepository struct {
	db *gorm.DB
}

// CreateModule implements ModuleRepository
func (mr *moduleRepository) CreateModule(module dto.ModuleTransaction) error {

	var moduleModel model.Module
	err := copier.Copy(&moduleModel, &module)
	if err != nil {
		return err
	}
	var checkModule []model.Module

	mr.db.Where("course_id=?", module.CourseID).Find(&checkModule)

	for _, data := range checkModule {

		if data.NoModule == module.NoModule {

			return gorm.ErrPrimaryKeyRequired
		}
	}
	err = mr.db.Model(&model.Module{}).Create(&moduleModel).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteModule implements ModuleRepository
func (mr *moduleRepository) DeleteModule(id string) error {
	// delete data Module from database by id
	err := mr.db.Select("media_modules", "assignments").Where("id = ?", id).Delete(&model.Module{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// GetAllModule implements ModuleRepository
func (mr *moduleRepository) GetAllModule() ([]dto.Module, error) {
	var moduleModels []model.Module
	// get data sub category from database by user
	err := mr.db.Model(&model.Module{}).Preload("MediaModules").Preload("Assignment").Find(&moduleModels).Error
	if err != nil {
		return nil, err
	}
	// copy data from model to dto
	var modules []dto.Module
	err = copier.Copy(&modules, &moduleModels)
	if err != nil {
		return nil, err
	}
	return modules, nil
}

// GetModuleByID implements ModuleRepository
func (mr *moduleRepository) GetModuleByID(id, customerID string) (dto.ModuleAcc, error) {
	var moduleModel model.Module
	err := mr.db.Model(&model.Module{}).Preload("MediaModules").Preload("Assignment").Where("id = ?", id).Find(&moduleModel)
	if err.Error != nil {
		return dto.ModuleAcc{}, err.Error
	}
	if err.RowsAffected <= 0 {
		return dto.ModuleAcc{}, gorm.ErrRecordNotFound
	}

	var CustomerCourses model.CustomerCourse

	mr.db.Where("course_id=?", moduleModel.CourseID).Where("customer_id=?", customerID).Find(&CustomerCourses)

	if !CustomerCourses.Status {
		return dto.ModuleAcc{}, gorm.ErrRecordNotFound
	}

	if moduleModel.NoModule > int(CustomerCourses.NoModule) {
		return dto.ModuleAcc{}, gorm.ErrRecordNotFound
	}

	// copy data from model to dto
	var Module dto.ModuleAcc
	errCopy := copier.Copy(&Module, &moduleModel)
	if errCopy != nil {
		return dto.ModuleAcc{}, errCopy
	}
	return Module, nil
}

func (mr *moduleRepository) GetModuleByCourseID(courseID, customerID string) ([]dto.Module, error) {

	var moduleModels []model.Module
	err := mr.db.Model(&model.Module{}).Where("course_id = ?", courseID).Find(&moduleModels).Error
	if err != nil {
		return nil, err
	}

	var CustomerCourses model.CustomerCourse

	mr.db.Where("course_id=?", courseID).Where("customer_id=?", customerID).Find(&CustomerCourses)

	if !CustomerCourses.Status {
		return nil, err
	}
	// copy data from model to dto
	var modules []dto.Module
	err = copier.Copy(&modules, &moduleModels)
	if err != nil {
		return nil, err
	}
	return modules, nil

}

// UpdateModule implements ModuleRepository
func (mr *moduleRepository) UpdateModule(module dto.ModuleTransaction) error {
	var moduleModel model.Module
	errCopy := copier.Copy(&moduleModel, &module)
	if errCopy != nil {
		return errCopy
	}

	// update account with new data
	err := mr.db.Model(&model.Module{}).Where("id = ?", module.ID).Updates(&moduleModel)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewModuleRepository(db *gorm.DB) ModuleRepository {
	return &moduleRepository{
		db: db,
	}
}
