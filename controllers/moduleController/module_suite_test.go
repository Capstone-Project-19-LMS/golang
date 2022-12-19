package moduleController

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	modulemockservice "golang/service/moduleService/moduleMockService"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestSuiteModule(t *testing.T) {
	suite.Run(t, new(suiteModule))
}
