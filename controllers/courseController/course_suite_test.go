package courseController

import (
	"bytes"
	"encoding/json"
	"errors"
	middlewareCustomer "golang/app/middlewares/costumer"
	middlewareInstructor "golang/app/middlewares/instructor"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/courseService/courseMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCourse struct {
	suite.Suite
	courseController *CourseController
	mock             *courseMockService.CourseMock
}

func (s *suiteCourse) SetupTest() {
	mock := &courseMockService.CourseMock{}
	s.mock = mock
	s.courseController = &CourseController{
		CourseService: s.mock,
	}
}

func (s *suiteCourse) TestCreateCourse() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CourseTransaction
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create course",
			"POST",
			dto.CourseTransaction{
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				CategoryID:  "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success create course",
		},
		{
			"fail bind data",
			"POST",
			dto.CourseTransaction{
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				CategoryID:  "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.CourseTransaction{
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				CategoryID:  "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create course",
			"POST",
			dto.CourseTransaction{
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				CategoryID:  "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail create course"),
			http.StatusInternalServerError,
			"fail create course",
		},
		{
			"fail create course",
			"POST",
			dto.CourseTransaction{
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    -1,
				CategoryID:  "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New(constantError.ErrorCapacityLowerThanZero),
			http.StatusBadRequest,
			"fail create course",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateCourse", v.Body, v.User).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/course", bytes.NewBuffer(res))
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
			ctx.SetPath("/course")
			ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.courseController.CreateCourse(ctx)
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

func (s *suiteCourse) TestDeleteCourse() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success delete course",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success delete course",
		},
		{
			"fail delete course because course not found",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail delete course",
		},
		{
			"fail delete course because error",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail delete course"),
			http.StatusInternalServerError,
			"fail delete course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteCourse", v.ParamID, v.User.ID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/course/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/course/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.courseController.DeleteCourse(ctx)
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

func (s *suiteCourse) TestGetAllCourse() {
	testCase := []struct {
		Name                   string
		Method                 string
		User                   dto.User
		MockReturnBody         []dto.GetCourse
		MockReturnError        error
		HasReturnBody          bool
		ExpectedBodyCustomer   []dto.GetCourse
		ExpectedBodyInstructor []dto.GetCourseInstructor
		ExpectedStatusCode     int
		ExpectedMesaage        string
	}{
		{
			"success get all course by customer",
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
			[]dto.GetCourseInstructor{
				{
					ID:              "test1",
					Name:            "test1",
					Description:     "test1",
					Price:           10000,
					Discount:        10000,
					Capacity:        10,
					InstructorID:    "insructor1",
					Rating:          4.5,
					NumberOfModules: 10,
					AmountCustomer:  5,
				},
				{
					ID:              "test2",
					Name:            "test2",
					Description:     "test2",
					Price:           10000,
					Discount:        10000,
					Capacity:        10,
					InstructorID:    "insructor1",
					Rating:          4.5,
					NumberOfModules: 10,
					AmountCustomer:  5,
				},
			},
			http.StatusOK,
			"success get all courses",
		},
		{
			"success get all course by instructor",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
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
			[]dto.GetCourseInstructor{
				{
					ID:              "test1",
					Name:            "test1",
					Description:     "test1",
					Price:           10000,
					Discount:        10000,
					Capacity:        10,
					InstructorID:    "insructor1",
					Rating:          4.5,
					NumberOfModules: 10,
					AmountCustomer:  5,
				},
				{
					ID:              "test2",
					Name:            "test2",
					Description:     "test2",
					Price:           10000,
					Discount:        10000,
					Capacity:        10,
					InstructorID:    "insructor1",
					Rating:          4.5,
					NumberOfModules: 10,
					AmountCustomer:  5,
				},
			},
			http.StatusOK,
			"success get all courses",
		},
		{
			"fail get all course by customer",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.GetCourse{},
			errors.New("error"),
			false,
			[]dto.GetCourse{},
			[]dto.GetCourseInstructor{},
			http.StatusInternalServerError,
			"fail get all courses",
		},
		{
			"fail get all course by instructor",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			[]dto.GetCourse{},
			errors.New("error"),
			false,
			[]dto.GetCourse{},
			[]dto.GetCourseInstructor{},
			http.StatusInternalServerError,
			"fail get all courses",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllCourse", v.User).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/courses", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/courses")
			if v.User.Role == "customer" {
				ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})
			} else {
				ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.User.ID, Role: v.User.Role}})
			}

			err := s.courseController.GetAllCourse(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				if v.User.Role == "customer" {
					s.Equal(v.ExpectedBodyCustomer[0].Name, resp["courses"].([]interface{})[0].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyCustomer[0].Description, resp["courses"].([]interface{})[0].(map[string]interface{})["description"])
					s.Equal(v.ExpectedBodyCustomer[1].Name, resp["courses"].([]interface{})[1].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyCustomer[1].Description, resp["courses"].([]interface{})[1].(map[string]interface{})["description"])
				} else if v.User.Role == "instructor" {
					s.Equal(v.ExpectedBodyInstructor[0].Name, resp["courses"].([]interface{})[0].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyInstructor[0].Description, resp["courses"].([]interface{})[0].(map[string]interface{})["description"])
					s.Equal(v.ExpectedBodyInstructor[1].Name, resp["courses"].([]interface{})[1].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyInstructor[1].Description, resp["courses"].([]interface{})[1].(map[string]interface{})["description"])
				}
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCourse) TestGetCourseByID() {
	testCase := []struct {
		Name                   string
		Method                 string
		ParamID                string
		ParamUser              dto.User
		MockReturnBody         dto.GetCourseByID
		MockReturnError        error
		HasReturnBody          bool
		ExpectedBodyCustomer   dto.GetCourseByID
		ExpectedBodyInstructor dto.GetCourseInstructorByID
		ExpectedStatusCode     int
		ExpectedMesaage        string
	}{
		{
			"success get course by id and customer",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.GetCourseByID{
				ID:                 "test2",
				Name:               "test2",
				Description:        "test2",
				Price:              10000,
				Discount:           10000,
				Capacity:           10,
				InstructorID:       "insructor1",
				Rating:             4.5,
				NumberOfModules:    10,
				ProgressModule:     6,
				ProgressPercentage: 50,
				IsFinish:           false,
			},
			nil,
			true,
			dto.GetCourseByID{
				ID:                 "test2",
				Name:               "test2",
				Description:        "test2",
				Price:              10000,
				Discount:           10000,
				Capacity:           10,
				InstructorID:       "insructor1",
				Rating:             4.5,
				NumberOfModules:    10,
				ProgressModule:     6,
				ProgressPercentage: 50,
				IsFinish:           false,
			},
			dto.GetCourseInstructorByID{},
			http.StatusOK,
			"success get course by id",
		},
		{
			"success get course by id and instructor",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			dto.GetCourseByID{
				ID:                 "test2",
				Name:               "test2",
				Description:        "test2",
				Price:              10000,
				Discount:           10000,
				Capacity:           10,
				InstructorID:       "insructor1",
				Rating:             4.5,
				NumberOfModules:    10,
				ProgressModule:     6,
				ProgressPercentage: 50,
				IsFinish:           false,
			},
			nil,
			true,
			dto.GetCourseByID{},
			dto.GetCourseInstructorByID{
				ID:              "test2",
				Name:            "test2",
				Description:     "test2",
				Price:           10000,
				Discount:        10000,
				Capacity:        10,
				InstructorID:    "insructor1",
				Rating:          4.5,
				NumberOfModules: 10,
			},
			http.StatusOK,
			"success get course by id",
		},
		{
			"failed get course by id",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.GetCourseByID{},
			errors.New("error"),
			false,
			dto.GetCourseByID{},
			dto.GetCourseInstructorByID{},
			http.StatusInternalServerError,
			"fail get course by id",
		},
		{
			"failed get course by id",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.GetCourseByID{},
			gorm.ErrRecordNotFound,
			false,
			dto.GetCourseByID{},
			dto.GetCourseInstructorByID{},
			http.StatusNotFound,
			"fail get course by id",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetCourseByID", v.ParamID, v.ParamUser).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/courses/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/courses/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)
			if v.ParamUser.Role == "customer" {
				ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.ParamUser.ID, Role: v.ParamUser.Role}})
			} else {
				ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.ParamUser.ID, Role: v.ParamUser.Role}})
			}

			err := s.courseController.GetCourseByID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				if v.ParamUser.Role == "customer" {
					s.Equal(v.ExpectedBodyCustomer.Name, resp["course"].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyCustomer.Description, resp["course"].(map[string]interface{})["description"])
				} else {
					s.Equal(v.ExpectedBodyInstructor.Name, resp["course"].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyInstructor.Description, resp["course"].(map[string]interface{})["description"])
				}
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCourse) TestGetCourseEnrollByID() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		ParamUser          dto.User
		MockReturnBody     []dto.CustomerCourseEnroll
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.CustomerCourseEnroll
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get course with customer enrolled",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			[]dto.CustomerCourseEnroll{
				{
					ID:           "test1",
					Name:         "test1",
					Email:        "test1@gmail.com",
					ProfileImage: "test1.jpg",
					StatusEnroll: true,
				},
				{
					ID:           "test2",
					Name:         "test2",
					Email:        "test2@gmail.com",
					ProfileImage: "test2.jpg",
					StatusEnroll: true,
				},
			},
			nil,
			true,
			[]dto.CustomerCourseEnroll{
				{
					ID:           "test1",
					Name:         "test1",
					Email:        "test1@gmail.com",
					ProfileImage: "test1.jpg",
					StatusEnroll: true,
				},
				{
					ID:           "test2",
					Name:         "test2",
					Email:        "test2@gmail.com",
					ProfileImage: "test2.jpg",
					StatusEnroll: true,
				},
			},
			http.StatusOK,
			"success get course with customer enrolled",
		},
		{
			"fail get course with customer enrolled",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			[]dto.CustomerCourseEnroll{},
			errors.New("fail get course with customer enrolled"),
			false,
			[]dto.CustomerCourseEnroll{},
			http.StatusInternalServerError,
			"fail get course with customer enrolled",
		},
		{
			"fail get course with customer enrolled",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			[]dto.CustomerCourseEnroll{},
			errors.New(constantError.ErrorCourseNotFound),
			false,
			[]dto.CustomerCourseEnroll{},
			http.StatusNotFound,
			"fail get course with customer enrolled",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetCourseEnrollByID", v.ParamID, v.ParamUser).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/courses/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/courses/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.ParamID)
			if v.ParamUser.Role == "customer" {
				ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.ParamUser.ID, Role: v.ParamUser.Role}})
			} else {
				ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.ParamUser.ID, Role: v.ParamUser.Role}})
			}

			err := s.courseController.GetCourseEnrollByID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				s.Equal(v.ExpectedBody[0].Name, resp["customer_enroll"].([]interface{})[0].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody[0].Email, resp["customer_enroll"].([]interface{})[0].(map[string]interface{})["email"])
				s.Equal(v.ExpectedBody[0].ProfileImage, resp["customer_enroll"].([]interface{})[0].(map[string]interface{})["profile_image"])
				s.Equal(v.ExpectedBody[0].StatusEnroll, resp["customer_enroll"].([]interface{})[0].(map[string]interface{})["status_enroll"])
				s.Equal(v.ExpectedBody[1].Name, resp["customer_enroll"].([]interface{})[1].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody[1].Email, resp["customer_enroll"].([]interface{})[1].(map[string]interface{})["email"])
				s.Equal(v.ExpectedBody[1].ProfileImage, resp["customer_enroll"].([]interface{})[1].(map[string]interface{})["profile_image"])
				s.Equal(v.ExpectedBody[1].StatusEnroll, resp["customer_enroll"].([]interface{})[1].(map[string]interface{})["status_enroll"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCourse) TestUpdateCourse() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CourseTransaction
		ParamID            string
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update course",
			"POST",
			dto.CourseTransaction{
				ID: 		"abcde",
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				InstructorID:  "abcde",
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success update course",
		},
		{
			"fail bind data",
			"POST",
			dto.CourseTransaction{
				ID: 		"abcde",
						Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				InstructorID:  "abcde",
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
			"fail update course",
			"POST",
			dto.CourseTransaction{
				ID: 		"abcde",
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				InstructorID:  "abcde",
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail update course",
		},
		{
			"fail update course",
			"POST",
			dto.CourseTransaction{
				ID: 		"abcde",
				Name:        "test",
				Description: "test",
				Objective:   "test",
				Capacity:    10,
				InstructorID:  "abcde",
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail update course"),
			http.StatusInternalServerError,
			"fail update course",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateCourse", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/categories/"+v.ParamID, bytes.NewBuffer(res))
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
			ctx.SetPath("/categories/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.courseController.UpdateCourse(ctx)
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

func TestSuiteCourse(t *testing.T) {
	suite.Run(t, new(suiteCourse))
}
