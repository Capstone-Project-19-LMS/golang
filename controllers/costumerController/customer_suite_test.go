package costumerController

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	customerMockService "golang/service/costumerService/customerMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteCustomer struct {
	suite.Suite
	customerController *CostumerController
	mock               *customerMockService.CustomerMock
}

func (s *suiteCustomer) SetupTest() {
	mock := &customerMockService.CustomerMock{}
	s.mock = mock
	s.customerController = &CostumerController{
		CostumerService: s.mock,
	}
}

func (s *suiteCustomer) TestCreateCustomer() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CostumerRegister
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create user",
			"POST",
			dto.CostumerRegister{
				ID:             "abcde",
				Name:           "tes",
				Email:          "tes@gmail.com",
				Password:       "tes123",
				ProfileImage:   "tes",
				CustomerCodeID: "abcde",
			},
			nil,
			http.StatusOK,
			"success create user",
		},
		{
			"fail bind data",
			"POST",
			dto.CostumerRegister{
				Name:           "tes",
				Email:          "tes@gmail.com",
				Password:       "tes123",
				ProfileImage:   "tes",
				CustomerCodeID: "abcde",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.CostumerRegister{
				Name:  "tes",
				Email: "tes@gmail.com",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create user",
			"POST",
			dto.CostumerRegister{
				Name:           "tes",
				Email:          "tes@gmail.com",
				Password:       "tes123",
				ProfileImage:   "tes",
				CustomerCodeID: "abcde",
			},

			errors.New("fail create user"),
			http.StatusInternalServerError,
			"fail create user",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateCustomer", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/customer/", bytes.NewBuffer(res))
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
			ctx.SetPath("/customer/create")

			err := s.customerController.Register(ctx)
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

func (s *suiteCustomer) TestVerifikasiCustomer() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CustomerVerif
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success verif user",
			"POST",
			dto.CustomerVerif{
				Email: "tes@gmail.com",
				Code:  "1234",
			},
			nil,
			http.StatusOK,
			"success verif user",
		},
		{
			"fail bind data",
			"POST",
			dto.CustomerVerif{
				Email: "tes@gmail.com",
				Code:  "1234",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},

		{
			"fail verif user",
			"POST",
			dto.CustomerVerif{
				Email: "tes@gmail.com",
				Code:  "1234",
			},

			errors.New("fail verif user"),
			http.StatusInternalServerError,
			"fail verif user",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("VerifikasiCustomer", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/customer/", bytes.NewBuffer(res))
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
			ctx.SetPath("/customer/verifikasi")

			err := s.customerController.Verifikasi(ctx)
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

// func (s *suiteCustomer) TestLoginCustomer() {
// 	testCase := []struct {
// 		Name   string
// 		Method string

// 		Body               dto.CostumerLogin
// 		MockReturnBody     dto.CostumerResponse
// 		ExpectedBody       dto.CostumerResponse
// 		MockReturnError    error
// 		ExpectedStatusCode int
// 		ExpectedMesaage    string
// 	}{
// 		{
// 			"success login",
// 			"POST",
// 			dto.CostumerLogin{
// 				Email:    "tes@gmail.com",
// 				Password: "tes123",
// 			},
// 			dto.CostumerResponse{
// 				Name:  "tes",
// 				Email: "tes@gmail.com",
// 				Token: "fghi",
// 			},
// 			dto.CostumerResponse{
// 				Name:  "tes",
// 				Email: "tes@gmail.com",
// 				Token: "fghi",
// 			},
// 			nil,
// 			http.StatusOK,
// 			"success login",
// 		},

// 	}
// 	for i, v := range testCase {
// 		mockCall := s.mock.On("LoginCostumer", v.Body).Return(v.MockReturnError)
// 		s.T().Run(v.Name, func(t *testing.T) {
// 			res, _ := json.Marshal(v.Body)
// 			// Create request
// 			r := httptest.NewRequest(v.Method, "/customer/", bytes.NewBuffer(res))
// 			if i != 1 {
// 				r.Header.Set("Content-Type", "application/json")
// 			}
// 			// Create response recorder
// 			w := httptest.NewRecorder()

// 			// handler echo
// 			e := echo.New()
// 			e.Validator = &helper.CustomValidator{
// 				Validator: validator.New(),
// 			}
// 			ctx := e.NewContext(r, w)
// 			ctx.SetPath("/customer/create")

// 			err := s.customerController.Login(ctx)
// 			s.NoError(err)
// 			s.Equal(v.ExpectedStatusCode, w.Code)

// 			var resp map[string]interface{}
// 			err = json.NewDecoder(w.Result().Body).Decode(&resp)
// 			s.NoError(err)

// 			s.Equal(v.ExpectedMesaage, resp["message"])
// 		})
// 		// remove mock
// 		mockCall.Unset()
// 	}
// }

func TestSuiteCustomer(t *testing.T) {
	suite.Run(t, new(suiteCustomer))
}
