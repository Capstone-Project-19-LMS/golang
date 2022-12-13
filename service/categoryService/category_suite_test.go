package categoryService

import (
	"golang/models/dto"
	"golang/repository/categoryRepository/categoryMockRepository"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCategory struct {
	suite.Suite
	categoryService CategoryService
	mock               *categoryMockRepository.CategoryMock
}

func (s *suiteCategory) SetupTest() {
	mock := &categoryMockRepository.CategoryMock{}
	s.mock = mock
	NewCategoryService := NewCategoryService(s.mock)
	s.categoryService = NewCategoryService
}

func (s *suiteCategory) TestGetCategoryByID() {
	testCase := []struct {
		Name               string
		ParamID            string
		ParamUser          dto.User
		MockReturnBody     dto.Category
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.GetCategory
		ExpectedError                 error
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
						ID:              "abcde",
						Name:            "test",
						Description:     "test",
						Objective:       "test",
						Price:           10000,
						Discount:        0,
						Thumbnail:       "test",
						Capacity:        100,
						InstructorID:    "abcde",
						CategoryID:      "abcde",
						ProgressModule: 2,
						ProgressPercentage: 100,
						CustomerCourses: []dto.CustomerCourse{
							{
								ID:         "abcde",
								CustomerID: "abcde",
								CourseID:   "abcde",
								Status:     true,
								NoModule:  2,
								IsFinish: false,
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
								Rating: 5,
							},
						},
						Modules: []dto.Module{
							{
								ID:          "efgh",
								Name:        "test",
								Content:    "test",
								CourseID:    "abcde",
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
						ID:              "abcde",
						Name:            "test",
						Description:     "test",
						Objective:       "test",
						Price:           10000,
						Discount:        0,
						Thumbnail:       "test",
						Capacity:        100,
						InstructorID:    "abcde",
						Rating:          5,
						Favorite:        true,
						NumberOfModules: 1,
						StatusEnroll: 		true,
						ProgressModule: 2,
						ProgressPercentage: 100,
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
				s.Equal(v.ExpectedBody, category)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
