package assignmentcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	"golang/service/assignmentService/assignmentMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteAssignment struct {
	suite.Suite
	assignmentController *AssignmentController
	mock                 *assignmentMockService.AssignmentMock
}

func (s *suiteAssignment) SetupTest() {
	mock := &assignmentMockService.AssignmentMock{}
	s.mock = mock
	s.assignmentController = &AssignmentController{
		AssignmentService: s.mock,
	}
}

func (s *suiteAssignment) TestCreateAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.AssignmentTransaction
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create assignment",
			"POST",
			dto.AssignmentTransaction{
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			nil,
			http.StatusOK,
			"success create assignment",
		},
		{
			"fail bind data",
			"POST",
			dto.AssignmentTransaction{
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.AssignmentTransaction{
				Title: "tes",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create assignment",
			"POST",
			dto.AssignmentTransaction{
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},

			errors.New("Fail create assignment"),
			http.StatusInternalServerError,
			"fail create assignment",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateAssignment", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/assignment/", bytes.NewBuffer(res))
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
			ctx.SetPath("/assignment/create")

			err := s.assignmentController.CreateAssignment(ctx)
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

func (s *suiteAssignment) TestDeleteAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success delete assignment",
			"DELETE",
			"abcde",
			nil,
			http.StatusOK,
			"success delete assignment",
		},
		{
			"fail delete assignment",
			"DELETE",
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail delete assignment",
		},
		{
			"fail delete assignment",
			"DELETE",
			"abcde",

			errors.New("fail delete assignment"),
			http.StatusInternalServerError,
			"fail delete assignment",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteAssignment", v.ParamID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/assignment/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/assignment/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.assignmentController.DeleteAssignment(ctx)
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

func (s *suiteAssignment) TestGetAssignmentByID() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnBody     dto.Assignment
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.Assignment
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get assignment by id",
			"GET",
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
						AssignmentID: "tes",
						CustomerID:   "tes",
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
						AssignmentID: "tes",
						CustomerID:   "abcde",
					},
				},
			},

			http.StatusOK,
			"success get assignment by id",
		},
		{
			"failed get assignment by id",
			"GET",
			"abcde",
			dto.Assignment{},
			gorm.ErrRecordNotFound,
			false,
			dto.Assignment{},
			http.StatusInternalServerError,
			"failed get assignment by id",
		},
	}

	for _, v := range testCase {
		mockCall := s.mock.On("GetAssignmentByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/assignment/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/assignment/get_by_id/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.assignmentController.GetAssignmentByID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)
			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				s.Equal(v.ExpectedBody.Title, resp["assignment"].(map[string]interface{})["title"])
				s.Equal(v.ExpectedBody.Description, resp["assignment"].(map[string]interface{})["description"])

			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteAssignment) TestGetAllAssignment() {
	testCase := []struct {
		Name   string
		Method string

		MockReturnBody     []dto.Assignment
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.Assignment
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get all assignment",
			"GET",

			[]dto.Assignment{
				{
					ID:          "abcde1",
					Title:       "tes1",
					Description: "tes1",
					ModuleID:    "abcde1",
					CustomerAssignments: []dto.CustomerAssignment{
						{
							ID:           "abcde1",
							File:         "tes1",
							Grade:        1,
							AssignmentID: "tes1",
							CustomerID:   "tes1",
						},
					},
				},
				{
					ID:          "abcde2",
					Title:       "tes2",
					Description: "tes2",
					ModuleID:    "abcde2",
					CustomerAssignments: []dto.CustomerAssignment{
						{
							ID:           "abcde2",
							File:         "tes2",
							Grade:        2,
							AssignmentID: "tes2",
							CustomerID:   "tes2",
						},
					},
				},
			},
			nil,
			true,
			[]dto.Assignment{
				{},
				{},
			},
			http.StatusOK,
			"success get all assignment",
		},
		{
			"fail get all assignment",
			"GET",

			[]dto.Assignment{},
			errors.New("error"),
			false,
			[]dto.Assignment{},
			http.StatusInternalServerError,
			"fail get all assignment",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllAssignment").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/assignment", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/assignment")

			err := s.assignmentController.GetAllAssignment(ctx)
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

func (s *suiteAssignment) TestUpdateAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.AssignmentTransaction
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update assignment",
			"POST",
			dto.AssignmentTransaction{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			"abcde",
			nil,
			http.StatusOK,
			"success update assignment",
		},
		{
			"fail bind data",
			"POST",
			dto.AssignmentTransaction{
				Title:       "tes",
				Description: "tes",
			},
			"abcde",
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"fail update assignment",
			"POST",
			dto.AssignmentTransaction{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail update assignment",
		},
		{
			"fail update assignment",
			"POST",
			dto.AssignmentTransaction{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
			},
			"abcde",
			errors.New("fail update assignment"),
			http.StatusInternalServerError,
			"fail update assignment",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateAssignment", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/assignment/update/"+v.ParamID, bytes.NewBuffer(res))
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
			ctx.SetPath("/assignment/update/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.assignmentController.UpdateAssignment(ctx)
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

func TestSuiteAssignment(t *testing.T) {
	suite.Run(t, new(suiteAssignment))
}
