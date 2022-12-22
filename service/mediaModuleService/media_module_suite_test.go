package mediamoduleservice

import (
	"errors"
	"golang/models/dto"
	mediamodulemockrepository "golang/repository/mediaModuleRepository/mediaModuleMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteMediaModule struct {
	suite.Suite
	mediaModuleService MediaModuleService
	mock               *mediamodulemockrepository.MediaModuleMock
}

func (s *suiteMediaModule) SetupTest() {
	mock := &mediamodulemockrepository.MediaModuleMock{}
	s.mock = mock
	NewMediaModuleService := NewMediaModuleService(s.mock)
	s.mediaModuleService = NewMediaModuleService
}

func (s *suiteMediaModule) TestCreateMediaModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		Body            dto.MediaModuleTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success create media module",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.MediaModuleTransaction{
				ID:       "abcde",
				Url:      "tes",
				ModuleID: "abcde",
			},
			nil,
			false,
			nil,
		},
		{
			"fail create media module",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.MediaModuleTransaction{
				ID:       "abcde",
				Url:      "tes",
				ModuleID: "abcde",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("CreateMediaModule", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.mediaModuleService.CreateMediaModule(v.Body)
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

func (s *suiteMediaModule) TestDeleteMediaModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success delete media module",
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
			"fail delete media module",
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
		mockCall := s.mock.On("DeleteMediaModule", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.mediaModuleService.DeleteMediaModule(v.ParamID)
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

func (s *suiteMediaModule) TestGetMediaModuleByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		MockReturnBody  dto.MediaModule
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.MediaModule
		ExpectedError   error
	}{
		{
			"success get media module by id ",
			"abcde",
			dto.MediaModule{
				ID:        "abcde",
				Url:       "tes",
				ModuleID:  "abcde",
			},
			nil,
			true,
			dto.MediaModule{
				ID:        "abcde",
				Url:       "tes",
				ModuleID:  "abcde",
			},
			nil,
		},
		{
			"failed get media module by id",
			"abcde",

			dto.MediaModule{},
			gorm.ErrRecordNotFound,
			false,
			dto.MediaModule{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetMediaModuleByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			mediaModule, err := s.mediaModuleService.GetMediaModuleByID(v.ParamID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, mediaModule)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, mediaModule)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteMediaModule) TestGetAllMediaModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		MockReturnBody  []dto.MediaModule
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.MediaModule
		ExpectedError   error
	}{
		{
			"success get all media module",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.MediaModule{
				{
					ID:        "abcde",
					Url:       "tes",
					ModuleID:  "abcde",
				},
				{
					ID:        "abcde2",
					Url:       "tes2",
					ModuleID:  "abcde2",
				},
			},
			nil,
			true,
			[]dto.MediaModule{
				{
					ID:        "abcde",
					Url:       "tes",
					ModuleID:  "abcde",
				},
				{
					ID:        "abcde2",
					Url:       "tes2",
					ModuleID:  "abcde2",
				},
			},
			nil,
		},
		{
			"failed get all media module",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.MediaModule{},
			gorm.ErrRecordNotFound,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllMediaModule").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			mediaModule, err := s.mediaModuleService.GetAllMediaModule()
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, mediaModule)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, mediaModule)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteMediaModule) TestUpdateMediaModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		Body            dto.MediaModuleTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success update media module",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.MediaModuleTransaction{
				ID:       "abcde2",
				Url:      "tes2",
				ModuleID: "abcde2",
			},
			nil,
			false,
			nil,
		},
		{
			"fail update media module",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.MediaModuleTransaction{
				ID:       "abcde2",
				Url:      "tes2",
				ModuleID: "abcde2",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("UpdateMediaModule", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.mediaModuleService.UpdateMediaModule(v.Body)
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

func TestSuiteMediaModule(t *testing.T) {
	suite.Run(t, new(suiteMediaModule))
}
