package customerAssignmentService

import (
	"errors"
	"golang/models/dto"
	customerAssignmentMockrepository "golang/repository/customerAssignmentRepository/customerAssignmentMockRepository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCustomerAssignment struct {
	suite.Suite
	customerAssignmentService CustomerAssignmentService
	mock                      *customerAssignmentMockrepository.CustomerAssignmentMock
}

func (s *suiteCustomerAssignment) SetupTest() {
	mock := &customerAssignmentMockrepository.CustomerAssignmentMock{}
	s.mock = mock
	NewcustomerAssignmentService := NewcustomerAssignmentService(s.mock)
	s.customerAssignmentService = NewcustomerAssignmentService
}

func (s *suiteCustomerAssignment) TestCreateCustomerAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		Body            dto.CustomerAssignmentTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success create customer assignment",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			nil,
			false,
			nil,
		},
		{
			"fail create customer assignment",
			dto.User{
				ID:   "1",
				Role: "customer",
			},
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "abcde",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("CreateCustomerAssignment", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.customerAssignmentService.CreateCustomerAssignment(v.Body)
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

func (s *suiteCustomerAssignment) TestDeleteCustomerAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success delete customer assignment",
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
			"fail delete customer assignment",
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
		mockCall := s.mock.On("DeleteCustomerAssignment", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.customerAssignmentService.DeleteCustomerAssignment(v.ParamID)
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

func (s *suiteCustomerAssignment) TestGetCustomerAssignmentByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		MockReturnBody  dto.CustomerAssignmentAcc
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.CustomerAssignmentAcc
		ExpectedError   error
	}{
		{
			"success get customer assignment by id ",
			"abcde",
			dto.CustomerAssignmentAcc{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "avcde",
				Customer: struct {
					Name string "json:\"name\" gorm:\"notNull;size:255\""
				}{
					Name: "tes",
				},
			},
			nil,
			true,
			dto.CustomerAssignmentAcc{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "avcde",
				Customer: struct {
					Name string "json:\"name\" gorm:\"notNull;size:255\""
				}{
					Name: "tes",
				},
			},
			nil,
		},
		{
			"failed get customer assignment by id",
			"abcde",

			dto.CustomerAssignmentAcc{},
			gorm.ErrRecordNotFound,
			false,
			dto.CustomerAssignmentAcc{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetCustomerAssignmentByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			customerAssignment, err := s.customerAssignmentService.GetCustomerAssignmentByID(v.ParamID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, customerAssignment)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, customerAssignment)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCustomerAssignment) TestGetAllCustomerAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		MockReturnBody  []dto.CustomerAssignmentAcc
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    []dto.CustomerAssignmentAcc
		ExpectedError   error
	}{
		{
			"success get all customer assignment",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.CustomerAssignmentAcc{
				{
					ID:           "abcde",
					File:         "tes",
					Grade:        1,
					AssignmentID: "abcde",
					CustomerID:   "avcde",
				},
				{
					ID:           "abcde",
					File:         "tes",
					Grade:        1,
					AssignmentID: "abcde",
					CustomerID:   "avcde",
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
					CustomerID:   "avcde",
				},
				{
					ID:           "abcde",
					File:         "tes",
					Grade:        1,
					AssignmentID: "abcde",
					CustomerID:   "avcde",
				},
			},
			nil,
		},
		{
			"failed get all customer assignment",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.CustomerAssignmentAcc{},
			gorm.ErrRecordNotFound,
			false,
			nil,
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllCustomerAssignment").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			CustomerAssignment, err := s.customerAssignmentService.GetAllCustomerAssignment()
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, CustomerAssignment)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, CustomerAssignment)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCustomerAssignment) TestUpdateCustomerAssignment() {
	testCase := []struct {
		Name            string
		User            dto.User
		ParamID         string
		Body            dto.CustomerAssignmentTransaction
		MockReturnError error
		HasReturnError  bool
		ExpectedError   error
	}{
		{
			"success update customer assignment",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "avcde",
			},
			nil,
			false,
			nil,
		},
		{
			"fail update customer assignment",
			dto.User{
				ID:   "1",
				Role: "user",
			},
			"abcde",
			dto.CustomerAssignmentTransaction{
				ID:           "abcde",
				File:         "tes",
				Grade:        1,
				AssignmentID: "abcde",
				CustomerID:   "avcde",
			},
			errors.New("error"),
			true,
			errors.New("error"),
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("UpdateCustomerAssignment", mock.Anything).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			err := s.customerAssignmentService.UpdateCustomerAssignment(v.Body)
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

func TestSuiteCustomerAssignment(t *testing.T) {
	suite.Run(t, new(suiteCustomerAssignment))
}
