package assignmentcontroller

import (
	"encoding/json"
	"errors"
	middlewareCustomer "golang/app/middlewares/costumer"
	middlewareInstructor "golang/app/middlewares/instructor"
	"golang/models/dto"
	"golang/service/categoryService/categoryMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteCategory struct {
	suite.Suite
	assignmentController *AssignmentController
	mock                 *categoryMockService.CategoryMock
}

func (s *suiteCategory) SetupTest() {
	mock := &categoryMockService.CategoryMock{}
	s.mock = mock
	s.assignmentController = &AssignmentController{
		AssignmentService: s.mock,
	}
}

func (s *suiteCategory) TestGetCategoryByID() {
	testCase := []struct {
		Name               string
		Method             string
		ParamID            string
		ParamUser          dto.User
		MockReturnBody     dto.GetCategory
		MockReturnError    error
		HasReturnBody      bool
		ExpectedBody       dto.GetCategory
		ExpectedStatusCode int
		ExpectedMesaage    string
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
				s.Equal(v.ExpectedBody.Name, resp["category"].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody.Description, resp["category"].(map[string]interface{})["description"])
				s.Equal(v.ExpectedBody.Courses[0].Name, resp["category"].(map[string]interface{})["courses"].([]interface{})[0].(map[string]interface{})["name"])
				s.Equal(v.ExpectedBody.Courses[0].Description, resp["category"].(map[string]interface{})["courses"].([]interface{})[0].(map[string]interface{})["description"])
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteCategory(t *testing.T) {
	suite.Run(t, new(suiteCategory))
}
