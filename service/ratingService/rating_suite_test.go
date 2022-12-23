package ratingService

import (
	"errors"
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/repository/courseRepository/courseMockRepository"
	"golang/repository/ratingRepository/ratingMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteRating struct {
	suite.Suite
	ratingService RatingService
	mockRating    *ratingMockRepository.RatingMock
	mockCourse    *courseMockRepository.CourseMock
}

func (s *suiteRating) SetupTest() {
	s.mockRating = &ratingMockRepository.RatingMock{}
	s.mockCourse = &courseMockRepository.CourseMock{}
	NewRatingService := NewRatingService(s.mockRating, s.mockCourse)
	s.ratingService = NewRatingService
}

func (s *suiteRating) TestAddRating() {
	testCase := []struct {
		Name                     string
		User                     dto.User
		Body                     dto.RatingTransaction
		MockReturnGetRatingError error
		MockReturnGetCourse      dto.Course
		MockReturnGetCourseError error
		MockReturnAddRatingError error
		HasReturnError           bool
		ExpectedError            error
	}{
		{
			"success get rating",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			gorm.ErrRecordNotFound,
			dto.Course{
				ID:           "abcde",
				Name:         "abcde",
				Description:  "abcde",
				Objective:    "abcde",
				Capacity:     1,
				InstructorID: "abcde",
				IsFinish:     true,
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
				Role: "customer",
			},
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			nil,
			dto.Course{},
			gorm.ErrRecordNotFound,
			nil,
			true,
			gorm.ErrRecordNotFound,
		},
		{
			"fail customer not finish course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			nil,
			dto.Course{
				ID:           "abcde",
				Name:         "abcde",
				Description:  "abcde",
				Objective:    "abcde",
				Capacity:     1,
				InstructorID: "abcde",
				IsFinish:     false,
			},
			nil,
			nil,
			true,
			errors.New(constantError.ErrorCustomerNotFinishedCourse),
		},
		{
			"fail customer already rating course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			nil,
			dto.Course{
				ID:           "abcde",
				Name:         "abcde",
				Description:  "abcde",
				Objective:    "abcde",
				Capacity:     1,
				InstructorID: "abcde",
				IsFinish:     true,
			},
			nil,
			nil,
			true,
			errors.New(constantError.ErrorCustomerAlreadyRatingCourse),
		},
		{
			"fail add rating course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			gorm.ErrRecordNotFound,
			dto.Course{
				ID:           "abcde",
				Name:         "abcde",
				Description:  "abcde",
				Objective:    "abcde",
				Capacity:     1,
				InstructorID: "abcde",
				IsFinish:     true,
			},
			nil,
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallGetCourseByID := s.mockCourse.On("GetCourseByID", v.Body.CourseID).Return(v.MockReturnGetCourse, v.MockReturnGetCourseError)
		mockCallGetRating := s.mockRating.On("GetRatingByCourseIDCustomerID", v.Body.CourseID, v.Body.CustomerID).Return(dto.Rating{}, v.MockReturnGetRatingError)
		mockCallAddRating := s.mockRating.On("AddRating", mock.Anything).Return(v.MockReturnAddRatingError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.ratingService.AddRating(v.Body)
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
		mockCallGetRating.Unset()
		mockCallAddRating.Unset()
	}
}

func (s *suiteRating) TestDeleteRating() {
	testCase := []struct {
		Name                        string
		User                        dto.User
		CourseID                    string
		MockReturnGetRating         dto.Rating
		MockReturnGetRatingError    error
		MockReturnDeleteRatingError error
		HasReturnError              bool
		ExpectedError               error
	}{
		{
			"success delete course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{
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
			"fail delete rating",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{
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
			"fail get rating, rating not found",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{},
			gorm.ErrRecordNotFound,
			nil,
			true,
			errors.New(constantError.ErrorCustomerNotRatingCourse),
		},
		{
			"fail get rating",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{},
			errors.New("error getting rating"),
			nil,
			true,
			errors.New("error getting rating"),
		},
		{
			"fail not authorized",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{
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
		mockCallGetRating := s.mockRating.On("GetRatingByCourseIDCustomerID", v.CourseID, v.User.ID).Return(v.MockReturnGetRating, v.MockReturnGetRatingError)
		mockCallDeleteRating := s.mockRating.On("DeleteRating", v.CourseID).Return(v.MockReturnDeleteRatingError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.ratingService.DeleteRating(v.CourseID, v.User.ID)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallGetRating.Unset()
		mockCallDeleteRating.Unset()
	}
}

func (s *suiteRating) TestGetRatingByCourseID() {
	testCase := []struct {
		Name            string
		User            dto.User
		CourseID              string
		MockReturnGetCourseBody  dto.Course
		MockReturnGetCourseError error
		MockReturnGetRatingBody  []dto.Rating
		MockReturnGetRatingError error
		HasReturnBody   bool
		ExpectedBody    []dto.Rating
		ExpectedError   error
	}{
		{
			"success get rating by course",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			"course",
			dto.Course{
				ID: 		"course",
				Name: 		"test",
				Description: "test",
				Objective: 	"test",
				InstructorID: "abcde",
			},
			nil,
			[]dto.Rating{
				{
					ID: "rating1",
					Rating: 5,
					Testimonial: "test1",
					CourseID: "course",
				},
				{
					ID: "rating1",
					Rating: 4,
					Testimonial: "test1",
					CourseID: "course",
				},
			},
			nil,
			true,
			[]dto.Rating{
				{
					ID: "rating1",
					Rating: 5,
					Testimonial: "test1",
					CourseID: "course",
				},
				{
					ID: "rating1",
					Rating: 4,
					Testimonial: "test1",
					CourseID: "course",
				},
			},
			nil,
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			"course",
			dto.Course{},
			gorm.ErrRecordNotFound,
			[]dto.Rating{},
			nil,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
		{
			"fail not authorized",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			"course",
			dto.Course{
				ID: 		"course",
				Name: 		"test",
				Description: "test",
				Objective: 	"test",
				InstructorID: "abcdef",
			},
			nil,
			[]dto.Rating{},
			nil,
			false,
			nil,
			errors.New(constantError.ErrorNotAuthorized),
		},
		{
			"fail get rating",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			"course",
			dto.Course{
				ID: 		"course",
				Name: 		"test",
				Description: "test",
				Objective: 	"test",
				InstructorID: "abcde",
			},
			nil,
			[]dto.Rating{},
			errors.New("error"),
			false,
			[]dto.Rating{},
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallGetCourse := s.mockCourse.On("GetCourseByID", v.CourseID).Return(v.MockReturnGetCourseBody, v.MockReturnGetCourseError)
		mockCallGetRating := s.mockRating.On("GetRatingByCourseID", v.CourseID).Return(v.MockReturnGetRatingBody, v.MockReturnGetRatingError)
		s.T().Run(v.Name, func(t *testing.T) {
			ratings, err := s.ratingService.GetRatingByCourseID(v.CourseID, v.User.ID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, ratings)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, ratings)
			}
		})
		// remove mock
		mockCallGetCourse.Unset()
		mockCallGetRating.Unset()
	}
}

func (s *suiteRating) TestGetRatingByCourseIDCustomerID() {
	testCase := []struct {
		Name            string
		User            dto.User
		CourseID              string
		MockReturnGetRatingBody  dto.Rating
		MockReturnGetRatingError error
		HasReturnBody   bool
		ExpectedBody    dto.Rating
		ExpectedError   error
	}{
		{
			"success get rating by course and customer",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"course",
			dto.Rating{
					ID: "rating1",
					Rating: 5,
					Testimonial: "test1",
					CourseID: "course",
			},
			nil,
			true,
			dto.Rating{
					ID: "rating1",
					Rating: 5,
					Testimonial: "test1",
					CourseID: "course",
			},
			nil,
		},
		{
			"fail get course",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"course",
			dto.Rating{},
			gorm.ErrRecordNotFound,
			false,
			dto.Rating{},
			errors.New(constantError.ErrorCustomerNotRatingCourse),
		},
		{
			"fail get rating",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"course",
			dto.Rating{},
			errors.New("error"),
			false,
			dto.Rating{},
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallGetRating := s.mockRating.On("GetRatingByCourseIDCustomerID", v.CourseID, v.User.ID).Return(v.MockReturnGetRatingBody, v.MockReturnGetRatingError)
		s.T().Run(v.Name, func(t *testing.T) {
			ratings, err := s.ratingService.GetRatingByCourseIDCustomerID(v.CourseID, v.User.ID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, ratings)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, ratings)
			}
		})
		// remove mock
		mockCallGetRating.Unset()
	}
}

func (s *suiteRating) TestUpdateRating() {
	testCase := []struct {
		Name                     string
		User                     dto.User
		Body                     dto.RatingTransaction
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
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Rating: 5,
				IsPublish: false,
			},
			nil,
			false,
			nil,
		},
		{
			"fail update enrollment status",
			dto.User{
				ID:   "abcde",
				Role: "intructor",
			},
			dto.RatingTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Rating: 5,
				IsPublish: false,
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCallUpdateRating := s.mockRating.On("UpdateRating", v.Body).Return(v.MockReturnUpdateError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.ratingService.UpdateRating(v.Body)
			if v.HasReturnError {
				s.Error(err)
				s.Equal(v.ExpectedError, err)
				s.EqualError(err, v.ExpectedError.Error())
			} else {
				s.NoError(err)
			}
		})
		// remove mock
		mockCallUpdateRating.Unset()
	}
}

func TestSuiteRating(t *testing.T) {
	suite.Run(t, new(suiteRating))
}
