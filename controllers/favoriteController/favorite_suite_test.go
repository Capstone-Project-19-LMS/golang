package favoriteController

import (
	"encoding/json"
	"errors"
	middlewareCustomer "golang/app/middlewares/costumer"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/favoriteService/favoriteMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteFavorite struct {
	suite.Suite
	favoriteController *FavoriteController
	mock               *favoriteMockService.FavoriteMock
}

func (s *suiteFavorite) SetupTest() {
	mock := &favoriteMockService.FavoriteMock{}
	s.mock = mock
	s.favoriteController = &FavoriteController{
		FavoriteService: s.mock,
	}
}

func (s *suiteFavorite) TestAddFavorite() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.FavoriteTransaction
		User               dto.User
		ParamCourseID string
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success favorite course",
			"POST",
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			nil,
			http.StatusOK,
			"success favorite course",
		},
		{
			"There is an empty field",
			"POST",
			dto.FavoriteTransaction{
				CustomerID: "abcde",
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
			"fail favorite course",
			"POST",
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			errors.New("fail favorite course"),
			http.StatusInternalServerError,
			"fail favorite course",
		},
		{
			"fail favorite course",
			"POST",
			dto.FavoriteTransaction{
				CustomerID: "abcde",
				CourseID:   "abcde",
			},
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			"abcde",
			errors.New(constantError.ErrorCustomerAlreadyFavoriteCourse),
			http.StatusBadRequest,
			"fail favorite course",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("AddFavorite", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/favorite", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/favorite/:courseId")
			if i != 1 {
				ctx.SetParamNames("courseId")
				ctx.SetParamValues(v.ParamCourseID)
			}
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.favoriteController.AddFavorite(ctx)
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

func (s *suiteFavorite) TestDeleteFavorite() {
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
			"success delete favorite course",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			nil,
			http.StatusOK,
			"success delete favorite course",
		},
		{
			"fail delete favorite course because favorite course not found",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			gorm.ErrRecordNotFound,
			http.StatusNotFound,
			"fail delete favorite course",
		},
		{
			"fail delete favorite course because error",
			"DELETE",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			errors.New("fail delete favorite course"),
			http.StatusInternalServerError,
			"fail delete favorite course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("DeleteFavorite", v.CourseID, v.User.ID).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/favorite/"+v.CourseID, nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/favorite/:courseId")
			ctx.SetParamNames("courseId")
			ctx.SetParamValues(v.CourseID)
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.favoriteController.DeleteFavorite(ctx)
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

func (s *suiteFavorite) TestGetCourseByCustomerID() {
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
			"success get favorite course by customer",
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
			"success get favorite course",
		},
		{
			"fail get favorite course by customer",
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
			"fail get favorite course",
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetFavoriteByCustomerID", v.User.ID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			// Create request
			r := httptest.NewRequest(v.Method, "/favorite", nil)
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			ctx := e.NewContext(r, w)
			ctx.SetPath("/favorite")
			ctx.Set("user", &jwt.Token{Claims: &middlewareCustomer.JwtCostumerClaims{ID: v.User.ID, Role: v.User.Role}})

			err := s.favoriteController.GetFavoriteCourseByCustomerID(ctx)
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

func TestSuiteFavorite(t *testing.T) {
	suite.Run(t, new(suiteFavorite))
}