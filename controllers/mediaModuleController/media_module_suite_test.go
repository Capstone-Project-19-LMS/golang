package mediamodulecontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	mediamodulemockservice "golang/service/mediaModuleService/mediaModuleMockService"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteMediaModule struct {
	suite.Suite
	mediaModuleController *MediaModuleController
	mock                  *mediamodulemockservice.MediaModuleMock
}

func (s *suiteMediaModule) SetupTest() {
	mock := &mediamodulemockservice.MediaModuleMock{}
	s.mock = mock
	s.mediaModuleController = &MediaModuleController{
		MediaModuleService: s.mock,
	}
}

func (s *suiteMediaModule) TestCreateMediaModule() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.MediaModuleTransaction
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create media module",
			"POST",
			dto.MediaModuleTransaction{
				Url:      "tes",
				ModuleID: "tes",
			},
			nil,
			http.StatusOK,
			"success create media module",
		},
		{
			"fail bind data",
			"POST",
			dto.MediaModuleTransaction{
				Url:      "tes",
				ModuleID: "tes",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.MediaModuleTransaction{
				Url: "tes",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create media module",
			"POST",
			dto.MediaModuleTransaction{
				Url:      "tes",
				ModuleID: "tes",
			},

			errors.New("fail create media module"),
			http.StatusInternalServerError,
			"fail create media module",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateMediaModule", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/media_module/", bytes.NewBuffer(res))
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
			ctx.SetPath("/media_module/create")

			err := s.mediaModuleController.CreateMediaModule(ctx)
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

func (s *suiteMediaModule) TestDeleteMediaModule() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success delete media module",
			"DELETE",
			"abcde",
			nil,
			http.StatusOK,
			"success delete media module",
		},
		{
			"fail delete media module",
			"DELETE",
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail delete media module",
		},
		{
			"fail delete media module",
			"DELETE",
			"abcde",
			errors.New("fail delete media module"),
			http.StatusInternalServerError,
			"fail delete media module",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteMediaModule", v.ParamID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/media_module/delete/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/media_module/delete/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.mediaModuleController.DeleteMediaModule(ctx)
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

func (s *suiteMediaModule) TestGetMediaModuleByID() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnBody     dto.MediaModule
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.MediaModule
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get media module by id",
			"GET",
			"abcde",

			dto.MediaModule{
				ID:        "abcde",
				Url:       "tes",
				ModuleID:  "abcde",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			nil,
			true,
			dto.MediaModule{
				ID:        "abcde",
				Url:       "tes",
				ModuleID:  "abcde",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			http.StatusOK,
			"success get media module by id",
		},
		{
			"fail get media module by id",
			"GET",
			"abcde",

			dto.MediaModule{},
			gorm.ErrRecordNotFound,
			false,
			dto.MediaModule{},
			http.StatusInternalServerError,
			"fail get media module by id",
		},
		{
			"fail get media module by id",
			"GET",
			"abcde",

			dto.MediaModule{},
			gorm.ErrRecordNotFound,
			false,
			dto.MediaModule{},
			http.StatusInternalServerError,
			"fail get media module by id",
		},
	}

	for _, v := range testCase {
		mockCall := s.mock.On("GetMediaModuleByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {

			// Create request
			r := httptest.NewRequest(v.Method, "/media_module/get_by_id/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/media_module/get_by_id/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.mediaModuleController.GetMediaModuleByID(ctx)
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

func (s *suiteMediaModule) TestGetAllMediaModule() {
	testCase := []struct {
		Name               string
		Method             string
		MockReturnBody     []dto.MediaModule
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.MediaModule
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get all media module",
			"GET",

			[]dto.MediaModule{
				{
					ID:        "abcde",
					Url:       "tes",
					ModuleID:  "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
				},
				{
					ID:        "abcde",
					Url:       "tes",
					ModuleID:  "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
				},
			},
			nil,
			true,
			[]dto.MediaModule{
				{
					ID:        "abcde",
					Url:       "tes",
					ModuleID:  "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
				},
				{
					ID:        "abcde",
					Url:       "tes",
					ModuleID:  "abcde",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DeletedAt: gorm.DeletedAt{},
				},
			},
			http.StatusOK,
			"success get all media module",
		},
		{
			"fail get all media module",
			"GET",

			[]dto.MediaModule{},
			errors.New("error"),
			false,
			[]dto.MediaModule{},
			http.StatusInternalServerError,
			"fail get all media module",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllMediaModule").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/media_module/get_all", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/media_module/get_all")

			err := s.mediaModuleController.GetAllMediaModule(ctx)
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

func (s *suiteMediaModule) TestUpdateMediaModule() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.MediaModuleTransaction
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update media module",
			"POST",
			dto.MediaModuleTransaction{
				ID:       "abcde",
				Url:      "tes",
				ModuleID: "abcde",
			},
			"abcde",
			nil,
			http.StatusOK,
			"success update media module",
		},
		{
			"fail bind data",
			"POST",
			dto.MediaModuleTransaction{
				ID:       "abcde",
				Url:      "tes",
				ModuleID: "abcde",
			},
			"abcde",
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"fail update media module",
			"POST",
			dto.MediaModuleTransaction{
				ID:       "abcde",
				Url:      "tes",
				ModuleID: "abcde",
			},
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail update media module",
		},
		{
			"fail update media module",
			"POST",
			dto.MediaModuleTransaction{
				ID:       "abcde",
				Url:      "tes",
				ModuleID: "abcde",
			},
			"abcde",
			errors.New("fail update media module"),
			http.StatusInternalServerError,
			"fail update media module",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateMediaModule", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/media_module/update/"+v.ParamID, bytes.NewBuffer(res))
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
			ctx.SetPath("/media_module/update/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.mediaModuleController.UpdateMediaModule(ctx)
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

func TestSuiteMediaModule(t *testing.T) {
	suite.Run(t, new(suiteMediaModule))
}
