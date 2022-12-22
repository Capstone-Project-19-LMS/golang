package moduleservice

import (
	"errors"
	"golang/models/dto"
	moduleMockRepository "golang/repository/moduleRepository/moduleMockRepository"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
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
			"success delete module",
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
			"fail delete module",
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

func (s *suiteModule) TestGetModuleByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		CustomerID      string
		MockReturnBody  dto.ModuleCourseAcc
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.ModuleCourseAcc
		ExpectedError   error
	}{
		{
			"success get module by id ",
			"abcde",
			"abcde",
			dto.ModuleCourseAcc{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				Course: struct {
					Name        string "json:\"name\""
					Description string "json:\"description\""
					Objective   string "json:\"objective\""
				}{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
			},
			nil,
			true,
			dto.ModuleCourseAcc{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				Course: struct {
					Name        string "json:\"name\""
					Description string "json:\"description\""
					Objective   string "json:\"objective\""
				}{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
			},
			nil,
		},
		{
			"failed get module by id",
			"abcde",
			"abcde",
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			module, err := s.moduleService.GetModuleByID(v.ParamID, v.CustomerID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, module)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, module)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}
func (s *suiteModule) TestGetModuleByCourseIDifInstructror() {
	testCase := []struct {
		Name            string
		ParamID         string
		MockReturnBody  []dto.ModuleCourse
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.ModuleCourse
		ExpectedError   error
	}{
		{
			"success get module by id ",
			"abcde",
			[]dto.ModuleCourse{
				{
					ID:       "abcde",
					Name:     "tes",
					Content:  "tes",
					CourseID: "abcde",
					Course: struct {
						Name        string "json:\"name\""
						Description string "json:\"description\""
						Objective   string "json:\"objective\""
					}{},
					MediaModules: []dto.MediaModule{
						{
							ID:       "abcde",
							Url:      "tes",
							ModuleID: "abcde",
						},
					},
				},
			},
			nil,
			true,
			[]dto.ModuleCourse{
				{
					ID:       "abcde",
					Name:     "tes",
					Content:  "tes",
					CourseID: "abcde",
					Course: struct {
						Name        string "json:\"name\""
						Description string "json:\"description\""
						Objective   string "json:\"objective\""
					}{},
					MediaModules: []dto.MediaModule{
						{
							ID:       "abcde",
							Url:      "tes",
							ModuleID: "abcde",
						},
					},
				},
			},
			nil,
		},
		{
			"failed get module by course id",
			"abcde",
			[]dto.ModuleCourse{},
			gorm.ErrRecordNotFound,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByCourseIDifInstructror", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			module, err := s.moduleService.GetModuleByCourseIDifInstructror(v.ParamID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, module)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, module)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}
func (s *suiteModule) TestGetModuleByIDifInstructor() {
	testCase := []struct {
		Name    string
		ParamID string

		MockReturnBody  dto.ModuleCourseAcc
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.ModuleCourseAcc
		ExpectedError   error
	}{
		{
			"success get module by course id ",
			"abcde",
			dto.ModuleCourseAcc{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				Course: struct {
					Name        string "json:\"name\""
					Description string "json:\"description\""
					Objective   string "json:\"objective\""
				}{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
			},
			nil,
			true,
			dto.ModuleCourseAcc{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				Course: struct {
					Name        string "json:\"name\""
					Description string "json:\"description\""
					Objective   string "json:\"objective\""
				}{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
			},
			nil,
		},
		{
			"failed get module by course id",
			"abcde",
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByIDifInstructor", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			module, err := s.moduleService.GetModuleByIDifInstructor(v.ParamID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, module)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, module)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestGetAllModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		MockReturnBody  []dto.Module
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.Module
		ExpectedError   error
	}{
		{
			"success get all Module",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.Module{
				{
					ID:       "abcde",
					Name:     "tes",
					Content:  "tes",
					CourseID: "abcde",
				},
				{
					ID:       "abcde",
					Name:     "tes",
					Content:  "tes",
					CourseID: "abcde",
				},
			},
			nil,
			true,
			[]dto.Module{
				{
					ID:       "abcde",
					Name:     "tes",
					Content:  "tes",
					CourseID: "abcde",
				},
				{
					ID:       "abcde",
					Name:     "tes",
					Content:  "tes",
					CourseID: "abcde",
				},
			},
			nil,
		},
		{
			"failed get all Module",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.Module{},
			gorm.ErrRecordNotFound,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllModule").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			Module, err := s.moduleService.GetAllModule()
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, Module)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, Module)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) GetModuleByCourseID() {
	testCase := []struct {
		Name            string
		CourseID        string
		CustomerID      string
		MockReturnBody  []dto.ModuleCourse
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.ModuleCourse
		ExpectedError   error
	}{
		{
			"success get module by id ",
			"abcde",
			"abcde",
			[]dto.ModuleCourse{
				{
					ID:        "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
					Name:      "tes",
					Content:   "tes",
					CourseID:  "abcde",
					Course: struct {
						Name        string "json:\"name\""
						Description string "json:\"description\""
						Objective   string "json:\"objective\""
					}{},
					MediaModules: []dto.MediaModule{
						{

							ID:        "abcde",
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
							DeletedAt: gorm.DeletedAt{},
							Url:       "tes",
							ModuleID:  "abcde",
						},
						{

							ID:        "abcde",
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
							DeletedAt: gorm.DeletedAt{},
							Url:       "tes",
							ModuleID:  "abcde",
						},
					},
				},
			},
			nil,
			true,
			[]dto.ModuleCourse{},
			nil,
		},
		{
			"failed get module by id",
			"abcde",
			"abcde",
			[]dto.ModuleCourse{},
			gorm.ErrRecordNotFound,
			false,
			[]dto.ModuleCourse{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByCourseID", v.CourseID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			module, err := s.moduleService.GetModuleByCourseID(v.CourseID, v.CustomerID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, module)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, module)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestUpdateModule() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		Body            dto.ModuleTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success update module",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.ModuleTransaction{
				ID:            "abcde",
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
			"fail update module",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.ModuleTransaction{
				ID:            "abcde",
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
		mockCall := s.mock.On("UpdateModule", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.moduleService.UpdateModule(v.Body)
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
