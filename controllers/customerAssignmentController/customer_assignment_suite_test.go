package customerassignmentcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	customerAssignmentMockservice "golang/service/customerAssignmentService/customerAssignmentMockService"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCustomerAssignment struct {
	suite.Suite
	customerAssignmentController *CustomerAssignmentController
	mock                         *customerAssignmentMockservice.CustomerAssignmentMock
}

func (s *suiteCustomerAssignment) SetupTest() {
	mock := &customerAssignmentMockservice.CustomerAssignmentMock{}
	s.mock = mock
	s.customerAssignmentController = &CustomerAssignmentController{
		CustomerAssignmentService: s.mock,
	}
}

func (s *suiteCustomerAssignment) TestCreateCustomerAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CustomerAssignmentTransaction
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create customer assignment",
			"POST",
			dto.CustomerAssignmentTransaction{
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			nil,
			http.StatusOK,
			"success create customer assignment",
		},
		{
			"fail bind data",
			"POST",
			dto.CustomerAssignmentTransaction{
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.CustomerAssignmentTransaction{
				File: "tes",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create customer assignment",
			"POST",
			dto.CustomerAssignmentTransaction{
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},

			errors.New("fail create customer assignment"),
			http.StatusInternalServerError,
			"fail create customer assignment",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateCustomerAssignment", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/customer_assignment/", bytes.NewBuffer(res))
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
			ctx.SetPath("/customer_assignment/create")

			err := s.customerAssignmentController.CreateCustomerAssignment(ctx)
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

func (s *suiteCustomerAssignment) TestDeleteCustomerAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success delete customer assignment",
			"DELETE",
			"abcde",
			nil,
			http.StatusOK,
			"success delete customer assignment",
		},
		{
			"fail delete customer assignment",
			"DELETE",
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail delete customer assignment",
		},
		{
			"fail delete customer assignment",
			"DELETE",
			"abcde",
			errors.New("fail delete customer assignment"),
			http.StatusInternalServerError,
			"fail delete customer assignment",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteCustomerAssignment", v.ParamID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/customer_assignment/delete/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/customer_assignment/delete/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.customerAssignmentController.DeleteCustomerAssignment(ctx)
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

func (s *suiteCustomerAssignment) TestGetCustomerAssignmentByID() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		MockReturnBody     dto.CustomerAssignmentAcc
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.CustomerAssignmentAcc
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get customer assignment by id",
			"GET",
			"abcde",

			dto.CustomerAssignmentAcc{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				DeletedAt:    gorm.DeletedAt{},
			},
			nil,
			true,
			dto.CustomerAssignmentAcc{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
				DeletedAt:    gorm.DeletedAt{},
			},
			http.StatusOK,
			"success get customer assignment by id",
		},
		{
			"fail get customer assignment by id",
			"GET",
			"abcde",

			dto.CustomerAssignmentAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.CustomerAssignmentAcc{},
			http.StatusInternalServerError,
			"fail get customer assignment by id",
		},
		{
			"fail get customer assignment by id",
			"GET",
			"abcde",

			dto.CustomerAssignmentAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.CustomerAssignmentAcc{},
			http.StatusInternalServerError,
			"fail get customer assignment by id",
		},
	}

	for _, v := range testCase {
		mockCall := s.mock.On("GetCustomerAssignmentByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {

			// Create request
			r := httptest.NewRequest(v.Method, "/customer_assignment/get_by_id/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/customer_assignment/get_by_id/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.customerAssignmentController.GetCustomerAssignmentByID(ctx)
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

func (s *suiteCustomerAssignment) TestGetAllCustomerAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		MockReturnBody     []dto.CustomerAssignmentAcc
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.CustomerAssignmentAcc
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get all customer assignment",
			"GET",

			[]dto.CustomerAssignmentAcc{
				{
					ID:           "abcde",
					File:         "tes",
					Grade:        1,
					AssignmentID: "abcde",
					CustomerID:   "abcde",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    gorm.DeletedAt{},
				},
				{
					ID:           "abcdef",
					File:         "tes 2",
					Grade:        2,
					AssignmentID: "abcdef",
					CustomerID:   "abcdef",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    gorm.DeletedAt{},
				},
			},
			nil,
			true,
			[]dto.CustomerAssignmentAcc{
				{
					ID:           "abcde",
					File:         "tes",
					Grade:        1,
					AssignmentID: "abcde",
					CustomerID:   "abcde",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    gorm.DeletedAt{},
				},
				{
					ID:           "abcde",
					File:         "tes",
					Grade:        1,
					AssignmentID: "abcde",
					CustomerID:   "abcde",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
					DeletedAt:    gorm.DeletedAt{},
				},
			},
			http.StatusOK,
			"success get all customer assignment",
		},
		{
			"fail get all customer assignment",
			"GET",

			[]dto.CustomerAssignmentAcc{},
			errors.New("error"),
			false,
			[]dto.CustomerAssignmentAcc{},
			http.StatusInternalServerError,
			"fail get all customer assignment",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllCustomerAssignment").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/customer_assignment/get_all", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/customer_assignment/get_all")

			err := s.customerAssignmentController.GetAllCustomerAssignment(ctx)
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

func (s *suiteCustomerAssignment) TestUpdateCustomerAssignment() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CustomerAssignmentTransaction
		ParamID            string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update customer assignment",
			"POST",
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			"abcde",
			nil,
			http.StatusOK,
			"success update customer assignment",
		},
		{
			"fail bind data",
			"POST",
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			"abcde",
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"fail update customer assignment",
			"POST",
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			"abcde",
			gorm.ErrRecordNotFound,
			http.StatusInternalServerError,
			"fail update customer assignment",
		},
		{
			"fail update customer assignment",
			"POST",
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			"abcde",
			errors.New("fail update customer assignment"),
			http.StatusInternalServerError,
			"fail update customer assignment",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateCustomerAssignment", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/customer_assignment/update/"+v.ParamID, bytes.NewBuffer(res))
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
			ctx.SetPath("/customer_assignment/update/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)

			err := s.customerAssignmentController.UpdateCustomerAssignment(ctx)
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

func TestSuiteCustomerAssignment(t *testing.T) {
	suite.Run(t, new(suiteCustomerAssignment))
}
