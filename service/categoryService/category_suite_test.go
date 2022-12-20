package categoryService

import (
	"errors"
	"golang/models/dto"
	"golang/repository/categoryRepository/categoryMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCategory struct {
	suite.Suite
	categoryService CategoryService
	mock            *categoryMockRepository.CategoryMock
}

func (s *suiteCategory) SetupTest() {
	mock := &categoryMockRepository.CategoryMock{}
	s.mock = mock
	NewCategoryService := NewCategoryService(s.mock)
	s.categoryService = NewCategoryService
}

func (s *suiteCategory) TestCreateCategory() {
	testCase := []struct {
		Name            string
		User            dto.User
		Body            dto.CategoryTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success create category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
			},
			nil,
			false,
			nil,
		},
		{
			"fail create category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("CreateCategory", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.categoryService.CreateCategory(v.Body)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCategory) TestDeleteCategory() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID            string
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success delete category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			"abcde",
			nil,
			false,
			nil,
		},
		{
			"fail delete category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			"abcde",
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteCategory", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.categoryService.DeleteCategory(v.ParamID)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCategory) TestGetAllCategory() {
	testCase := []struct {
		Name            string
		User       dto.User
		MockReturnBody  []dto.CategoryTransaction
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.CategoryTransaction
		ExpectedError   error
	}{
		{
			"success get all category",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.CategoryTransaction{
				{
					ID:          "test1",
					Name:        "test1",
					Description: "test1",
				},
				{
					ID:          "test2",
					Name:        "test2",
					Description: "test2",
				},
			},
			nil,
			true,
			[]dto.CategoryTransaction{
				{
					ID:          "test1",
					Name:        "test1",
					Description: "test1",
				},
				{
					ID:          "test2",
					Name:        "test2",
					Description: "test2",
				},
			},
			nil,
		},
		{
			"failed get all category",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.CategoryTransaction{},
			gorm.ErrRecordNotFound,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllCategory").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			category, err := s.categoryService.GetAllCategory()
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, category)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, category)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCategory) TestGetCategoryByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		ParamUser       dto.User
		MockReturnBody  dto.Category
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.GetCategory
		ExpectedError   error
	}{
		{
			"success get category by id and customer",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.Category{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
				Courses: []dto.Course{
					{
						ID:                 "abcde",
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
								ID:         "abcde",
								CustomerID: "abcde",
								CourseID:   "abcde",
								Status:     true,
								NoModule:   2,
								IsFinish:   false,
							},
						},
						Favorites: []dto.Favorite{
							{
								ID:         "abcde",
								CustomerID: "abcde",
								CourseID:   "abcde",
							},
						},
						Ratings: []dto.Rating{
							{
								ID:         "abcde",
								CustomerID: "abcde",
								CourseID:   "abcde",
								Rating:     5,
							},
						},
						Modules: []dto.Module{
							{
								ID:       "efgh",
								Name:     "test",
								Content:  "test",
								CourseID: "abcde",
							},
						},
					},
				},
			},
			nil,
			true,
			dto.GetCategory{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
				Courses: []dto.GetCourseWithoutCategory{
					{
						ID:                 "abcde",
						Name:               "test",
						Description:        "test",
						Objective:          "test",
						Price:              10000,
						Discount:           0,
						Thumbnail:          "test",
						Capacity:           100,
						InstructorID:       "abcde",
						Rating:             5,
						Favorite:           true,
						NumberOfModules:    1,
						StatusEnroll:       true,
						ProgressModule:     2,
						ProgressPercentage: 100,
						IsFinish:           false,
					},
				},
			},
			nil,
		},
		{
			"failed get category by id and customer",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.Category{},
			gorm.ErrRecordNotFound,
			false,
			dto.GetCategory{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetCategoryByID", v.ParamID, v.ParamUser).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			category, err := s.categoryService.GetCategoryByID(v.ParamID, v.ParamUser)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, category)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCategory) TestUpdateCategory() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID            string
		Body            dto.CategoryTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success update category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			"abcde",
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
			},
			nil,
			false,
			nil,
		},
		{
			"fail update category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			"abcde",
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("UpdateCategory", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.categoryService.UpdateCategory(v.Body)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
