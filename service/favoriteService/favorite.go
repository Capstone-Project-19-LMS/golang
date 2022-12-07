package favoriteService

import (
	"errors"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/repository/courseRepository"
	"golang/repository/favoriteRepository"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type FavoriteService interface {
	AddFavorite(favorite dto.FavoriteTransaction) error
	DeleteFavorite(courseID, customerID string) error
	GetFavoriteByCustomerID(customerID string) ([]dto.GetCourse, error)
}

type favoriteService struct {
	courseRepo   courseRepository.CourseRepository
	favoriteRepo favoriteRepository.FavoriteRepository
}

// AddFavorite implements FavoriteService
func (fs *favoriteService) AddFavorite(favorite dto.FavoriteTransaction) error {
	// check if course is not found
	_, err := fs.courseRepo.GetCourseByID(favorite.CourseID)
	if err != nil {
		return err
	}

	// check if customer already favorite the course
	_, err = fs.favoriteRepo.GetFavorite(favorite.CourseID, favorite.CustomerID)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New(constantError.ErrorCustomerAlreadyFavoriteCourse)
	}

	// create uuid for favorite
	id := helper.GenerateUUID()
	favorite.ID = id

	// call repository to create favorite
	err = fs.favoriteRepo.AddFavorite(favorite)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFavorite implements FavoriteService
func (fs *favoriteService) DeleteFavorite(courseID string, customerID string) error {
	// get customer course by id
	favorite, err := fs.favoriteRepo.GetFavorite(courseID, customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constantError.ErrorCustomerNotFavoriteCourse)
		}
		return err
	}

	// check if customer course is not belong to customer
	if favorite.CustomerID != customerID {
		return errors.New(constantError.ErrorNotAuthorized)
	}

	// delete favorite course
	err = fs.favoriteRepo.DeleteFavorite(favorite.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetFavoriteByCustomerID implements FavoriteService
func (fs *favoriteService) GetFavoriteByCustomerID(customerID string) ([]dto.GetCourse, error) {
	courses, err := fs.favoriteRepo.GetFavoriteByCustomerID(customerID)
	if err != nil {
		return nil, err
	}

	// get rating of all courses
	for i, course := range courses {
		rating := helper.GetRatingCourse(course)
		courses[i].Rating = rating
	}

	// get favorite of all courses
	for i := 0; i < len(courses); i++ {
		courses[i].Favorite = true
	}

	// get number of module
	for i, course := range courses {
		numberOfModule := len(course.Modules)
		courses[i].NumberOfModules = numberOfModule
	}

	// copy courses from dto.course to dto.GetCustomerCourse
	var customerCourses []dto.GetCourse
	err = copier.Copy(&customerCourses, &courses)
	if err != nil {
		return nil, err
	}
	return customerCourses, nil
}

func NewFavoriteService(favoriteRepo favoriteRepository.FavoriteRepository,
	courseRepo courseRepository.CourseRepository) FavoriteService {
	return &favoriteService{
		favoriteRepo: favoriteRepo,
		courseRepo:   courseRepo,
	}
}
