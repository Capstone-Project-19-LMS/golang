package favoriteService

import (
	"errors"
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/repository/courseRepository/courseMockRepository"
	"golang/repository/favoriteRepository/favoriteMockRepository"
	"golang/service/courseService"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteFavorite struct {
	suite.Suite
	favoriteService FavoriteService
	courseService   courseService.CourseService
	mockFavorite    *favoriteMockRepository.FavoriteMock
	mockCourse      *courseMockRepository.CourseMock
}

func (s *suiteFavorite) SetupTest() {
	s.mockFavorite = &favoriteMockRepository.FavoriteMock{}
	s.mockCourse = &courseMockRepository.CourseMock{}
	NewFavoriteService := NewFavoriteService(s.mockFavorite, s.mockCourse)
	s.favoriteService = NewFavoriteService
}

func (s *suiteFavorite) TestAddFavorite() {
	testCase := []struct {
		Name                       string
		User                       dto.User
		Body                       dto.FavoriteTransaction
		MockReturnGetFavoriteError error
		MockReturnGetCourseError   error
		MockReturnAddFavoriteError error
		HasReturnError             bool
		ExpectedError              error
	}{
		{
			"success take course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			gorm.ErrRecordNotFound,
			nil,
			nil,
			false,
			nil,
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			gorm.ErrRecordNotFound,
			gorm.ErrRecordNotFound,
			nil,
			true,
			gorm.ErrRecordNotFound,
		},
		{
			"fail get favorite",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			nil,
			nil,
			nil,
			true,
			errors.New(constantError.ErrorCustomerAlreadyFavoriteCourse),
		},
		{
			"fail add favorite",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			gorm.ErrRecordNotFound,
			nil,
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallGetCourseByID := s.mockCourse.On("GetCourseByID", v.Body.CourseID).Return(dto.Course{}, v.MockReturnGetCourseError)
		mockCallGetFavorite := s.mockFavorite.On("GetFavorite", v.Body.CourseID, v.Body.CustomerID).Return(dto.Favorite{}, v.MockReturnGetFavoriteError)
		mockCallAddFavorite := s.mockFavorite.On("AddFavorite", mock.Anything).Return(v.MockReturnAddFavoriteError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.favoriteService.AddFavorite(v.Body)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetCourseByID.Unset()
		mockCallGetFavorite.Unset()
		mockCallAddFavorite.Unset()
	}
}

func (s *suiteFavorite) TestDeleteFavorite() {
	testCase := []struct {
		Name                       string
		User                       dto.User
		CourseID                   string
		MockReturnGetFavorite      dto.Favorite
		MockReturnGetFavoriteError error
		MockReturnDeleteFavoriteError      error
		HasReturnError             bool
		ExpectedError              error
	}{
		{
			"success delete course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Favorite{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			nil,
			nil,
			false,
			nil,
		},
		{
			"fail delete favorite",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Favorite{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			nil,
			errors.New("error"),
			true,
			errors.New("error"),
		},
		{
			"fail get favorite, favorite not found",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Favorite{},
			gorm.ErrRecordNotFound,
			nil,
			true,
			errors.New(constantError.ErrorCustomerNotFavoriteCourse),
		},
		{
			"fail get favorite",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Favorite{},
			errors.New("error getting favorite"),
			nil,
			true,
			errors.New("error getting favorite"),
		},
		{
			"fail not authorized",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Favorite{
				ID:         "abcde",
				CustomerID: "abcdefg",
				CourseID:   "abcde",
			},
			nil,
			nil,
			true,
			errors.New(constantError.ErrorNotAuthorized),
		},
	}
	for _, v := range testCase {
		mockCallGetFavorite := s.mockFavorite.On("GetFavorite", v.CourseID, v.User.ID).Return(v.MockReturnGetFavorite, v.MockReturnGetFavoriteError)
		mockCallDeleteFavorite := s.mockFavorite.On("DeleteFavorite", v.CourseID).Return(v.MockReturnDeleteFavoriteError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.favoriteService.DeleteFavorite(v.CourseID, v.User.ID)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetFavorite.Unset()
		mockCallDeleteFavorite.Unset()
	}
}

func (s *suiteFavorite) TestGetFavoriteByCustomerID() {
	testCase := []struct {
		Name            string
		User            dto.User
		MockReturnBody  []dto.Course
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.GetCourse
		ExpectedError   error
	}{
		{
			"success history course",
			dto.User{
				ID:   "customer1",
				Role: "customer",
			},
			[]dto.Course{
				{
					ID:                 "course1",
					Name:               "test",
					Description:        "test",
					Objective:          "test",
					Price:              10000,
					Discount:           0,
					Thumbnail:          "test",
					Capacity:           100,
					InstructorID:       "abcde",
					CategoryID:         "abcde",
					ProgressModule:     2,
					ProgressPercentage: 100,
					IsFinish:           false,
					CustomerCourses: []dto.CustomerCourse{
						{
							ID:         "customerCourses1",
							CustomerID: "customer1",
							CourseID:   "course1",
							Status:     true,
							NoModule:   2,
							IsFinish:   false,
						},
					},
					Favorites: []dto.Favorite{
						{
							ID:         "favorite1",
							CustomerID: "customer1",
							CourseID:   "course1",
						},
					},
					Ratings: []dto.Rating{
						{
							ID:         "rating1",
							CustomerID: "customer1",
							CourseID:   "course1",
							Rating:     5,
						},
					},
					Modules: []dto.Module{
						{
							ID:       "module1",
							Name:     "test",
							Content:  "test",
							CourseID: "course1",
						},
					},
				},
			},
			nil,
			true,
			[]dto.GetCourse{
				{
					ID:                 "course1",
					Name:               "test",
					Description:        "test",
					Objective:          "test",
					Price:              10000,
					Discount:           0,
					Thumbnail:          "test",
					Capacity:           100,
					InstructorID:       "abcde",
					NumberOfModules:    1,
					AmountCustomer:     1,
					ProgressModule:     2,
					ProgressPercentage: 100,
					IsFinish:           false,
					Rating:             5,
					Favorite:           true,
					StatusEnroll:       true,
				},
			},
			nil,
		},
		{
			"failed history course",
			dto.User{
				ID:   "customer1",
				Role: "customer",
			},
			[]dto.Course{},
			errors.New("error"),
			false,
			nil,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mockFavorite.On("GetFavoriteByCustomerID", v.User.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			course, err := s.favoriteService.GetFavoriteByCustomerID(v.User.ID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, course)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, course)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteFavorite(t *testing.T) {
	suite.Run(t, new(suiteFavorite))
}
