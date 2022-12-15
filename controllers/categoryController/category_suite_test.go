package categoryController

import (
	"bytes"
	"encoding/json"
	"errors"
	middlewareCustomer "golang/app/middlewares/costumer"
	middlewareInstructor "golang/app/middlewares/instructor"
	"golang/helper"
	"golang/models/dto"
	"golang/service/categoryService/categoryMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCategory struct {
	suite.Suite
	categoryController *CategoryController
	mock               *categoryMockService.CategoryMock
}

func (s *suiteCategory) SetupTest() {
	mock := &categoryMockService.CategoryMock{}
	s.mock = mock
	s.categoryController = &CategoryController{
		CategoryService: s.mock,
	}
}

func (s *suiteCategory) TestCreateCategory() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CategoryTransaction
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create category",
			"POST",
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success create category",
		},
		{
			"fail bind data",
			"POST",
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
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
			dto.CategoryTransaction{
				Description: "test",
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
			"fail create category",
			"POST",
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
			},
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail create category"),
			http.StatusInternalServerError,
			"fail create category",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateCategory", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/categories", bytes.NewBuffer(res))
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
			ctx.SetPath("/categories")
			ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.categoryController.CreateCategory(ctx)
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

func (s *suiteCategory) TestDeleteCategory() {
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
			"success delete category",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success delete category",
		},
		{
			"fail delete category because category not found",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail delete category",
		},
		{
			"fail delete category because error",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail delete category"),
			http.StatusInternalServerError,
			"fail delete category",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteCategory", v.ParamID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/categories/"+v.ParamID, nil)
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

			err := s.categoryController.DeleteCategory(ctx)
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

func (s *suiteCategory) TestGetAllCategory() {
	testCase := []struct {
		Name               string
		Method             string
		User               dto.User
		MockReturnBody     []dto.CategoryTransaction
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       []dto.CategoryTransaction
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success get all category",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.CategoryTransaction{
				{
					ID:          "test1",
					Name:        "test1",
					Description: "test1",
				},
				{
					ID:          "test2",
					Name:        "test2",
					Description: "test2",
				},
			},
			nil,
			true,
			[]dto.CategoryTransaction{
				{
					ID:          "test1",
					Name:        "test1",
					Description: "test1",
				},
				{
					ID:          "test2",
					Name:        "test2",
					Description: "test2",
				},
			},
			http.StatusOK,
			"success get all category",
		},
		{
			"fail get all category",
			"GET",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			[]dto.CategoryTransaction{},
			errors.New("error"),
			false,
			[]dto.CategoryTransaction{},
			http.StatusInternalServerError,
			"fail get all category",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAllCategory").Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/categories", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/categories")
			if v.User.Role == "customer" {
				ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})
			} else {
				ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.User.ID, Role: v.User.Role}})
			}

			err := s.categoryController.GetAllCategory(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				s.Equal(v.ExpectedBody[0].Name, resp["categories"].([]interface{})[0].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody[0].Description, resp["categories"].([]interface{})[0].(map[string]interface{})["description"])
				s.Equal(v.ExpectedBody[1].Name, resp["categories"].([]interface{})[1].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody[1].Description, resp["categories"].([]interface{})[1].(map[string]interface{})["description"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCategory) TestGetCategoryByID() {
	testCase := []struct {
		Name                   string
		Method                 string
		ParamID                string
		ParamUser              dto.User
		MockReturnBody         dto.GetCategory
		MockReturnError        error
		HasReturnBody          bool
		ExpectedBodyCustomer   dto.GetCategory
		ExpectedBodyInstructor dto.GetCategoryInstructor
		ExpectedStatusCode     int
		ExpectedMesaage        string
	}{
		{
			"success get category by id and customer",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.GetCategory{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
				Courses: []dto.GetCourseWithoutCategory{
					{
						ID:              "abcde",
						Name:            "test",
						Description:     "test",
						Objective:       "test",
						Price:           10000,
						Discount:        0,
						Thumbnail:       "test",
						Capacity:        100,
						InstructorID:    "abcde",
						Rating:          5,
						Favorite:        false,
						NumberOfModules: 10,
					},
				},
			},
			nil,
			true,
			dto.GetCategory{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
				Courses: []dto.GetCourseWithoutCategory{
					{
						ID:              "abcde",
						Name:            "test",
						Description:     "test",
						Objective:       "test",
						Price:           10000,
						Discount:        0,
						Thumbnail:       "test",
						Capacity:        100,
						InstructorID:    "abcde",
						Rating:          5,
						Favorite:        false,
						NumberOfModules: 10,
					},
				},
			},
			dto.GetCategoryInstructor{},
			http.StatusOK,
			"success get category by id",
		},
		{
			"success get category by id and instructor",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			dto.GetCategory{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
				Courses: []dto.GetCourseWithoutCategory{
					{
						ID:              "abcde",
						Name:            "test",
						Description:     "test",
						Objective:       "test",
						Price:           10000,
						Discount:        0,
						Thumbnail:       "test",
						Capacity:        100,
						InstructorID:    "abcde",
						Rating:          5,
						Favorite:        false,
						NumberOfModules: 10,
					},
				},
			},
			nil,
			true,
			dto.GetCategory{},
			dto.GetCategoryInstructor{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
				Courses: []dto.GetCourseInstructorWithoutCategory{
					{
						ID:              "abcde",
						Name:            "test",
						Description:     "test",
						Objective:       "test",
						Price:           10000,
						Discount:        0,
						Thumbnail:       "test",
						Capacity:        100,
						InstructorID:    "abcde",
						Rating:          5,
						NumberOfModules: 10,
					},
				},
			},
			http.StatusOK,
			"success get category by id",
		},
		{
			"failed get category by id",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.GetCategory{},
			errors.New("error"),
			false,
			dto.GetCategory{},
			dto.GetCategoryInstructor{},
			http.StatusInternalServerError,
			"fail get category by id",
		},
		{
			"failed get category by id",
			"GET",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.GetCategory{},
			gorm.ErrRecordNotFound,
			false,
			dto.GetCategory{},
			dto.GetCategoryInstructor{},
			http.StatusNotFound,
			"fail get category by id",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetCategoryByID", v.ParamID, v.ParamUser).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/categories/"+v.ParamID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/categories/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(v.ParamID)
			if v.ParamUser.Role == "customer" {
				ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.ParamUser.ID, Role: v.ParamUser.Role}})
			} else {
				ctx.Set("user", &jwt.Token{Claims: &middlewareInstructor.JwtInstructorClaims{ID: v.ParamUser.ID, Role: v.ParamUser.Role}})
			}

			err := s.categoryController.GetCategoryByID(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])

			if v.HasReturnBody {
				if v.ParamUser.Role == "customer" {
					s.Equal(v.ExpectedBodyCustomer.Name, resp["category"].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyCustomer.Description, resp["category"].(map[string]interface{})["description"])
					s.Equal(v.ExpectedBodyCustomer.Courses[0].Name, resp["category"].(map[string]interface{})["courses"].([]interface{})[0].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyCustomer.Courses[0].Description, resp["category"].(map[string]interface{})["courses"].([]interface{})[0].(map[string]interface{})["description"])
				} else {
					s.Equal(v.ExpectedBodyInstructor.Name, resp["category"].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyInstructor.Description, resp["category"].(map[string]interface{})["description"])
					s.Equal(v.ExpectedBodyInstructor.Courses[0].Name, resp["category"].(map[string]interface{})["courses"].([]interface{})[0].(map[string]interface{})["name"])
					s.Equal(v.ExpectedBodyInstructor.Courses[0].Description, resp["category"].(map[string]interface{})["courses"].([]interface{})[0].(map[string]interface{})["description"])
				}
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func (s *suiteCategory) TestUpdateCategory() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.CategoryTransaction
		ParamID            string
		User               dto.User
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success update category",
			"POST",
			dto.CategoryTransaction{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			nil,
			http.StatusOK,
			"success update category",
		},
		{
			"fail bind data",
			"POST",
			dto.CategoryTransaction{
				Name:        "test",
				Description: "test",
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
			"fail update category",
			"POST",
			dto.CategoryTransaction{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail update category",
		},
		{
			"fail update category",
			"POST",
			dto.CategoryTransaction{
				ID:          "abcde",
				Name:        "test",
				Description: "test",
			},
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "instructor",
			},
			errors.New("fail update category"),
			http.StatusInternalServerError,
			"fail update category",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("UpdateCategory", v.Body).Return(v.MockReturnError)
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

			err := s.categoryController.UpdateCategory(ctx)
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

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
