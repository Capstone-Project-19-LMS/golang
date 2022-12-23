package customerCourseService

import (
	"errors"
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/repository/courseRepository/courseMockRepository"
	"golang/repository/customerCourseRepository/customerCourseMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCustomerCourse struct {
	suite.Suite
	customerCourseService CustomerCourseService
	mockCustomerCourse    *customerCourseMockRepository.CustomerCourseMock
	mockCourse            *courseMockRepository.CourseMock
}

func (s *suiteCustomerCourse) SetupTest() {
	s.mockCustomerCourse = &customerCourseMockRepository.CustomerCourseMock{}
	s.mockCourse = &courseMockRepository.CourseMock{}
	NewCustomerCourseService := NewCustomerCourseService(s.mockCustomerCourse, s.mockCourse)
	s.customerCourseService = NewCustomerCourseService
}

func (s *suiteCustomerCourse) TestDeleteCustomerCourse() {
	testCase := []struct {
		Name                             string
		User                             dto.User
		CourseID                         string
		MockReturnGetCustomerCourse      dto.CustomerCourse
		MockReturnGetCustomerCourseError error
		MockReturnGetCourse              dto.Course
		MockReturnGetCourseError         error
		MockReturnDeleteError            error
		BodyUpdateCourse                 dto.CourseTransaction
		MockReturnUpdateCourse           error
		HasReturnError                   bool
		ExpectedError                    error
	}{
		{
			"success delete course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			nil, dto.Course{
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
			dto.CourseTransaction{
				ID:       "abcde",
				Capacity: 11,
			},
			nil,
			false,
			nil,
		},
		{
			"fail delete course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			nil, dto.Course{
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
			dto.CourseTransaction{
				ID:       "abcde",
				Capacity: 11,
			},
			nil,
			true,
			errors.New("error"),
		},
		{
			"fail get customer course not found",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{},
			gorm.ErrRecordNotFound,
			dto.Course{},
			nil,
			nil,
			dto.CourseTransaction{},
			nil,
			true,
			errors.New(constantError.ErrorCustomerNotEnrolled),
		},
		{
			"fail get customer course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{},
			errors.New(constantError.ErrorCustomerNotEnrolled),
			dto.Course{},
			nil,
			nil,
			dto.CourseTransaction{},
			nil,
			true,
			errors.New(constantError.ErrorCustomerNotEnrolled),
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false},
			nil,
			dto.Course{},
			gorm.ErrRecordNotFound,
			nil,
			dto.CourseTransaction{},
			nil,
			true,
			gorm.ErrRecordNotFound,
		},
		{
			"fail not authorized",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{
				ID:         "abcde",
				CustomerID: "abcdefg",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			nil, dto.Course{},
			nil,
			nil,
			dto.CourseTransaction{},
			nil,
			true,
			errors.New(constantError.ErrorNotAuthorized),
		},
		{
			"fail update course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourse{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			nil, 
			dto.Course{
				ID:           "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     10,
				CategoryID:   "abcde",
				InstructorID: "abcde",},
			nil,
			nil,
			dto.CourseTransaction{
				ID:       "abcde",
				Capacity: 11,},
			gorm.ErrRecordNotFound,
			true,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCallGetCustomerCourse := s.mockCustomerCourse.On("GetCustomerCourse", v.CourseID, v.User.ID).Return(v.MockReturnGetCustomerCourse, v.MockReturnGetCustomerCourseError)
		GetCourse := s.mockCourse.On("GetCourseByID", v.CourseID).Return(v.MockReturnGetCourse, v.MockReturnGetCourseError)
		UpdateCourse := s.mockCourse.On("UpdateCourse", v.BodyUpdateCourse).Return(v.MockReturnUpdateCourse)
		mockCallDeleteCustomerCourse := s.mockCustomerCourse.On("DeleteCustomerCourse", v.CourseID).Return(v.MockReturnDeleteError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.customerCourseService.DeleteCustomerCourse(v.CourseID, v.User.ID)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetCustomerCourse.Unset()
		GetCourse.Unset()
		UpdateCourse.Unset()
		mockCallDeleteCustomerCourse.Unset()
	}
}

func (s *suiteCustomerCourse) TestGetHistoryCourseByCustomerID() {
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
			"failed history course",
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
		mockCall := s.mockCustomerCourse.On("GetHistoryCourseByCustomerID", v.User.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			course, err := s.customerCourseService.GetHistoryCourseByCustomerID(v.User.ID)
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

func (s *suiteCustomerCourse) TestGetCustomerCourseEnrollByID() {
	testCase := []struct {
		Name            string
		User            dto.User
		ID string
		MockReturnBody  dto.CustomerCourseEnroll
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.CustomerCourseEnroll
		ExpectedError   error
	}{
		{
			"success history course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourseEnroll{
				ID:         "abcde",
				CustomerID: "abcde",
				Name: 	 "test",
				Email: 	 "test@gmail.com",
				ProfileImage: "test",
				StatusEnroll: true,
			},
			nil,
			true,
			dto.CustomerCourseEnroll{
				ID:         "abcde",
				CustomerID: "abcde",
				Name: 	 "test",
				Email: 	 "test@gmail.com",
				ProfileImage: "test",
				StatusEnroll: true,
			},
			nil,
		},
		{
			"failed get customer enroll by id",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.CustomerCourseEnroll{},
			errors.New("error"),
			false,
			dto.CustomerCourseEnroll{},
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mockCustomerCourse.On("GetCustomerCourseEnrollByID", v.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			course, err := s.customerCourseService.GetCustomerCourseEnrollByID(v.ID)
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

func (s *suiteCustomerCourse) TestTakeCourse() {
	testCase := []struct {
		Name                    string
		User                    dto.User
		Body                    dto.CustomerCourseTransaction
		MockReturnGetCustomerCourse      dto.CustomerCourse
		MockReturnGetCustomerCourseError error
		MockReturnGetCourse              dto.Course
		MockReturnGetCourseError         error
		MockReturnTakeCourseError   error
		BodyUpdateCourse                 dto.CourseTransaction
		MockReturnUpdateCourse           error
		HasReturnError          bool
		ExpectedError           error
	}{
		{
			"success take course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			dto.CustomerCourse{},
			gorm.ErrRecordNotFound, 
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
			dto.CourseTransaction{
				ID:       "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     9,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
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
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			dto.CustomerCourse{},
			gorm.ErrRecordNotFound, 
			dto.Course{},
			gorm.ErrRecordNotFound,
			nil,
			dto.CourseTransaction{},
			nil,
			true,
			gorm.ErrRecordNotFound,
		},
		{
			"fail capacity is 0",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			dto.CustomerCourse{},
			gorm.ErrRecordNotFound, 
			dto.Course{
				ID:           "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     0,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			nil,
			nil,
			dto.CourseTransaction{},
			nil,
			true,
			errors.New(constantError.ErrorCourseCapacity),
		},
		{
			"fail customer already enroll",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			dto.CustomerCourse{},
			nil, 
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
			dto.CourseTransaction{},
			nil,
			true,
			errors.New(constantError.ErrorCustomerAlreadyTakeCourse),
		},
		{
			"fail capacity is 0",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			dto.CustomerCourse{},
			gorm.ErrRecordNotFound, 
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
			dto.CourseTransaction{},
			nil,
			true,
			errors.New("error"),
		},
		{
			"fail update course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			dto.CustomerCourse{},
			gorm.ErrRecordNotFound, 
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
			dto.CourseTransaction{
				ID:       "abcde",
				Name:         "test",
				Description:  "test",
				Objective:    "test",
				Capacity:     9,
				CategoryID:   "abcde",
				InstructorID: "abcde",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallGetCourseByID := s.mockCourse.On("GetCourseByID", v.Body.CourseID).Return(v.MockReturnGetCourse, v.MockReturnGetCourseError)
		mockCallGetCustomerCourse := s.mockCustomerCourse.On("GetCustomerCourse", v.Body.CourseID, v.Body.CustomerID).Return(dto.CustomerCourse{}, v.MockReturnGetCustomerCourseError)
		mockCallTakeCourse := s.mockCustomerCourse.On("TakeCourse", mock.Anything).Return(v.MockReturnTakeCourseError)
		mockCallUpdateCourse := s.mockCourse.On("UpdateCourse", v.BodyUpdateCourse).Return(v.MockReturnUpdateCourse)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.customerCourseService.TakeCourse(v.Body)
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
		mockCallGetCustomerCourse.Unset()
		mockCallTakeCourse.Unset()
		mockCallUpdateCourse.Unset()
	}
}

func (s *suiteCustomerCourse) TestUpdateEnrollmentStatus() {
	testCase := []struct {
		Name                     string
		User                     dto.User
		Body                     dto.CustomerCourseTransaction
		MockReturnCustomerCourseError    error
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
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			nil,
			nil,
			false,
			nil,
		},
		{
			"fail error get customer course",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			errors.New("error"),
			nil,
			true,
			errors.New("error"),
		},
		{
			"fail update enrollment status",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:    true,
				NoModule: 0,
				IsFinish:           false,
			},
			nil,
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallGetCustomerCourse := s.mockCustomerCourse.On("GetCustomerCourseByID", v.Body.ID).Return(dto.CustomerCourse{}, v.MockReturnCustomerCourseError)
		mockCallUpdateEnrollmentStatus := s.mockCustomerCourse.On("UpdateEnrollmentStatus", v.Body).Return(v.MockReturnUpdateError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.customerCourseService.UpdateEnrollmentStatus(v.Body, v.User.ID)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetCustomerCourse.Unset()
		mockCallUpdateEnrollmentStatus.Unset()
	}
}

func TestSuiteCustomerCourse(t *testing.T) {
	suite.Run(t, new(suiteCustomerCourse))
}
