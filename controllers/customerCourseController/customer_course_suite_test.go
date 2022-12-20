package customerCourseController

import (
	"bytes"
	"encoding/json"
	"errors"
	middlewareCustomer "golang/app/middlewares/costumer"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/customerCourseService/customerCourseMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCustomerCourse struct {
	suite.Suite
	customerCourseController *CustomerCourseController
	mock                     *customerCourseMockService.CustomerCourseMock
}

func (s *suiteCustomerCourse) SetupTest() {
	mock := &customerCourseMockService.CustomerCourseMock{}
	s.mock = mock
	s.customerCourseController = &CustomerCourseController{
		CustomerCourseService: s.mock,
	}
}

func (s *suiteCustomerCourse) TestDeleteCustomerCourse() {
	testCase := []struct {
		Name               string
		Method             string
		CourseID           string
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success delete customer course",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			nil,
			http.StatusOK,
			"success delete customer course",
		},
		{
			"fail delete customer course because customer course not found",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail delete customer course",
		},
		{
			"fail delete customer course because error",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			errors.New("fail delete customer course"),
			http.StatusInternalServerError,
			"fail delete customer course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteCustomerCourse", v.CourseID, v.User.ID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/customerCourse/"+v.CourseID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/customerCourse/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.CourseID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.customerCourseController.DeleteCustomerCourse(ctx)
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

func (s *suiteCustomerCourse) TestGetHistoryCourseByCustomerID() {
	testCase := []struct {
		Name               string
		Method             string
		User               dto.User
		MockReturnBody     []dto.GetCourse
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.GetCourse
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get history course by customer",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.GetCourse{
				{
					ID:                 "test1",
					Name:               "test1",
					Description:        "test1",
					Price:              10000,
					Discount:           10000,
					Capacity:           10,
					InstructorID:       "insructor1",
					Rating:             4.5,
					Favorite:           true,
					StatusEnroll:       true,
					NumberOfModules:    10,
					AmountCustomer:     5,
					ProgressModule:     6,
					ProgressPercentage: 50,
					IsFinish:           false,
				},
				{
					ID:                 "test2",
					Name:               "test2",
					Description:        "test2",
					Price:              10000,
					Discount:           10000,
					Capacity:           10,
					InstructorID:       "insructor1",
					Rating:             4.5,
					Favorite:           true,
					StatusEnroll:       true,
					NumberOfModules:    10,
					AmountCustomer:     5,
					ProgressModule:     6,
					ProgressPercentage: 50,
					IsFinish:           false,
				},
			},
			nil,
			true,
			[]dto.GetCourse{
				{
					ID:                 "test1",
					Name:               "test1",
					Description:        "test1",
					Price:              10000,
					Discount:           10000,
					Capacity:           10,
					InstructorID:       "insructor1",
					Rating:             4.5,
					Favorite:           true,
					StatusEnroll:       true,
					NumberOfModules:    10,
					AmountCustomer:     5,
					ProgressModule:     6,
					ProgressPercentage: 50,
					IsFinish:           false,
				},
				{
					ID:                 "test2",
					Name:               "test2",
					Description:        "test2",
					Price:              10000,
					Discount:           10000,
					Capacity:           10,
					InstructorID:       "insructor1",
					Rating:             4.5,
					Favorite:           true,
					StatusEnroll:       true,
					NumberOfModules:    10,
					AmountCustomer:     5,
					ProgressModule:     6,
					ProgressPercentage: 50,
					IsFinish:           false,
				},
			},
			http.StatusOK,
			"success get history course",
		},
		{
			"fail get history course by customer",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.GetCourse{},
			errors.New("error"),
			false,
			[]dto.GetCourse{},
			http.StatusInternalServerError,
			"fail get history course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetHistoryCourseByCustomerID", v.User.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/courses", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/courses")
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.customerCourseController.GetHistoryCourseByCustomerID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				s.Equal(v.ExpectedBody[0].Name, resp["courses"].([]interface{})[0].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody[0].Description, resp["courses"].([]interface{})[0].(map[string]interface{})["description"])
				s.Equal(v.ExpectedBody[1].Name, resp["courses"].([]interface{})[1].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody[1].Description, resp["courses"].([]interface{})[1].(map[string]interface{})["description"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCustomerCourse) TestTakeCourse() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CustomerCourseTransaction
		User               dto.User
		CourseID           string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success take course",
			"POST",
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			nil,
			http.StatusOK,
			"success take course",
		},
		{
			"There is an empty field",
			"POST",
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail take course",
			"POST",
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			errors.New("fail take course"),
			http.StatusInternalServerError,
			"fail take course",
		},
		{
			"fail take course",
			"POST",
			dto.CustomerCourseTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			errors.New(constantError.ErrorCourseCapacity),
			http.StatusBadRequest,
			"fail take course",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("TakeCourse", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/enroll", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/enroll/:courseId")
			if i != 1 {
				ctx.SetParamNames("courseId")
				ctx.SetParamValues(v.CourseID)
			}
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.customerCourseController.TakeCourse(ctx)
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

func (s *suiteCustomerCourse) TestUpdateEnrollmentStatus() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CustomerCourseTransaction
		ParamID            string
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update enrollment status",
			"POST",
			dto.CustomerCourseTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success update enrollment status",
		},
		{
			"fail bind data",
			"POST",
			dto.CustomerCourseTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusBadRequest,
			"fail bind request body",
		},
		{
			"There is an empty field",
			"POST",
			dto.CustomerCourseTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail update enrollment status",
			"POST",
			dto.CustomerCourseTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status:     true,
				NoModule:   0,
				IsFinish:   false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail update enrollment status",
		},
		{
			"fail update enrollment status",
			"POST",
			dto.CustomerCourseTransaction{
				ID: 		"abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				Status: true,
				NoModule: 0,
				IsFinish: false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail update enrollment status"),
			http.StatusInternalServerError,
			"fail update enrollment status",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateEnrollmentStatus", v.Body, v.User.ID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/course/enroll/"+v.ParamID, bytes.NewBuffer(res))
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
			ctx.SetPath("/course/enroll/:courseId")
			if i != 2 {
				ctx.SetParamNames("courseId")
				ctx.SetParamValues(v.ParamID)
			}
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.customerCourseController.UpdateEnrollmentStatus(ctx)
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

func TestSuiteCustomerCourse(t *testing.T) {
	suite.Run(t, new(suiteCustomerCourse))
}
