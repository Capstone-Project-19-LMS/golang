package courseService

import (
	"errors"
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/repository/categoryRepository/categoryMockRepository"
	"golang/repository/courseRepository/courseMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCourse struct {
	suite.Suite
	courseService   CourseService
	mockCourse      *courseMockRepository.CourseMock
	mockCategory    *categoryMockRepository.CategoryMock
}

func (s *suiteCourse) SetupTest() {
	s.mockCourse = &courseMockRepository.CourseMock{}
	s.mockCategory = &categoryMockRepository.CategoryMock{}
	NewCourseService := NewCourseService(s.mockCourse, s.mockCategory)
	s.courseService = NewCourseService
}

func (s *suiteCourse) TestCreateCourse() {
	testCase := []struct {
		Name                    string
		User                    dto.User
		Body                    dto.CourseTransaction
		MockReturnCategoryError error
		MockReturnCourseError   error
		HasReturnError          bool
		ExpectedError           error
	}{
		{
			"success create course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			dto.CourseTransaction{
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			nil,
			nil,
			false,
			nil,
		},
		{
			"fail create course",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			dto.CourseTransaction{
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			nil,
			errors.New("error"),
			true,
			errors.New("error"),
		},
		{
			"fail capacity is lower than 0",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			dto.CourseTransaction{
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     -1,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			nil,
			nil,
			true,
			errors.New(constantError.ErrorCapacityLowerThanZero),
		},
		{
			"fail error category",
			dto.User{
				ID:   "1",
				Role: "intructor",
			},
			dto.CourseTransaction{
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     -1,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			errors.New("error category"),
			nil,
			true,
			errors.New(constantError.ErrorCategoryNotFound),
		},
	}
	for _, v := range testCase {
		mockCallCategory := s.mockCategory.On("GetCategoryByID", v.Body.CategoryID, v.User).Return(dto.Category{}, v.MockReturnCategoryError)
		mockCallCourse := s.mockCourse.On("CreateCourse", mock.Anything).Return(v.MockReturnCourseError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.courseService.CreateCourse(v.Body, v.User)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallCategory.Unset()
		mockCallCourse.Unset()
	}
}

func (s *suiteCourse) TestDeleteCourse() {
	testCase := []struct {
		Name                     string
		User                     dto.User
		ParamID                  string
		MockReturnGetCourse      dto.Course
		MockReturnGetCourseError error
		MockReturnDeleteError    error
		HasReturnError           bool
		ExpectedError            error
	}{
		{
			"success delete course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.Course{
				ID:           "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			nil,
			nil,
			false,
			nil,
		},
		{
			"fail delete course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.Course{
				ID:           "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			nil,
			errors.New("error"),
			true,
			errors.New("error"),
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.Course{
				ID:           "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcdef",
			},
			errors.New("error"),
			nil,
			true,
			errors.New("error"),
		},
		{
			"fail not authorized",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.Course{
				ID:           "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcdef",
			},
			nil,
			nil,
			true,
			errors.New(constantError.ErrorNotAuthorized),
		},
	}
	for _, v := range testCase {
		mockCallGetCourse := s.mockCourse.On("GetCourseByID", v.ParamID).Return(v.MockReturnGetCourse, v.MockReturnGetCourseError)
		mockCallDeleteCourse := s.mockCourse.On("DeleteCourse", v.ParamID).Return(v.MockReturnDeleteError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.courseService.DeleteCourse(v.ParamID, v.User.ID)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetCourse.Unset()
		mockCallDeleteCourse.Unset()
	}
}

func (s *suiteCourse) TestGetAllCourse() {
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
			"success get all course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.Course{
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
			nil,
			true,
			[]dto.GetCourse{
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
			"fail len course is 0",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.Course{},
			nil,
			true,
			[]dto.GetCourse{},
			nil,
		},
		{
			"failed get all course",
			dto.User{
				ID:   "abcde",
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
		mockCall := s.mockCourse.On("GetAllCourse", v.User).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			course, err := s.courseService.GetAllCourse(v.User)
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

func (s *suiteCourse) TestGetCourseByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		ParamUser       dto.User
		MockReturnBody  dto.Course
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.GetCourseByID
		ExpectedError   error
	}{
		{
			"success get course by id and customer",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.Course{
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
			nil,
			true,
			dto.GetCourseByID{
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
				Ratings: []dto.Rating{
					{
						ID:         "abcde",
						CustomerID: "abcde",
						CourseID:   "abcde",
						Rating:     5,
					},
				},
				Modules: []dto.ModuleTransaction{
					{
						ID:       "efgh",
						Name:     "test",
						Content:  "test",
						CourseID: "abcde",
					},
				},
			},
			nil,
		},
		{
			"success get course by id and instructor",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			dto.Course{
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
			nil,
			true,
			dto.GetCourseByID{
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
				Favorite:           false,
				NumberOfModules:    1,
				StatusEnroll:       false,
				ProgressModule:     2,
				ProgressPercentage: 100,
				IsFinish:           false,
				Ratings: []dto.Rating{
					{
						ID:         "abcde",
						CustomerID: "abcde",
						CourseID:   "abcde",
						Rating:     5,
					},
				},
				Modules: []dto.ModuleTransaction{
					{
						ID:       "efgh",
						Name:     "test",
						Content:  "test",
						CourseID: "abcde",
					},
				},
			},
			nil,
		},
		{
			"fail not authorized instructor",
			"abcde",
			dto.User{
				ID:   "abcdefg",
				Role: "instructor",
			},
			dto.Course{
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
			nil,
			false,
			dto.GetCourseByID{},
			errors.New(constantError.ErrorNotAuthorized),
		},
		{
			"failed get course by id and customer",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.Course{},
			gorm.ErrRecordNotFound,
			false,
			dto.GetCourseByID{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mockCourse.On("GetCourseByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			course, err := s.courseService.GetCourseByID(v.ParamID, v.ParamUser)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, course)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCourse) TestGetCourseEnrollByID() {
	testCase := []struct {
		Name                             string
		User                             dto.User
		ParamID                          string
		MockReturnGetCourse              dto.Course
		MockReturnGetCourseError         error
		MockReturnGetCourseEnrolled      []dto.CustomerCourseEnroll
		MockReturnGetCourseEnrolledError error
		HasReturnBody                    bool
		ExpectedBody                     []dto.CustomerCourseEnroll
		ExpectedError                    error
	}{
		{
			"success get all course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Course{
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
			nil,
			[]dto.CustomerCourseEnroll{
				{
					ID:           "test1",
					Name:         "test1",
					Email:        "test1@test.com",
					ProfileImage: "abcde.jpg",
					StatusEnroll: true,
				},
				{
					ID:           "test2",
					Name:         "test2",
					Email:        "test2@test.com",
					ProfileImage: "abcde.jpg",
					StatusEnroll: false,
				},
			},
			nil,
			true,
			[]dto.CustomerCourseEnroll{
				{
					ID:           "test1",
					Name:         "test1",
					Email:        "test1@test.com",
					ProfileImage: "abcde.jpg",
					StatusEnroll: true,
				},
				{
					ID:           "test2",
					Name:         "test2",
					Email:        "test2@test.com",
					ProfileImage: "abcde.jpg",
					StatusEnroll: false,
				},
			},
			nil,
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Course{},
			gorm.ErrRecordNotFound,
			[]dto.CustomerCourseEnroll{},
			nil,
			false,
			nil,
			errors.New(constantError.ErrorCourseNotFound),
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Course{},
			errors.New("fail get course"),
			[]dto.CustomerCourseEnroll{},
			nil,
			false,
			nil,
			errors.New("fail get course"),
		},
		{
			"success get all course",
			dto.User{
				ID:   "abcdefg",
				Role: "instructor",
			},
			"abcde",
			dto.Course{
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
			nil,
			[]dto.CustomerCourseEnroll{},
			nil,
			false,
			nil,
			errors.New(constantError.ErrorNotAuthorized),
		},
		{
			"success get all course",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			"abcde",
			dto.Course{
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
			nil,
			[]dto.CustomerCourseEnroll{},
			errors.New("fail get customer enroll"),
			false,
			nil,
			errors.New("fail get customer enroll"),
		},
	}
	for _, v := range testCase {
		mockCallGetCourse := s.mockCourse.On("GetCourseByID", v.ParamID).Return(v.MockReturnGetCourse, v.MockReturnGetCourseError)
		mockCallGetCourseEnrolled := s.mockCourse.On("GetCourseEnrollByID", v.ParamID).Return(v.MockReturnGetCourseEnrolled, v.MockReturnGetCourseEnrolledError)
		s.T().Run(v.Name, func(t *testing.T) {
			customerCourse, err := s.courseService.GetCourseEnrollByID(v.ParamID, v.User)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, customerCourse)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, customerCourse)
			}
		})
		// remove mock
		mockCallGetCourse.Unset()
		mockCallGetCourseEnrolled.Unset()
	}
}

func (s *suiteCourse) TestUpdateCourse() {
	testCase := []struct {
		Name                     string
		User                     dto.User
		ParamID                  string
		Body                     dto.CourseTransaction
		MockReturnGetCourse      dto.Course
		MockReturnGetCourseError error
		MockReturnUpdateError    error
		HasReturnError           bool
		ExpectedError            error
	}{
		{
			"success update course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.CourseTransaction{
				ID:   "abcde",
				Name:        "test",
				Description: "test",
				InstructorID: "abcde",
			},
			dto.Course{
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
			nil,
			nil,
			false,
			nil,
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.CourseTransaction{
				ID:   "abcde",
				Name:        "test",
				Description: "test",
				InstructorID: "abcde",
			},
			dto.Course{},
			errors.New("fail get course"),
			nil,
			true,
			errors.New("fail get course"),
		},
		{
			"fail no authorized instructor",
			dto.User{
				ID:   "abcdefg",
				Role: "intructor",
			},
			"abcde",
			dto.CourseTransaction{
				ID:   "abcde",
				Name:        "test",
				Description: "test",
				InstructorID: "abcdefg",
			},
			dto.Course{
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
			errors.New("fail get course"),
			nil,
			true,
			errors.New("fail get course"),
		},
		{
			"fail update course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			"abcde",
			dto.CourseTransaction{
				ID:   "abcde",
				Name:        "test",
				Description: "test",
				InstructorID: "abcde",
			},
			dto.Course{
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
			nil,
			errors.New(constantError.ErrorNotAuthorized),
			true,
			errors.New(constantError.ErrorNotAuthorized),
		},
	}
	for _, v := range testCase {
		mockCallGetCourse := s.mockCourse.On("GetCourseByID", v.ParamID).Return(v.MockReturnGetCourse, v.MockReturnGetCourseError)
		mockCallUpdate := s.mockCourse.On("UpdateCourse", mock.Anything).Return(v.MockReturnUpdateError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.courseService.UpdateCourse(v.Body)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetCourse.Unset()
		mockCallUpdate.Unset()
	}
}

func TestSuiteCourse(t *testing.T) {
	suite.Run(t, new(suiteCourse))
}
