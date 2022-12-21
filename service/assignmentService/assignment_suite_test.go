package assignmentservice

import (
	"errors"
	"golang/models/dto"
	assignmentmockrepository "golang/repository/assignmentRepository/assignmentMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteAssignment struct {
	suite.Suite
	assignmentService AssignmentService
	mock              *assignmentmockrepository.AssignmentMock
}

func (s *suiteAssignment) SetupTest() {
	mock := &assignmentmockrepository.AssignmentMock{}
	s.mock = mock
	NewAssignmentService := NewAssignmentService(s.mock)
	s.assignmentService = NewAssignmentService
}

func (s *suiteAssignment) TestCreateAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		Body            dto.AssignmentTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success create assignment",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.AssignmentTransaction{
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			nil,
			false,
			nil,
		},
		{
			"fail create assignment",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.AssignmentTransaction{
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("CreateAssignment", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.assignmentService.CreateAssignment(v.Body)
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

func (s *suiteAssignment) TestDeleteAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success delete assignment",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			"abcde",
			nil,
			false,
			nil,
		},
		{
			"fail delete assignment",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			"abcde",
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteAssignment", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.assignmentService.DeleteAssignment(v.ParamID)
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

func (s *suiteAssignment) TestGetAssignmentByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		MockReturnBody  dto.Assignment
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.Assignment
		ExpectedError   error
	}{
		{
			"success get assignment by id ",
			"abcde",
			dto.Assignment{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
				CustomerAssignments: []dto.CustomerAssignment{
					{
						ID:           "abcde",
						File:         "tes",
						Grade:        1,
						AssignmentID: "abcde",
						CustomerID:   "abcde",
					},
				},
			},
			nil,
			true,
			dto.Assignment{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
				CustomerAssignments: []dto.CustomerAssignment{
					{
						ID:           "abcde",
						File:         "tes",
						Grade:        1,
						AssignmentID: "abcde",
						CustomerID:   "abcde",
					},
				},
			},
			nil,
		},
		{
			"failed get assignment by id",
			"abcde",

			dto.Assignment{},
			gorm.ErrRecordNotFound,
			false,
			dto.Assignment{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAssignmentByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			assignment, err := s.assignmentService.GetAssignmentByID(v.ParamID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, assignment)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, assignment)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteAssignment) TestGetAllAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		MockReturnBody  []dto.Assignment
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.Assignment
		ExpectedError   error
	}{
		{
			"success get all assignment",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.Assignment{
				{
					ID:          "abcde",
					Title:       "tes",
					Description: "tes",
					ModuleID:    "tes",
				},
				{
					ID:          "abcde",
					Title:       "tes2",
					Description: "tes2",
					ModuleID:    "tes2",
				},
			},
			nil,
			true,
			[]dto.Assignment{
				{
					ID:          "abcde",
					Title:       "tes",
					Description: "tes",
					ModuleID:    "tes",
				},
				{
					ID:          "abcde",
					Title:       "tes2",
					Description: "tes2",
					ModuleID:    "tes2",
				},
			},
			nil,
		},
		{
			"failed get all assignment",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.Assignment{},
			gorm.ErrRecordNotFound,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllAssignment").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			assignment, err := s.assignmentService.GetAllAssignment()
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, assignment)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, assignment)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteAssignment) TestUpdateAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		Body            dto.AssignmentTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success update assignment",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.AssignmentTransaction{
				ID:          "abcde",
				Title:       "tes2",
				Description: "tes2",
				ModuleID:    "tes2",
			},
			nil,
			false,
			nil,
		},
		{
			"fail update assignment",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.AssignmentTransaction{
				ID:          "abcde",
				Title:       "tes2",
				Description: "tes2",
				ModuleID:    "tes2",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("UpdateAssignment", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.assignmentService.UpdateAssignment(v.Body)
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

func TestSuiteAssignment(t *testing.T) {
	suite.Run(t, new(suiteAssignment))
}
