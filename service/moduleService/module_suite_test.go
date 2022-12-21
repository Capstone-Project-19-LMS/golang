package moduleservice

import (
	"errors"
	"golang/models/dto"
	moduleMockRepository "golang/repository/moduleRepository/moduleMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type suiteModule struct {
	suite.Suite
	moduleService ModuleService
	mock          *moduleMockRepository.ModuleMock
}

func (s *suiteModule) SetupTest() {
	mock := &moduleMockRepository.ModuleMock{}
	s.mock = mock
	NewmoduleService := NewModuleService(s.mock)
	s.moduleService = NewmoduleService
}

func (s *suiteModule) TestCreateModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		Body            dto.ModuleTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success create module",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.ModuleTransaction{
				Name:          "tes",
				Content:       "tes",
				CourseID:      "abcde",
				NoModule:      1,
				MediaModuleID: "abcde",
				Url:           "tes",
			},
			nil,
			false,
			nil,
		},
		{
			"fail create module",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.ModuleTransaction{
				Name:          "tes",
				Content:       "tes",
				CourseID:      "abcde",
				NoModule:      1,
				MediaModuleID: "abcde",
				Url:           "tes",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("CreateModule", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.moduleService.CreateModule(v.Body)
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

func (s *suiteModule) TestDeleteModule() {
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
		mockCall := s.mock.On("DeleteModule", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.moduleService.DeleteModule(v.ParamID)
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

func TestSuiteModule(t *testing.T) {
	suite.Run(t, new(suiteModule))
}
