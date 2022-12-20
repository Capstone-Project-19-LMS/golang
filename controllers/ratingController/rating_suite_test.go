package ratingController

import (
	"bytes"
	"encoding/json"
	"errors"
	middlewareCustomer "golang/app/middlewares/costumer"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/ratingService/ratingMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteRating struct {
	suite.Suite
	ratingController *RatingController
	mock             *ratingMockService.RatingMock
}

func (s *suiteRating) SetupTest() {
	mock := &ratingMockService.RatingMock{}
	s.mock = mock
	s.ratingController = &RatingController{
		RatingService: s.mock,
	}
}

func (s *suiteRating) TestAddRating() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.RatingTransaction
		User               dto.User
		ParamCourseID      string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success rating course",
			"POST",
			dto.RatingTransaction{
				Rating:      5,
				Testimonial: "abcde",
				IsPublish:   true,
				CustomerID:  "abcde",
				CourseID:    "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			nil,
			http.StatusOK,
			"success rating course",
		},
		{
			"fail bind data",
			"POST",
			dto.RatingTransaction{
				Rating:      5,
				Testimonial: "abcde",
				IsPublish:   true,
				CustomerID:  "abcde",
				CourseID:    "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.RatingTransaction{
				Rating:     5,
				IsPublish:  true,
				CustomerID: "abcde",
				CourseID:   "abcde",
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
			"fail rating course",
			"POST",
			dto.RatingTransaction{
				Rating:      5,
				Testimonial: "abcde",
				IsPublish:   true,
				CustomerID:  "abcde",
				CourseID:    "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			errors.New("fail rating course"),
			http.StatusInternalServerError,
			"fail rating course",
		},
		{
			"fail rating course",
			"POST",
			dto.RatingTransaction{
				Rating:      5,
				Testimonial: "abcde",
				IsPublish:   true,
				CustomerID:  "abcde",
				CourseID:    "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			errors.New(constantError.ErrorCustomerAlreadyRatingCourse),
			http.StatusBadRequest,
			"fail rating course",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("AddRating", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			res, _ := json.Marshal(v.Body)
			r := httptest.NewRequest(v.Method, "/rating", bytes.NewBuffer(res))
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
			ctx.SetPath("/rating/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.ParamCourseID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.ratingController.AddRating(ctx)
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

func (s *suiteRating) TestDeleteRating() {
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
			"success delete rating course",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			nil,
			http.StatusOK,
			"success delete rating course",
		},
		{
			"fail delete rating course because rating course not found",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail delete rating course",
		},
		{
			"fail delete rating course because error",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			errors.New("fail delete rating course"),
			http.StatusInternalServerError,
			"fail delete rating course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteRating", v.CourseID, v.User.ID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/rating/"+v.CourseID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/rating/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.CourseID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.ratingController.DeleteRating(ctx)
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

func (s *suiteRating) TestGetCourseByCustomerID() {
	testCase := []struct {
		Name               string
		Method             string
		User               dto.User
		CourseID           string
		MockReturnBody     dto.Rating
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.Rating
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get rating course by customer",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{
				ID:          "test1",
				Rating:      5,
				Testimonial: "test1",
				IsPublish:   true,
				CustomerID:  "customer1",
				CourseID:    "course1",
			},
			nil,
			true,
			dto.Rating{
				ID:          "test1",
				Rating:      5,
				Testimonial: "test1",
				IsPublish:   true,
				CustomerID:  "customer1",
				CourseID:    "course1",
			},
			http.StatusOK,
			"success get rating course",
		},
		{
			"fail get rating course by customer",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{},
			errors.New("error"),
			false,
			dto.Rating{},
			http.StatusInternalServerError,
			"fail get rating course",
		},
		{
			"fail get rating course by customer",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			dto.Rating{},
			errors.New(constantError.ErrorCustomerNotRatingCourse),
			false,
			dto.Rating{},
			http.StatusNotFound,
			"fail get rating course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetRatingByCourseIDCustomerID", v.CourseID, v.User.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/rating", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/rating/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.CourseID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.ratingController.GetRatingByCourseIDCustomerID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				s.Equal(v.ExpectedBody.Rating, int(resp["data"].(map[string]interface{})["rating"].(float64)))
				s.Equal(v.ExpectedBody.Testimonial, resp["data"].(map[string]interface{})["testimonial"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteRating) TestGetRatingByCourseID() {
	testCase := []struct {
		Name               string
		Method             string
		User               dto.User
		CourseID           string
		MockReturnBody     []dto.Rating
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.Rating
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get rating of course",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			[]dto.Rating{
				{
					ID:          "test1",
					Rating:      5,
					Testimonial: "test1",
					IsPublish:   true,
					CustomerID:  "customer1",
					CourseID:    "course1",
				},
				{
					ID:          "test2",
					Rating:      5,
					Testimonial: "test2",
					IsPublish:   true,
					CustomerID:  "customer2",
					CourseID:    "course2",
				},
			},
			nil,
			true,
			[]dto.Rating{
				{
					ID:          "test1",
					Rating:      5,
					Testimonial: "test1",
					IsPublish:   true,
					CustomerID:  "customer1",
					CourseID:    "course1",
				},
				{
					ID:          "test2",
					Rating:      5,
					Testimonial: "test2",
					IsPublish:   true,
					CustomerID:  "customer2",
					CourseID:    "course2",
				},
			},
			http.StatusOK,
			"success get rating of course",
		},
		{
			"fail get rating of course",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			[]dto.Rating{},
			errors.New("error"),
			false,
			[]dto.Rating{},
			http.StatusInternalServerError,
			"fail get rating of course",
		},
		{
			"fail get rating of course",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			[]dto.Rating{},
			errors.New(constantError.ErrorCustomerNotRatingCourse),
			false,
			[]dto.Rating{},
			http.StatusNotFound,
			"fail get rating of course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetRatingByCourseID", v.CourseID, v.User.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/rating", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/rating/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.CourseID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.ratingController.GetRatingByCourseID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				s.Equal(v.ExpectedBody[0].Rating, int(resp["rating"].([]interface{})[0].(map[string]interface{})["rating"].(float64)))
				s.Equal(v.ExpectedBody[0].Testimonial, resp["rating"].([]interface{})[0].(map[string]interface{})["testimonial"])
				s.Equal(v.ExpectedBody[1].Rating, int(resp["rating"].([]interface{})[1].(map[string]interface{})["rating"].(float64)))
				s.Equal(v.ExpectedBody[1].Testimonial, resp["rating"].([]interface{})[1].(map[string]interface{})["testimonial"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteRating) TestUpdateRating() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.RatingTransaction
		ParamID            string
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update rating course",
			"PUT",
			dto.RatingTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				IsPublish:  false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success update rating course",
		},
		{
			"fail bind data",
			"PUT",
			dto.RatingTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				IsPublish:  false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"fail update rating course",
			"PUT",
			dto.RatingTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				IsPublish:  false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail update rating course",
		},
		{
			"fail update rating course",
			"PUT",
			dto.RatingTransaction{
				ID:         "abcde",
				CustomerID: "abcde",
				CourseID:   "abcde",
				IsPublish:  false,
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail update rating course"),
			http.StatusInternalServerError,
			"fail update rating course",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateRating", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/rating/"+v.ParamID, bytes.NewBuffer(res))
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
			ctx.SetPath("/rating/:ratingId")
			ctx.SetParamNames("ratingId")
			ctx.SetParamValues(v.ParamID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.ratingController.UpdateRating(ctx)
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

func TestSuiteRating(t *testing.T) {
	suite.Run(t, new(suiteRating))
}
