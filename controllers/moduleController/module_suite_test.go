package moduleController

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	"golang/models/model"
	modulemockservice "golang/service/moduleService/moduleMockService"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteModule struct {
	suite.Suite
	moduleController *ModuleController
	mock             *modulemockservice.ModuleMock
}

func (s *suiteModule) SetupTest() {
	mock := &modulemockservice.ModuleMock{}
	s.mock = mock
	s.moduleController = &ModuleController{
		ModuleService: s.mock,
	}
}

func (s *suiteModule) TestCreateModule() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.ModuleTransaction
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create module",
			"POST",
			dto.ModuleTransaction{
				Name:     "tes",
				Content:  "tes",
				CourseID: "tes",
				NoModule: 1,
				Url:      "tes",
			},
			nil,
			http.StatusOK,
			"success create module",
		},
		{
			"fail bind data",
			"POST",
			dto.ModuleTransaction{
				Name:     "tes",
				Content:  "tes",
				CourseID: "tes",
				NoModule: 1,
				Url:      "tes",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.ModuleTransaction{
				Name:    "tes",
				Content: "tes",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create module",
			"POST",
			dto.ModuleTransaction{
				Name:     "tes",
				Content:  "tes",
				CourseID: "tes",
				NoModule: 1,
				Url:      "tes",
			},

			errors.New("fail create module"),
			http.StatusInternalServerError,
			"fail create module",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateModule", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/module/", bytes.NewBuffer(res))
			if i != 1 {
				r.Header.Set("Content-Type", "application/json")
			}
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module/create")

			err := s.moduleController.CreateModule(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestDeleteModule() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success delete module",
			"DELETE",
			"abcde",
			nil,
			http.StatusOK,
			"success delete module",
		},
		{
			"fail delete module",
			"DELETE",
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail delete module",
		},
		{
			"fail delete module",
			"DELETE",
			"abcde",

			errors.New("fail delete module"),
			http.StatusInternalServerError,
			"fail delete module",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteModule", v.ParamID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/module/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.moduleController.DeleteModule(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestGetModuleByID() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		CustomerCourseID   model.CustomerCourse
		MockReturnBody     dto.ModuleCourseAcc
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.ModuleCourseAcc
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get module by id",
			"GET",
			"abcde",
			model.CustomerCourse{
				ID: "abcde",
			},
			dto.ModuleCourseAcc{
				ID:        "abcde",
				Name:      "tes",
				Content:   "tes",
				CourseID:  "tes",
				NoModule:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
				Assignment: dto.Assignment{
					ID:                  "abcde",
					Title:               "tes",
					ModuleID:            "abcde",
					CustomerAssignments: []dto.CustomerAssignment{},
				},
			},
			nil,
			true,
			dto.ModuleCourseAcc{
				ID:        "abcde",
				Name:      "tes",
				Content:   "tes",
				CourseID:  "tes",
				NoModule:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
				Assignment: dto.Assignment{
					ID:                  "abcde",
					Title:               "tes",
					ModuleID:            "abcde",
					CustomerAssignments: []dto.CustomerAssignment{},
				},
			},
			http.StatusOK,
			"success get module by id",
		},
		{
			"fail get module by id",
			"GET",
			"abcde",
			model.CustomerCourse{
				ID: "abcde",
			},
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.ModuleCourseAcc{},
			http.StatusInternalServerError,
			"fail get module by id",
		},
		{
			"fail get module by id",
			"GET",
			"abcde",
			model.CustomerCourse{
				ID: "abcde",
			},
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.ModuleCourseAcc{},
			http.StatusInternalServerError,
			"fail get module by id",
		},
		// {
		// 	"fail bind data",
		// 	"GET",
		// 	"abcde",
		// 	model.CustomerCourse{
		// 		ID:         "abcde",
		// 		CustomerID: "abcde",
		// 		CourseID:   "abcde",
		// 		NoModule:   1,
		// 		IsFinish:   true,
		// 		Status:     true,
		// 		CreatedAt:  time.Now(),
		// 		UpdatedAt:  time.Now(),
		// 		DeletedAt:  gorm.DeletedAt{},
		// 	},
		// 	dto.ModuleCourseAcc{},
		// 	nil,
		// 	false,
		// 	dto.ModuleCourseAcc{},
		// 	http.StatusInternalServerError,
		// 	"fail bind data",
		// },
	}

	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {

			// Create request
			r := httptest.NewRequest(v.Method, "/module/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module/get_by_id/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.moduleController.GetModuleByID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)
			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

		})
		// remove mock
		mockCall.Unset()
	}
}
func (s *suiteModule) TestGetModuleByIDifInstructor() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnBody     dto.ModuleCourseAcc
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.ModuleCourseAcc
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get module by id",
			"GET",
			"abcde",
			dto.ModuleCourseAcc{
				ID:        "abcde",
				Name:      "tes",
				Content:   "tes",
				CourseID:  "tes",
				NoModule:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
				Assignment: dto.Assignment{
					ID:                  "abcde",
					Title:               "tes",
					ModuleID:            "abcde",
					CustomerAssignments: []dto.CustomerAssignment{},
				},
			},
			nil,
			true,
			dto.ModuleCourseAcc{
				ID:        "abcde",
				Name:      "tes",
				Content:   "tes",
				CourseID:  "tes",
				NoModule:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
				Assignment: dto.Assignment{
					ID:                  "abcde",
					Title:               "tes",
					ModuleID:            "abcde",
					CustomerAssignments: []dto.CustomerAssignment{},
				},
			},
			http.StatusOK,
			"success get module by id",
		},
		{
			"fail get module by id",
			"GET",
			"abcde",
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.ModuleCourseAcc{},
			http.StatusInternalServerError,
			"fail get module by id",
		},
		{
			"fail get module by id",
			"GET",
			"abcde",
			dto.ModuleCourseAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.ModuleCourseAcc{},
			http.StatusInternalServerError,
			"fail get module by id",
		},
		// {
		// 	"fail bind data",
		// 	"GET",
		// 	"abcde",
		// 	model.CustomerCourse{
		// 		ID:         "abcde",
		// 		CustomerID: "abcde",
		// 		CourseID:   "abcde",
		// 		NoModule:   1,
		// 		IsFinish:   true,
		// 		Status:     true,
		// 		CreatedAt:  time.Now(),
		// 		UpdatedAt:  time.Now(),
		// 		DeletedAt:  gorm.DeletedAt{},
		// 	},
		// 	dto.ModuleCourseAcc{},
		// 	nil,
		// 	false,
		// 	dto.ModuleCourseAcc{},
		// 	http.StatusInternalServerError,
		// 	"fail bind data",
		// },
	}

	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByIDifInstructor", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {

			// Create request
			r := httptest.NewRequest(v.Method, "/module/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module/get_by_id/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.moduleController.GetModuleByIDifInstructor(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)
			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestGetModuleByCourseIDifInstructror() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnBody     []dto.ModuleCourse
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.ModuleCourse
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get module by course id",
			"GET",
			"abcde",
			[]dto.ModuleCourse{{
				ID:        "abcde",
				Name:      "tes",
				Content:   "tes",
				CourseID:  "tes",
				NoModule:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
				Assignment: dto.Assignment{
					ID:                  "abcde",
					Title:               "tes",
					ModuleID:            "abcde",
					CustomerAssignments: []dto.CustomerAssignment{},
				},
			},
			},
			nil,
			true,
			[]dto.ModuleCourse{{
				ID:        "abcde",
				Name:      "tes",
				Content:   "tes",
				CourseID:  "tes",
				NoModule:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
				MediaModules: []dto.MediaModule{
					{
						ID:       "abcde",
						Url:      "tes",
						ModuleID: "abcde",
					},
				},
				Assignment: dto.Assignment{
					ID:                  "abcde",
					Title:               "tes",
					ModuleID:            "abcde",
					CustomerAssignments: []dto.CustomerAssignment{},
				},
			},
			},
			http.StatusOK,
			"success get module by course id",
		},
		{
			"fail get module by course id",
			"GET",
			"abcde",
			[]dto.ModuleCourse{},
			gorm.ErrRecordNotFound,
			false,
			[]dto.ModuleCourse{},
			http.StatusInternalServerError,
			"fail get module by course id",
		},
		{
			"fail get module by course id",
			"GET",
			"abcde",
			[]dto.ModuleCourse{},
			gorm.ErrRecordNotFound,
			false,
			[]dto.ModuleCourse{},
			http.StatusInternalServerError,
			"fail get module by course id",
		},
	}

	for _, v := range testCase {
		mockCall := s.mock.On("GetModuleByCourseIDifInstructror", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {

			// Create request
			r := httptest.NewRequest(v.Method, "/module/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module/get_by_course_id/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.moduleController.GetModuleByCourseIDifInstructror(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)
			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestGetAllModule() {
	testCase := []struct {
		Name               string
		Method             string
		MockReturnBody     []dto.ModuleCourse
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.ModuleCourse
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get all module",
			"GET",

			[]dto.ModuleCourse{
				{
					ID:        "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
					Name:      "tes",
					Content:   "tes",
					CourseID:  "abcde",
					NoModule:  1,
				},
				{
					ID:        "abcdef",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
					Name:      "tes 2",
					Content:   "tes 2",
					CourseID:  "abcdef",
					NoModule:  2,
				},
			},
			nil,
			true,
			[]dto.ModuleCourse{
				{
					ID:        "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
					Name:      "tes",
					Content:   "tes",
					CourseID:  "abcde",
					NoModule:  1,
				},
				{
					ID:        "abcdef",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
					Name:      "tes 2",
					Content:   "tes 2",
					CourseID:  "abcdef",
					NoModule:  2,
				},
			},
			http.StatusOK,
			"success get all module",
		},
		{
			"fail get all module",
			"GET",

			[]dto.ModuleCourse{},
			errors.New("error"),
			false,
			[]dto.ModuleCourse{},
			http.StatusInternalServerError,
			"fail get all module",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllModule").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/module", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module")

			err := s.moduleController.GetAllModule(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteModule) TestUpdateModule() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.ModuleTransaction
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update module",
			"POST",
			dto.ModuleTransaction{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				NoModule: 1,
			},
			"abcde",
			nil,
			http.StatusOK,
			"success update module",
		},
		{
			"fail bind data",
			"POST",
			dto.ModuleTransaction{
				Name:    "tes",
				Content: "tes",
			},
			"abcde",
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"fail update module",
			"POST",
			dto.ModuleTransaction{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				NoModule: 1,
			},
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail update module",
		},
		{
			"fail update module",
			"POST",
			dto.ModuleTransaction{
				ID:       "abcde",
				Name:     "tes",
				Content:  "tes",
				CourseID: "abcde",
				NoModule: 1,
			},
			"abcde",
			errors.New("fail update module"),
			http.StatusInternalServerError,
			"fail update module",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateModule", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/module/update/"+v.ParamID, bytes.NewBuffer(res))
			if i != 1 {
				r.Header.Set("Content-Type", "application/json")
			}
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/module/update/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.moduleController.UpdateModule(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteModule(t *testing.T) {
	suite.Run(t, new(suiteModule))
}
